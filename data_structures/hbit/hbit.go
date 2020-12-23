// Package hbit provides a safe and fast utility for storing and operating on bits.
package hbit

import (
	"fmt"
	"strings"
	"go/token"
	"io"
)


var (
	// This is the standard error message when trying to use an invalid buffer.
	errBadBuf = fmt.Errorf("Must create bit buffer with New() first")
)


// Buffer is the main type for this package. It holds the internal information about the bit buffer.
type Buffer struct {
	head *bnode
	tail *bnode
}

type bnode struct {
	next *bnode
	prev *bnode
	v     bool
}


// New creates a new bit buffer.
func New() *Buffer {
	return new(Buffer)
}

// Bit gets the boolean status (set or unset) of the bit at the provided index.
func (b *Buffer) Bit(index int) bool {
	node, err := b.getNode(index)
	if err != nil {
		return false
	}

	return node.v
}

// Bits gets the number of bits in the buffer, or -1 on error.
func (b *Buffer) Bits() int {
	if b == nil {
		return -1
	}

	cnt := 0
	node := b.head
	for node != nil {
		cnt++
		node = node.next
	}

	return cnt
}

// Offset gets ther number of bits forward ther buffer has been advanced, or -1 on error.
func (b *Buffer) Offset() int {
	if b == nil {
		return -1
	}

	cnt := 0
	node := b.tail
	for node != nil {
		cnt++
		node = node.prev
	}

	return cnt
}

// Copy creates a new buffer with the first n bits of the original buffer.
func (b *Buffer) Copy(n int) *Buffer {
	if b == nil {
		return nil
	}

	// We will copy these reference values over ...
	ref := b.head
	// ... into this new buffer.
	nb := New()

	// If the buffer is empty or no nodes were requested, then we don't have anything to copy.
	if b.head == nil || n < 1 {
		return nb
	}

	// Set up the start of the buffer.
	nb.head = new(bnode)
	node := nb.head
	node.v = ref.v

	// Go through the rest of the nodes in the original buffer.
	for i := 1; i < n; i++ {
		// Check if it's safe to move on first.
		ref = ref.next
		if ref == nil {
			break
		}

		// It's safe. Move on to the next node.
		node.appendNodeVal(nil, ref.v)
		node = node.next
	}

	return nb
}

// Recalibrate realigns the bits to the beginning of the buffer.
func (b *Buffer) Recalibrate() error {
	if b == nil {
		return errBadBuf
	}

	// All we have to do is cut off the tail.
	b.tail = nil

	return nil
}

// Reset resets the bit buffer to its initial state.
func (b *Buffer) Reset() error {
	if b == nil {
		return errBadBuf
	}

	b.head = nil
	b.tail = nil

	return nil
}

// String gets a string representation of the binary data in the buffer.
func (b *Buffer) String() string {
	return b.string_int(false)
}

// Display gets a string representation of the binary data in the buffer, with a single space between nibbles and a
// double space between bytes.
func (b *Buffer) Display() string {
	return b.string_int(true)
}


// Read reads len(p) bytes of bits from the buffer into p. It will return the number of bytes read into p, or io.EOF if
// the buffer is empty. io.EOF will only be returned if the buffer is empty before any bytes have been read into p. If
// there are not enough bits to fill all of the last byte, then the rest of the byte will be false bits.
func (b *Buffer) Read(p []byte) (int, error) {
	if b == nil {
		return 0, errBadBuf
	}

	// Check if our buffer has anything to read.
	if b.head == nil {
		return 0, io.EOF
	}

	length := len(p)
	node := b.head
	cnt := 0
	for i := 0; i < length; i++ {
		if node == nil {
			break
		}
		p[i] = 0

		for j := 0; j < 8; j++ {
			if node == nil {
				break
			}

			if node.v {
				p[i] |= (1 << uint(j))
			}
			node = node.next
			cnt++
		}
	}

	// Note: The calculation (cnt+7)/8 ensures that we account for untouched (and therefore false) bits in the last byte.
	_, err := b.Advance(cnt)
	return (cnt+7)/8, err
}

// ReadByte reads out one byte of bits at the index. This will not advance the buffer.
func (b *Buffer) ReadByte(index int) (byte, error) {
	node, err := b.getNode(index)
	if err != nil {
		return 0, err
	}

	var bt byte
	for i := 0; i < 8; i++ {
		if node == nil {
			break
		}

		if node.v {
			bt |= (1 << uint(i))
		}
		node = node.next
	}

	return bt, nil
}

