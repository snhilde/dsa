// Package hbit provides a safe and fast utility for storing and operating on bits.
package hbit

import (
	"errors"
	"strings"
	"go/token"
	"io"
)


// Buffer is the main type for this package. It holds the internal information about the bit buffer.
type Buffer struct {
	head *bnode
	tail *bnode
}

type bnode struct {
	next *bnode
	prev *bnode
	val   bool
}


// Create a new bit buffer.
func New() *Buffer {
	return new(Buffer)
}

// Get boolean status (on or off) of bit at provided index.
func (b *Buffer) Bit(index int) bool {
	node, err := b.getNode(index)
	if err != nil {
		return false
	}

	return node.val
}

// Get number of bits in buffer, or -1 on error.
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

// Get number of bits forward buffer has been advanced, or -1 on error.
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

// Create a new buffer with the first n bits of the original buffer.
func (b *Buffer) Copy(num int) *Buffer {
	if b == nil {
		return nil
	}

	// We will copy these reference values over ...
	ref := b.head
	// ... into this new buffer.
	nb := New()

	// If the buffer is empty or no nodes were requested, then we don't have anything to copy.
	if b.head == nil || num < 1 {
		return nb
	}

	// Set up the start of the buffer.
	nb.head = new(bnode)
	node := nb.head
	node.val = ref.val

	// Go through the rest of the nodes in the original buffer.
	for i := 1; i < num; i++ {
		// Check if it's safe to move on first.
		ref = ref.next
		if ref == nil {
			break
		}

		// It's safe. Move on to the next node.
		node.appendNodeVal(nil, ref.val)
		node = node.next
	}

	return nb
}

// Realign the bits to the beginning of the buffer.
func (b *Buffer) Recalibrate() error {
	if b == nil {
		return bufErr()
	}

	// All we have to do is cut off the tail.
	b.tail = nil

	return nil
}

// Reset the bit buffer to its initial state.
func (b *Buffer) Reset() error {
	if b == nil {
		return bufErr()
	}

	b.head = nil
	b.tail = nil

	return nil
}

// Get a string representation of the binary data in the buffer.
func (b *Buffer) String() string {
	return b.string_int(false)
}

// Get a string representation of the binary data in the buffer, with a single space between nibbles and a double space
// between bytes.
func (b *Buffer) Display() string {
	return b.string_int(true)
}


// Read len(p) bytes of bits from the buffer into p.
// Returns number of bytes read into p, or io.EOF if the buffer is empty. If there are not enough bits to fill all of
// the last byte, then the rest of the byte will be false bits. In the same vein, the returned number of bytes does not
// guarantee that the last bytes is full.
// io.EOF will only be returned if the buffer is empty before any bytes have been read into p.
func (b *Buffer) Read(p []byte) (int, error) {
	if b == nil {
		return 0, bufErr()
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

			if node.val {
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

// Read out one byte of bits at the index. This will not advance the buffer.
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

		if node.val {
			bt |= (1 << uint(i))
		}
		node = node.next
	}

	return bt, nil
}

// Read out the 32-bit decimal representation of the bits at the index. This will not advance the buffer.
func (b *Buffer) ReadInt(index int) (int, error) {
	node, err := b.getNode(index)
	if err != nil {
		return 0, err
	}

	var num int32
	for i := 0; i < 32; i++ {
		if node == nil {
			break
		}

		if node.val {
			num |= (1 << uint(i))
		}
		node = node.next
	}

	return int(num), nil
}

// Read from r and append bytes to buffer.
// Returns number of bytes read, and possibly an error.
func (b *Buffer) ReadFrom(r io.Reader) (int, error) {
	if b == nil {
		return 0, bufErr()
	} else if r == nil {
		return 0, io.EOF
	}

	n, err := io.Copy(b, r)
	if n == 0 && err == nil {
		err = io.ErrNoProgress
	}
	return int(n), err
}

// Append the entire contents of p to the buffer.
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
			val := bitOn(p[i], j)
			end.appendNodeVal(nil, val)
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

// Append a bit to the end of the buffer.
func (b *Buffer) WriteBit(val bool) error {
	end, err := b.getEnd()
	if err != nil {
		return err
	}

	if end == nil {
		// This means the buffer is empty.
		b.head = new(bnode)
		b.head.val = val
	} else {
		end.appendNodeVal(nil, val)
	}

	return nil
}

// Append an octet to the end of the buffer. The bits will be added low to high.
func (b *Buffer) WriteByte(nb byte) error {
	_, err := b.Write([]byte{nb})
	return err
}

// Append bytes to the end of the buffer.
func (b *Buffer) WriteBytes(bytes ...byte) error {
	_, err := b.Write(bytes)
	return err
}

// Append string bytes to the end of the buffer.
func (b *Buffer) WriteString(s string) error {
	_, err := b.Write([]byte(s))
	return err
}


// Set the value of a particular bit in the buffer.
func (b *Buffer) SetBit(index int, val bool) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	node.val = val

	return nil
}

