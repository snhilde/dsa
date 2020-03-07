// Package hbit provides a safe and fast utility for storing and operating on bits.
package hbit

import (
	"errors"
	"strings"
	"go/token"
)


// --- PACKAGE TYPES ---
// Buffer is the main type for this package. It holds the internal information about the bit buffer.
type Buffer struct {
	node *bnode
}

type bnode struct {
	next *bnode
	prev *bnode
	val   bool
}


// --- ENTRY FUNCTIONS AND OVERVIEW METHODS ---
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
	node := b.node
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

	// If the buffer is empty, then we don't have any offset.
	if b.node == nil {
		return 0
	}

	cnt := 0
	node := b.node
	for node.prev != nil {
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
	ref := b.node
	// ... into this new buffer.
	nb := New()

	// If the buffer is empty or no nodes were requested, then we don't have anything to copy.
	if b.node == nil || num < 1 {
		return nb
	}

	// Set up the start of the buffer.
	nb.node = new(bnode)
	node := nb.node
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

	// If we have anything in the buffer, remove the offset.
	if b.node != nil {
		b.node.prev = nil
	}

	return nil
}

// Reset the bit buffer to its initial state.
func (b *Buffer) Reset() error {
	if b == nil {
		return bufErr()
	}

	b.node = nil
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


// --- METHODS FOR ADDING AND REMOVING BITS ---
// Add a bit to the end of the buffer.
func (b *Buffer) AddBit(val bool) error {
	end, err := b.getEnd()
	if err != nil {
		return err
	}

	if end == nil {
		// This means the buffer is empty.
		b.node = new(bnode)
		b.node.val = val
	} else {
		end.appendNodeVal(nil, val)
	}

	return nil
}

// Add an octet to the end of the buffer. The bits will be added low to high.
func (b *Buffer) AddByte(nb byte) error {
	end, err := b.getEnd()
	if err != nil {
		return err
	}

	val := bitOn(nb, 0)
	if end == nil {
		// This means the buffer is empty.
		b.node = new(bnode)
		b.node.val = val
		end = b.node
	} else {
		end.appendNodeVal(nil, val)
		end = end.next
	}

	for i := 1; i < 8; i++ {
		val := bitOn(nb, i)
		end.appendNodeVal(nil, val)
		end = end.next
	}

	return nil
}

// Add bytes to the end of the buffer.
func (b *Buffer) AddBytes(nbs []byte) error {
	buf := New()
	for _, nb := range nbs {
		if err := buf.AddByte(nb); err != nil {
			return err
		}
	}

	return b.Merge(buf)
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
		b.node = next
	}

	return nil
}

// Cut out the bits at the index.
func (b *Buffer) RemoveBits(index, n int) error {
	// TODO: should removing all bits to the end also remove all offset bits before the start of the buffer?
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
		b.node = end
	}

	return nil
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

// Move start of buffer forward a number of bits. This will not go past the last bit in the buffer.
// Returns the number of bits moved.
func (b *Buffer) Advance(n int) (int, error) {
	if b == nil {
		return 0, bufErr()
	} else if n < 0 {
		return 0, errors.New("Invalid number")
	}

	if b.node == nil {
		return 0, nil
	}

	var i int
	for i = 0; i < n; i++ {
		if b.node.next == nil {
			break
		}
		b.node = b.node.next
	}

	return i, nil
}

// Move start of buffer back a number of bits, or to the initial start.
// Returns the number of bits moved.
func (b *Buffer) Rewind(n int) (int, error) {
	if b == nil {
		return 0, bufErr()
	} else if n < 0 {
		return 0, errors.New("Invalid number")
	}

	if b.node == nil {
		return 0, nil
	}

	var i int
	for i = 0; i < n; i++ {
		if b.node.prev == nil {
			break
		}
		b.node = b.node.prev
	}

	return i, nil
}

// Append a different buffer to the end of the current one. For safety, the current buffer will take ownership of the
// second buffer.
func (b *Buffer) Merge(nb *Buffer) error {
	end, err := b.getEnd()
	if err != nil {
		return err
	}

	// Sanity check the new buffer.
	if nb == nil || nb.node == nil {
		// Nothing to add.
		return nil
	}

	if end == nil {
		// This means the buffer is empty.
		b.node = nb.node
	} else {
		end.appendNode(nb.node)
	}

	nb.node = nil
	return nil
}


// --- METHODS FOR BITWISE OPERATIONS ---
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

// Negate the specified bit. This is equivalent to the bitwise operation '~'.
func (b *Buffer) NOTBit(index int) error {
	node, err := b.getNode(index)
	if err != nil {
		return err
	}

	return opBit(node, false, token.NOT)
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

// Negate the first n bits in the buffer. This is equivalent to the bitwise operation '~'.
func (b *Buffer) NOTBytes(n int) error {
	if n < 0 {
		return errors.New("Invalid range")
	}
	ref := make([]byte, n)
	return b.opBytes(ref, token.NOT)
}

// TODO: ANDBuf

// TODO: ORBuf

// TODO: XORBuf

// TODO: NOTBuf

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
		next := b.node.next
		prev := b.node.prev
		next.prev = prev
		if prev != nil {
			prev.next = next
		}
		b.node = next
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
		next := b.node
		prev := b.node.prev
		node.next = next
		node.prev = prev
		next.prev = node
		if prev != nil {
			prev.next = node
		}
		b.node = node

		end = end.prev
		end.next = nil
	}

	return nil
}


// --- METHODS FOR WRITING OUT BITS ---
// Write out the 32-bit decimal representation of the bits at the index, or -1 on error.
func (b *Buffer) WriteInt(index int) int {
	node, err := b.getNode(index)
	if err != nil {
		return 0
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

	return int(num)
}


// --- HELPER FUNCTIONS ---
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
	} else if b.node == nil {
		return nil, nil
	}

	node := b.node
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

	node := b.node
	for i := 0; i < index; i++ {
		if node == nil {
			break
		}
		node = node.next
	}

	if node == nil {
		return nil, errors.New("Index out of range")
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

	node := b.node
	for _, octet := range ref {
		for i := 0; i < 8; i++ {
			if node == nil {
				// We're out of bits to operate on in the buffer.
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

// Print string from data..
func (b *Buffer) string_int(pretty bool) string {
	var sb strings.Builder

	if b == nil {
		return "<nil>"
	} else if b.node == nil {
		return "<empty>"
	}

	node := b.node
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