// ReadInt reads out the 32-bit decimal representation of the bits at the index. This will not advance the buffer.
func (b *Buffer) ReadInt(index int) (int, error) {
	node, err := b.getNode(index)
	if err != nil {
		return 0, err
	}

	var n int32
	for i := 0; i < 32; i++ {
		if node == nil {
			break
		}

		if node.v {
			n |= (1 << uint(i))
		}
		node = node.next
	}

	return int(n), nil
}

// ReadFrom reads from r and append the bytes to the buffer. It will return the number of bytes read, and possibly an
// error.
func (b *Buffer) ReadFrom(r io.Reader) (int, error) {
	if b == nil {
		return 0, errBadBuf
	} else if r == nil {
		return 0, io.EOF
	}

	n, err := io.Copy(b, r)
	if n == 0 && err == nil {
		err = io.ErrNoProgress
	}
	return int(n), err
}

// Write appends the entire contents of p to the buffer.
func (b *Buffer) Write(p []byte) (int, error) {
	end, err := b.getEnd()
	if err != nil {
		return 0, err
	}

	skip := false
	if end == nil {
		// This means the buffer is empty. We'll create a node now to make setup easy and then skip past it later.
		b.head = new(bnode)
		end = b.head
		skip = true
	}

	length := len(p)
	for i := 0; i < length; i++ {
		for j := 0; j < 8; j++ {
			v := bitOn(p[i], j)
			end.appendNodeVal(nil, v)
			end = end.next
		}
	}

	if skip {
		// Skip past the dummy node we had to create earlier.
		b.head = b.head.next
		b.head.prev = nil
	}

	return length, nil
}

// WriteBit appends a bit to the end of the buffer.
func (b *Buffer) WriteBit(v bool) error {
	end, err := b.getEnd()
	if err != nil {
		return err
	}

	if end == nil {
		// This means the buffer is empty.
		b.head = new(bnode)
		b.head.v = v
	} else {
		end.appendNodeVal(nil, v)
	}

	return nil
}

// WriteByte appends an octet of bits to the end of the buffer. The bits will be added low to high.
func (b *Buffer) WriteByte(nb byte) error {
	_, err := b.Write([]byte{nb})
	return err
}

// WriteBytes appends the provided bytes to the end of the buffer.
func (b *Buffer) WriteBytes(bytes ...byte) error {
	_, err := b.Write(bytes)
	return err
}

// WriteString appends the provided string (in bytes) to the end of the buffer.
func (b *Buffer) WriteString(s string) error {
	_, err := b.Write([]byte(s))
	return err
}

// SetBit sets the value of a particular bit in the buffer.
func (b *Buffer) SetBit(index int, v bool) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	node.v = v

	return nil
}

// SetBytes sets the value of a range of bits in the buffer.
func (b *Buffer) SetBytes(index int, ref []byte) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	for _, octet := range ref {
		for i := 0; i < 8; i++ {
			if node == nil {
				return nil
			}
			node.v = bitOn(octet, i)
			node = node.next
		}
	}

	return nil
}

// RemoveBit cuts out the bit at the index.
func (b *Buffer) RemoveBit(index int) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	next := node.next
	prev := node.prev
	if next != nil {
		next.prev = prev
	}
	if prev != nil {
		prev.next = next
	}

	// If the first bit was specified, then we have to move the start of the buffer forward.
	if index == 0 {
		b.head = next
	}

	return nil
}

// RemoveBits cuts out the bits at the index.
func (b *Buffer) RemoveBits(index, n int) error {
	if n < 1 {
		return nil
	}

	start, err := b.getNode(index)
	if err != nil {
		return err
	}

	// Figure out how far out we're going to cut. If end is nil, then all the way.
	end, _ := b.getNode(index + n)

	// Remove all nodes inclusive of the start node and exclusive of the end node.
	prev := start.prev
	if prev != nil {
		prev.next = end
	}
	if end != nil {
		end.prev = prev
	}

	// If the first bit was specified, then we have to move the start of the buffer forward.
	if index == 0 {
		b.head = end
	}

	return nil
}