// Set the value of a range of bits in the buffer.
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
			node.val = bitOn(octet, i)
			node = node.next
		}
	}

	return nil
}

// Cut out the bit at the index.
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

// Cut out the bits at the index.
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

// Move start of buffer forward a number of bits.
// Returns the number of bits moved.
func (b *Buffer) Advance(n int) (int, error) {
	if b == nil {
		return 0, bufErr()
	} else if n < 0 {
		return 0, errors.New("Invalid number")
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

// Move start of buffer back a number of bits, or to the initial start.
// Returns the number of bits moved.
func (b *Buffer) Rewind(n int) (int, error) {
	if b == nil {
		return 0, bufErr()
	} else if n < 0 {
		return 0, errors.New("Invalid number")
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

// Append a different buffer to the end of the current one. For safety, the current buffer will take ownership of the
// second buffer.
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


// AND the specified bit with the reference bit. This is equivalent to the bitwise operation '&'.
func (b *Buffer) ANDBit(index int, ref bool) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, ref, token.AND)
}

// OR the specified bit with the reference bit. This is equivalent to the bitwise operation '|'.
func (b *Buffer) ORBit(index int, ref bool) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, ref, token.OR)
}

// XOR the specified bit with the reference bit. This is equivalent to the bitwise operation '^'.
func (b *Buffer) XORBit(index int, ref bool) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, ref, token.XOR)
}

// AND the buffer against the reference bytes. This is equivalent to the bitwise operation '&'.
func (b *Buffer) ANDBytes(ref []byte) error {
	return b.opBytes(ref, token.AND)
}

// OR the buffer against the reference bytes. This is equivalent to the bitwise operation '|'.
func (b *Buffer) ORBytes(ref []byte) error {
	return b.opBytes(ref, token.OR)
}

// XOR the buffer against the reference bytes. This is equivalent to the bitwise operation '^'.
func (b *Buffer) XORBytes(ref []byte) error {
	return b.opBytes(ref, token.XOR)
}

// AND the buffer against the reference buffer. This is equivalent to the bitwise operation '&'.
func (b *Buffer) ANDBuffer(ref *Buffer) error {
	return b.opBuf(ref, token.AND)
}

// OR the buffer against the reference buffer. This is equivalent to the bitwise operation '|'.
func (b *Buffer) ORBuffer(ref *Buffer) error {
	return b.opBuf(ref, token.OR)
}

// XOR the buffer against the reference buffer. This is equivalent to the bitwise operation '^'.
func (b *Buffer) XORBuffer(ref *Buffer) error {
	return b.opBuf(ref, token.XOR)
}

// Shift the bits in the buffer to the left. This is equivalent to the bitwise operation '<<'.
func (b *Buffer) ShiftLeft(n int) error {
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

// Shift the bits in the buffer to the right. This is equivalent to the bitwise operation '>>'.
func (b *Buffer) ShiftRight(n int) error {
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

// Negate the specified bit. This is equivalent to the bitwise operation '~'.
func (b *Buffer) NOTBit(index int) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, false, token.NOT)
}

// Negate the first n bits in the buffer. This is equivalent to the bitwise operation '~'.
func (b *Buffer) NOTBits(n int) error {
	if n < 0 {
		return errors.New("Invalid range")
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
func (bn *bnode) appendNodeVal(node *bnode, val bool) {
	bn.appendNode(node)
	bn.next.val = val
}

// Check if a certain bit in a certain byte is set or not.
func bitOn(b byte, bit int) bool {
	if b & (1 << uint(bit)) > 0 {
		return true
	}

	return false
}

// This is the standard error message for passing an invalid buffer.
func bufErr() error {
	return errors.New("Must create bit buffer with New() first")
}

// Get the last node in the buffer.
func (b *Buffer) getEnd() (*bnode, error) {
	if b == nil {
		return nil, bufErr()
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
		return nil, bufErr()
	} else if index < 0 {
		return nil, errors.New("Invalid index")
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
		if bit.val && !ref {
			bit.val = false
		}
	case token.OR:
		if ref {
			bit.val = true
		}
	case token.XOR:
		if ref {
			if bit.val {
				bit.val = false
			} else {
				bit.val = true
			}
		}
	case token.NOT:
		bit.val = !bit.val
	default:
		return errors.New("internal misuse of opBit method")
	}

	return nil
}

// Perform a bitwise operation over a byte range.
func (b *Buffer) opBytes(ref []byte, t token.Token) error {
	if b == nil {
		return bufErr()
	}

	node := b.head
	for _, octet := range ref {
		for i := 0; i < 8; i++ {
			if node == nil {
				return nil
			}

			val := bitOn(octet, i)
			if err := opBit(node, val, t); err != nil {
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
		return bufErr()
	}

	node := b.head
	refNode := ref.head
	for node != nil && refNode != nil {
		if err := opBit(node, refNode.val, t); err != nil {
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
		if node.val {
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
