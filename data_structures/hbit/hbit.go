// Package hbit provides a compact and efficient utility for storing and operating on bits.
package hbit

import (
	"errors"
	"strings"
	"go/token"
)


// --- PACKAGE TYPES ---
// Buffer is the main type for this package. It holds the internal information about the bit buffer.
type Buffer struct {
	buf []byte // internal bit buffer

	// We are not going to resize our byte slice to constantly fit tightly around the data, and we're not going to align
	// bits to byte boundaries lest we interfere with other bytes. Instead, we're going to grow the buffer as needed but
	// keep the same size if the buffer is advanced or subtracted from. We'll use the values below to keep track of
	// where in the buffer the actual data is stored.
	start    int // index of the first byte currently used in the internal buffer
	end      int // index of the last byte currently used in the internal buffer
	offset   int // number of bits forward (low to high) the first byte is offset
	end_bits int // number of bits (low to high) currently used in the last byte
}


// --- ENTRY FUNCTIONS AND OVERVIEW METHODS ---
// Create a new bit buffer.
func New() *Buffer {
	var b Buffer

	b.buf = make([]byte, 1)

	return &b
}

// Get length of internal buffer, or -1 on invalid object.
func (b *Buffer) Len() int {
	if !sanityCheck(b) {
		return -1
	}

	return len(b.buf)
}

// Get capacity of internal buffer, or -1 on invalid object.
func (b *Buffer) Cap() int {
	if !sanityCheck(b) {
		return -1
	}

	return cap(b.buf)
}

// Get boolean status (on or off) of bit at provided index.
func (b *Buffer) Bit(index int) bool {
	if !sanityCheck(b) {
		return false
	} else if index < 0 || index >= b.Bits() {
		return false
	}

	// Figure out the actual byte and bit in the internal buffer where we're looking.
	byte_i, bit_i := b.iToPos(index)

	return bitOn(b.buf[byte_i], bit_i)
}

// Get number of bits in buffer, or -1 on invalid object.
func (b *Buffer) Bits() int {
	if !sanityCheck(b) {
		return -1
	}

	if b.start == b.end {
		return b.end_bits
	}

	// We'll start by assuming that all bytes up to the last byte are full.
	bytes := b.end - b.start
	bits := bytes * 8

	// If the starting bit in the buffer is not the first bit in the first byte (like can happen after an Advance()
	// operation), then we have to subtract the unused bits from the total.
	bits -= b.offset

	// Add in the number of bits being used at the end of the buffer.
	bits += b.end_bits

	return bits
}

// Get the current number of bits away from a byte boundary, or -1 on invalid object.
func (b *Buffer) Offset() int {
	if !sanityCheck(b) {
		return -1
	}

	return b.offset
}

// Reset the bit buffer to its initial state.
func (b *Buffer) Reset() error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	}

	nb := New()
	b.buf = nb.buf
	b.start = nb.start
	b.end = nb.end
	b.offset = nb.offset
	b.end_bits = nb.end_bits

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


// --- METHODS FOR OPERATING ON INTERNAL BUFFER ---
// Increase the internal buffer length by n bytes.
func (b *Buffer) Grow(n int) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if n < 0 {
		return errors.New("Invalid size")
	}

	buf := make([]byte, b.Len() + n)
	copy(buf, b.buf)
	b.buf = buf

	return nil
}

// Create a new Buffer with the first n bits of the original Buffer.
func (b *Buffer) Copy(n int) (*Buffer, error) {
	if !sanityCheck(b) {
		return nil, errors.New("Must create bit buffer with New() first")
	}

	// Make a mirror image of the original Buffer.
	var nb Buffer
	nb.buf = b.buf
	nb.start = b.start
	nb.end = b.end
	nb.offset = b.offset
	nb.end_bits = b.end_bits

	// Calculate our cut-off point.
	byte_i, bit_i := nb.iToPos(n)
	nb.end = byte_i
	nb.end_bits = bit_i + 1 // +1 to convert from index to count

	// Clean up our new Buffer.
	if err := nb.Recalibrate(); err != nil {
		return nil, err
	}

	return &nb, nil
}

// Realign the bits to the beginning of the buffer, and shrink the buffer to the appropriate size.
func (b *Buffer) Recalibrate() error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	}

	// Align to a byte boundary. This method will also set the appropriate values for b.offset and b.end_bits.
	if err := b.ShiftLeft(b.offset); err != nil {
		return err
	}

	// Resize the buffer to fit snuggly around the bits, and move everything to the far left.
	bytes := b.end - b.start
	buf := make([]byte, bytes + 1)
	copy(buf, b.buf[b.start:])
	b.buf = buf

	b.start = 0
	b.end = bytes

	// Clear out any bit beyond what we're now using.
	var tmp byte
	for i := 0; i < b.end_bits; i++ {
		tmp |= (1 << uint(i))
	}
	b.buf[b.end] |= tmp

	return nil
}