// Advance moves the start of the buffer forward a number of bits. It will return the number of bits moved.
func (b *Buffer) Advance(n int) (int, error) {
	if b == nil {
		return 0, errBadBuf
	} else if n < 0 {
		return 0, fmt.Errorf("Invalid number")
	}

	for i := 0; i < n; i++ {
		node := b.head
		if node == nil {
			return i, nil
		}

		// Move the start of the head forward one.
		b.head = b.head.next
		if b.head != nil {
			b.head.prev = nil
		}

		// Move the node from the head to the tail.
		if b.tail != nil {
			b.tail.appendNode(node)
		}
		b.tail = node
		b.tail.next = nil
	}

	return n, nil
}

// Rewind moves the start of buffer back a number of bits, or to the initial start. It will return the number of bits
// moved.
func (b *Buffer) Rewind(n int) (int, error) {
	if b == nil {
		return 0, errBadBuf
	} else if n < 0 {
		return 0, fmt.Errorf("Invalid number")
	}

	for i := 0; i < n; i++ {
		node := b.tail
		if node == nil {
			return i, nil
		}

		// Move the start of the tail back one.
		b.tail = node.prev
		if b.tail != nil {
			b.tail.next = nil
		}

		// Move the node from the tail to the head.
		if b.head != nil {
			node.appendNode(b.head)
		}
		b.head = node
		b.head.prev = nil
	}

	return n, nil
}

// Join appends a different buffer to the end of the current one. For safety, the current buffer will take ownership of
// the second buffer.
func (b *Buffer) Join(nb *Buffer) error {
	end, err := b.getEnd()
	if err != nil {
		return err
	}

	// Sanity check the new buffer.
	if nb == nil || nb.head == nil {
		// Nothing to add.
		return nil
	}

	if end == nil {
		// This means the buffer is empty.
		b.head = nb.head
	} else {
		end.appendNode(nb.head)
	}

	nb.head = nil
	return nil
}


// ANDBit performs the bitwise operation AND ('&') on the specified bit with the reference bit.
func (b *Buffer) ANDBit(index int, ref bool) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, ref, token.AND)
}

// ORBit performs the bitwise operation OR ('|') on the specified bit with the reference bit.
func (b *Buffer) ORBit(index int, ref bool) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, ref, token.OR)
}

// XORBit performs the bitwise operation XOR ('^') on the specified bit with the reference bit.
func (b *Buffer) XORBit(index int, ref bool) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, ref, token.XOR)
}

// ANDBytes performs the bitwise operation AND ('&') on the buffer with the reference bytes.
func (b *Buffer) ANDBytes(ref []byte) error {
	return b.opBytes(ref, token.AND)
}

// ORBytes performs the bitwise operation OR ('|') on the buffer with the reference bytes.
func (b *Buffer) ORBytes(ref []byte) error {
	return b.opBytes(ref, token.OR)
}

// XORBytes performs the bitwise operation XOR ('^') on the buffer with the reference bytes.
func (b *Buffer) XORBytes(ref []byte) error {
	return b.opBytes(ref, token.XOR)
}

// ANDBuffer performs the bitwise operation AND ('&') on the buffer with the reference buffer.
func (b *Buffer) ANDBuffer(ref *Buffer) error {
	return b.opBuf(ref, token.AND)
}

// ORBuffer performs the bitwise operation OR ('|') on the buffer with the reference buffer.
func (b *Buffer) ORBuffer(ref *Buffer) error {
	return b.opBuf(ref, token.OR)
}

// XORBuffer performs the bitwise operation XOR ('^') on the buffer with the reference buffer.
func (b *Buffer) XORBuffer(ref *Buffer) error {
	return b.opBuf(ref, token.XOR)
}

// ShiftLeft shifts the bits in the buffer to the left. This is equivalent to the bitwise operation '<<'.
func (b *Buffer) ShiftLeft(n int) error {
	if n < 0 {
		return fmt.Errorf("Invalid number")
	}

	end, err := b.getEnd()
	if err != nil {
		return err
	}

	if end == nil {
		// This means the buffer is empty.
		return nil
	}

	for i := 0; i < n; i++ {
		// Add a false bit at the end of the buffer.
		end.appendNode(nil)
		end = end.next

		// Pop the first bit.
		next := b.head.next
		b.head = next
		b.head.prev = nil
	}

	return nil
}

