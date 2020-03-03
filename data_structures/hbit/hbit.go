// Package hbit provides a utility for storing and operating on bits.
package hbit

import (
	"errors"
	"strings"
)


// --- PACKAGE TYPES ---
// Buffer is the main type for this package. It holds the internal information about the bit buffer.
type Buffer struct {
	buf []byte // internal bit buffer

	// We are not going to resize our byte slice to constantly fit tightly around the data, and we're not going to align
	// bits to byte boundaries lest we interfere with other bytes. Instead, we're going to grow the buffer as needed but
	// keep the same size if the buffer is advanced or subtracted from. We'll use the values below to keep track of
	// where in the buffer the actual data is stored.
	index_begin int // index at the start of the first byte currently used
	index_end   int // index at the start of the last byte currently used
	offset      int // number of bits forward (low to high) the first byte is offset
	end_bits    int // number of bits (low to high) currently used in the last byte
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

// Get number of bits in buffer, or -1 on invalid object.
func (b *Buffer) Bits() int {
	if !sanityCheck(b) {
		return -1
	}

	// We'll start by assuming that all bytes up to the last byte are full.
	bytes := b.index_end - b.index_begin
	bits := bytes * 8

	// If the starting bit in the buffer is not the first bit in the first byte (like can happen after an Advance()
	// operation), then we have to subtract the unused bits from the total.
	if b.index_begin != b.index_end {
		bits -= b.offset
	}

	// Add in the number of bits being used at the end of the buffer.
	bits += b.end_bits

	return bits
}

// Reset the bit buffer to its initial state.
func (b *Buffer) Reset() error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	}

	b.buf = make([]byte, 1)
	b.index_begin = 0
	b.index_end = 0
	b.offset = 0
	b.end_bits = 0

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
func (b *Buffer) AddBit(bit byte) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	} else if bit != 0 && bit != 1 {
		return errors.New("Invalid bit")
	}

	if b.end_bits == 8  || (b.index_begin == b.index_end && b.offset + b.end_bits == 8) {
		// This byte is full. Move to the next byte.
		b.index_end++
		b.end_bits = 0
	}

	// If needed, add more room.
	b.checkSpace()

	b.buf[b.index_end] |= byte(bit << uint(b.end_bits))
	b.end_bits++

	return nil
}

// Add a byte to the end of the buffer.
// will be split between two bytes in the buffer.
func (b *Buffer) AddByte(nb byte) error {
	var err error

	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	}

	// If the end of the buffer is aligned to a byte boundary, then we can add in the new byte directly.
	if (b.index_begin != b.index_end && b.end_bits == 0) ||
	   (b.index_begin == b.index_end && b.offset == 0 && b.end_bits == 0) {
		// If needed, add more room.
		b.checkSpace()
		b.buf[b.index_end] = nb
		b.index_end++
	} else {
		// The last bit in the buffer is not aligned to a byte boundary. We'll add the new bits in one-by-one to easily
		// and safely account for the overflow.
		for i := 0; i < 8; i++ {
			if checkBit(nb, i) {
				err = b.AddBit(1)
			} else {
				err = b.AddBit(0)
			}
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Add bytes to the end of the buffer.
func (b *Buffer) AddBytes(nbs []byte) error {
	for _, nb := range nbs {
		err := b.AddByte(nb)
		if err != nil {
			return err
		}
	}

	return nil
}

// Move buffer forward a number of bits.
func (b *Buffer) Advance(n int) error {
	if !sanityCheck(b) {
		return errors.New("Must create bit buffer with New() first")
	}

	// Calculate the big and small steps we'll have to take.
	bytes := n / 8
	bits := n % 8

	// Take the big step forward.
	b.index_begin += bytes

	// Take the small step forward.
	b.offset += bits
	if b.offset >= 8 {
		// Overflow to the next byte.
		b.index_begin ++
		b.offset -= 8
	}

	// Figure out if the end of the buffer was affected by this at all.
	if b.index_begin > b.index_end {
		// We blew past the end of the buffer. Let's line them up again.
		b.index_end = b.index_begin
		b.end_bits = 0
	} else if b.index_begin == b.index_end {
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


// --- HELPER FUNCTIONS ---
// Check if certain bit in certain byte is set or not.
func checkBit(b byte, bit int) bool {
	if b & (1 << uint(bit)) > 0 {
		return true
	}

	return false
}

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

func (b *Buffer) checkSpace() {
	if b.index_end >= b.Len() {
		// We need more space in the buffer. We'll opt for twice as much room as we currently need.
		l := b.index_end * 2
		if b.Cap() >= l {
			// We only have to reslice.
			b.buf = b.buf[:l]
		} else {
			// Grow our buffer to make room.
			s := make([]byte, l)
			copy(s, b.buf)
			b.buf = s
		}
	}
}

// internal helper for printing string from data.
func (b *Buffer) string_int(pretty bool) string {
	var sb strings.Builder

	if !sanityCheck(b) {
		return "<nil>"
	} else if b.Bits() == 0 {
		return "<empty>"
	}

	for i := b.index_begin; i <= b.index_end; i++ {
		low := 0
		high := 8
		if b.index_begin == b.index_end {
			// If we're only printing the one byte, then we have to tweak the calculation.
			low = b.offset
			high = b.offset + b.end_bits
		} else {
			if i == b.index_begin {
				low = b.offset
			} else if i == b.index_end {
				high = b.end_bits
			}
		}
		for j := low; j < high; j++ {
			if j == 0 && pretty {
				sb.WriteString("  ")
			} else if j == 4 && pretty {
				sb.WriteString(" ")
			}

			if checkBit(b.buf[i], j) {
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