// --- METHODS FOR ADDING AND REMOVING BITS ---
// Add a bit to the end of the buffer.
func (b *Buffer) AddBit(bit byte) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if bit != 0 && bit != 1 {
		return errors.New("Invalid bit")
	}

	if b.end_bits == 8  || (b.start == b.end && b.offset + b.end_bits == 8) {
		// This byte is full. Move to the next byte.
		b.end++
		b.end_bits = 0
	}

	// If needed, add more room.
	b.checkSpace()

	b.setBit(b.end, b.end_bits, int(bit))
	b.end_bits++

	return nil
}

// Add an octet to the end of the buffer.
func (b *Buffer) AddByte(nb byte) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	}

	// We'll add the new bits in one-by-one to easily and safely account for any overflow.
	for i := 0; i < 8; i++ {
		bit := 0
		if bitOn(nb, i) {
			bit = 1
		}
		if err := b.AddBit(bit); err != nil {
			return err
		}
	}

	return nil
}

// Add bytes to the end of the buffer.
func (b *Buffer) AddBytes(nbs []byte) error {
	for _, nb := range nbs {
		if err := b.AddByte(nb); err != nil {
			return err
		}
	}

	return nil
}

// Move start of buffer forward a number of bits.
func (b *Buffer) Advance(n int) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if n < 0 {
		return errors.New("Invalid number")
	}

	// Calculate the new position.
	b.start, b.offset = b.iToPos(n)

	// Figure out if the end of the buffer was affected by this at all.
	if b.start > b.end {
		// We blew past the end of the buffer. Let's line them up again.
		b.end = b.start
		b.end_bits = 0
	} else if b.start == b.end {
		// We're only using one byte in the buffer. Let's see if we have any bits left.
		if b.offset >= b.end_bits {
			b.end_bits = 0
		} else {
			b.end_bits -= b.offset
		}
	}

	// Now, let's make sure our internal buffer is OK with all these changes.
	b.checkSpace()

	return nil
}

// Move start of buffer back a number of bits.
func (b *Buffer) Reverse(n int) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if n < 0 {
		return errors.New("Invalid number")
	}

	// TODO: implement

	return nil
}


// --- METHODS FOR BITWISE OPERATIONS ---
// AND the specified bit with the reference bit. This is equivalent to the bitwise operation '&'.
func (b *Buffer) ANDBit(index, ref int) error {
	return b.opBit(index, ref, token.AND)
}

// AND the buffer against the reference bytes. This is equivalent to the bitwise operation '&'.
func (b *Buffer) ANDBytes(ref []byte) error {
	return b.opBytes(ref, token.AND)
}

// TODO: ANDBuf

// OR the specified bit with the reference bit. This is equivalent to the bitwise operation '|'.
func (b *Buffer) ORBit(index, ref int) error {
	return b.opBit(index, ref, token.OR)
}

// OR the buffer against the reference bytes. This is equivalent to the bitwise operation '|'.
func (b *Buffer) ORBytes(ref []byte) error {
	return b.opBytes(ref, token.OR)
}

// XOR the specified bit with the reference bit. This is equivalent to the bitwise operation '^'.
func (b *Buffer) XORBit(index, ref int) error {
	return b.opBit(index, ref, token.XOR)
}

// XOR the buffer against the reference bytes. This is equivalent to the bitwise operation '^'.
func (b *Buffer) XORBytes(ref []byte) error {
	return b.opBytes(ref, token.XOR)
}

// Negate the specified bit. This is equivalent to the bitwise operation '~'.
func (b *Buffer) NOTBit(index int) error {
	return b.opBit(index, 0, token.NOT)
}

// Negate the first n bits in the buffer. This is equivalent to the bitwise operation '~'.
func (b *Buffer) NOTBytes(n int) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if n < 0 {
		return errors.New("Invalid range")
	} else if n == 0 {
		// Already done.
		return nil
	}

	return nil
}

// Shift the bits in the buffer to the left. This is equivalent to the bitwise operation '<<'.
func (b *Buffer) ShiftLeft(width int) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if width < 0 {
		return errors.New("Invalid range")
	}

	if width >= b.Len() * 8 || (b.start == 0 && b.offset == 0 && width >= b.Bits()) {
		// All we have to do is turn everything off.
		b.buf = make([]byte, b.Len(), b.Cap())
		return nil
	}

	return nil
}

// Shift the bits in the buffer to the right. This is equivalent to the bitwise operation '>>'.
func (b *Buffer) ShiftRight(width int) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if width < 0 {
		return errors.New("Invalid range")
	}

	if width >= b.Len() * 8 || (b.start == 0 && b.offset == 0 && width >= b.Bits()) {
		// All we have to do is turn everything off.
		b.buf = make([]byte, b.Len(), b.Cap())
		return nil
	}

	return nil
}


// --- HELPER FUNCTIONS ---
// Check if a certain bit in a certain byte is set or not.
func bitOn(b byte, bit int) bool {
	if b & (1 << uint(bit)) > 0 {
		return true
	}

	return false
}