// ShiftRight shifts the bits in the buffer to the right. This is equivalent to the bitwise operation '>>'.
func (b *Buffer) ShiftRight(n int) error {
	if n < 0 {
		return fmt.Errorf("Invalid number")
	}

	end, err := b.getEnd()
	if err != nil {
		return err
	}

	if end == nil {
		// This means the buffer is empty.
		return nil
	}

	for i := 0; i < n; i++ {
		// Insert a false bit into the beginning of the buffer.
		node := new(bnode)
		node.appendNode(b.head)
		b.head = node

		// Cut off the last bit.
		end = end.prev
		end.next = nil
	}

	return nil
}

// NOTBit negates the specified bit. This is equivalent to the bitwise operation '~'.
func (b *Buffer) NOTBit(index int) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, false, token.NOT)
}

// NOTBits negates the first n bits in the buffer. This is equivalent to the bitwise operation '~'.
func (b *Buffer) NOTBits(n int) error {
	if n < 0 {
		return fmt.Errorf("Invalid range")
	}

	ref := make([]byte, n)
	return b.opBytes(ref, token.NOT)
}


// Create a new node and link it after the given node.
func (bn *bnode) appendNode(node *bnode) {
	if node == nil {
		node = new(bnode)
	}
	bn.next = node
	node.prev = bn
}

// Same as appendNode(), but also assign a value to the new node.
func (bn *bnode) appendNodeVal(node *bnode, v bool) {
	bn.appendNode(node)
	bn.next.v = v
}

// Check if a certain bit in a certain byte is set or not.
func bitOn(b byte, bit int) bool {
	if b & (1 << uint(bit)) > 0 {
		return true
	}

	return false
}

// Get the last node in the buffer.
func (b *Buffer) getEnd() (*bnode, error) {
	if b == nil {
		return nil, errBadBuf
	} else if b.head == nil {
		return nil, nil
	}

	node := b.head
	for node.next != nil {
		node = node.next
	}

	return node, nil
}

// Get the node at a given index.
func (b *Buffer) getNode(index int) (*bnode, error) {
	if b == nil {
		return nil, errBadBuf
	} else if index < 0 {
		return nil, fmt.Errorf("Invalid index")
	}

	node := b.head
	for i := 0; i < index; i++ {
		if node == nil {
			break
		}
		node = node.next
	}

	if node == nil {
		return nil, io.EOF
	}

	return node, nil
}

// Perform a bitwise operation on a bit.
func opBit(bit *bnode, ref bool, t token.Token) error {
	switch t {
	case token.AND:
		if bit.v && !ref {
			bit.v = false
		}
	case token.OR:
		if ref {
			bit.v = true
		}
	case token.XOR:
		if ref {
			if bit.v {
				bit.v = false
			} else {
				bit.v = true
			}
		}
	case token.NOT:
		bit.v = !bit.v
	default:
		return fmt.Errorf("internal misuse of opBit method")
	}

	return nil
}

// Perform a bitwise operation over a byte range.
func (b *Buffer) opBytes(ref []byte, t token.Token) error {
	if b == nil {
		return errBadBuf
	}

	node := b.head
	for _, octet := range ref {
		for i := 0; i < 8; i++ {
			if node == nil {
				return nil
			}

			v := bitOn(octet, i)
			if err := opBit(node, v, t); err != nil {
				return err
			}
			node = node.next
		}
	}

	return nil
}

// Perform a bitwise operation using another buffer as the reference.
func (b *Buffer) opBuf(ref *Buffer, t token.Token) error {
	if b == nil || ref == nil {
		return errBadBuf
	}

	node := b.head
	refNode := ref.head
	for node != nil && refNode != nil {
		if err := opBit(node, refNode.v, t); err != nil {
			return err
		}
		node = node.next
		refNode = refNode.next
	}

	return nil
}

// Print string from data.
func (b *Buffer) string_int(pretty bool) string {
	var sb strings.Builder

	if b == nil {
		return "<nil>"
	} else if b.head == nil {
		return "<empty>"
	}

	node := b.head
	cnt := 1
	for node != nil {
		if node.v {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		node = node.next

		if pretty {
			if cnt % 8 == 0 {
				sb.WriteString("  ")
			} else if cnt % 4 == 0 {
				sb.WriteString(" ")
			}
		}
		cnt++
	}

	s := sb.String()
	s = strings.Trim(sb.String(), " ")

	return s
}