// Make sure the Buffer was created properly.
func sanityCheck(b *Buffer) bool {
	// Check if Buffer was probably created with New() or not. If we only have a pointer and no underlying structure or
	// the internal buffer has not been created yet, then the Buffer was not created with New().
	if b == nil {
		return false
	} else if b.buf == nil {
		return false
	}

	return true
}

// Check if the internal buffer has enough space.
func (b *Buffer) checkSpace() {
	if b.end >= b.Len() {
		// We need more space in the buffer. We'll opt for twice as much room as we currently need.
		l := b.end * 2
		if b.Cap() >= l {
			// We only have to reslice.
			b.buf = b.buf[:l]
		} else {
			// Grow our buffer to make room.
			b.Grow(l - b.Len())
		}
	}
}

// Convert the bit index to the actual location (byte and bit) in the internal buffer.
func (b *Buffer) iToPos(index int) (int, int) {
	byte_i := b.start + (index / 8)
	bit_i := b.offset + (index % 8)
	if bit_i > 7 {
		byte_i++
		bit_i -= 8
	}

	return byte_i, bit_i
}

// Perform a bitwise operation on a bit.
func (b *Buffer) opBit(index, ref int, t token.Token) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if index < 0 {
		return errors.New("Invalid index")
	} else if index >= b.Bits() {
		return errors.New("Index out of range")
	} else if ref != 0 && ref != 1 {
		return errors.New("Invalid reference bit")
	}

	// Find the position of the bit in question.
	byte_i, bit_i := b.iToPos(index)

	on := bitOn(b.buf[byte_i], bit_i)
	bit := 0
	switch t {
	case token.AND:
		if on && ref == 1 {
			bit = 1
		}
	case token.OR:
		if on || ref == 1 {
			bit = 1
		}
	case token.XOR:
		if (on && ref == 0) || (!on && ref == 1) {
			bit = 1
		}
	case token.NOT:
		if !on {
			bit = 1
		}
	default:
		return errors.New("internal misuse of opBit method")
	}

	b.setBit(byte_i, bit_i, bit)

	return nil
}

// Perform a bitwise operation over a byte range.
func (b *Buffer) opBytes(ref []byte, t token.Token) error {
	// Here's how we're going to do this for efficiency:
	// 1. Create a Buffer
	// 2. Offset the new Buffer to match the original Buffer (if there is an offset).
	// 3. Add the reference bytes to the new Buffer.
	// 4. Apply the operation to the head and tail dangling bits, if there are any.
	// 5. Apply the operation to the bytes in the middle.
	// This will allow us to operate at the byte level as much as possible.
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if len(ref) == 0 {
		// Already done.
		return nil
	}

	// 1. Create a new Buffer.
	buf := New()

	// 2. Set an identical offset.
	offset := (b.start * 8) + b.offset
	if err := buf.Advance(offset); err != nil {
		return err
	}

	// 3. Add the reference bytes.
	if err := buf.AddBytes(ref); err != nil {
		return err
	}

	// 4. Apply the operation to the dangling bytes.
	for i := 0; i < 8 - b.offset; i++ {
		ref := 0
		if buf.Bit(i) {
			ref = 1
		}
		if err := b.opBit(i, ref, t); err != nil {
			return err
		}
	}

	if b.Bits() <= buf.Bits() {
		for i := b.Bits() - b.end_bits; i < b.end_bits; i++ {
			ref := 0
			if buf.Bit(i) {
				ref = 1
			}
			if err := b.opBit(i, ref, t); err != nil {
				return err
			}
		}
	}

	// 5. Apply the operation to all bytes in the middle.
	end := b.end
	if buf.end < b.end {
		end = buf.end
	}
	for i := b.start + 1; i < end; i++ {
		switch t {
		case token.AND:
			b.buf[i] &= buf.buf[i]
		case token.OR:
			b.buf[i] |= buf.buf[i]
		case token.XOR:
			b.buf[i] ^= buf.buf[i]
		default:
			return errors.New("internal misuse of opBytes method")
		}
	}

	return nil
}

// Set the bit at the given byte and index to val.
func (b *Buffer) setBit(byte_i, bit_i, val int) {
	b.buf[byte_i] |= byte(val << uint(bit_i))
}

// Print string from data..
func (b *Buffer) string_int(pretty bool) string {
	var sb strings.Builder

	if !sanityCheck(b) {
		return "<nil>"
	} else if b.Bits() == 0 {
		return "<empty>"
	}

	for i := b.start; i <= b.end; i++ {
		low := 0
		high := 8
		if b.start == b.end {
			// If we're only printing the one byte, then we have to tweak the calculation.
			low = b.offset
			high = b.offset + b.end_bits
		} else {
			if i == b.start {
				low = b.offset
			} else if i == b.end {
				high = b.end_bits
			}
		}
		for j := low; j < high; j++ {
			if j == 0 && pretty {
				sb.WriteString("  ")
			} else if j == 4 && pretty {
				sb.WriteString(" ")
			}

			if bitOn(b.buf[i], j) {
				sb.WriteString("1")
			} else {
				sb.WriteString("0")
			}
		}
	}

	s := sb.String()
	s = strings.Trim(sb.String(), " ")

	return s
}
