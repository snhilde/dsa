// Package hbit provides a utility for storing and operating on bits.
package hbit

import (
	"errors"
)


// --- PACKAGE TYPES ---
// Buffer is the main type for this package. It holds the internal information about the bit buffer.
type Buffer struct {
	bits []byte // internal bit buffer

	// We are not going to resize our byte slice to constantly fit tightly around the data. Instead, we're going to grow
	// the buffer as needed but keep the same size if the buffer is advanced or subtracted from. We'll use the values
	// below to keep track of where in the buffer the actual data is stored.
	begin_index int // index at the start of the first byte currently used
	begin_bits  int // number of bits (low to high) currently used in the first byte
	end_index   int // index immediately after the last byte currently used
	end_bits    int // number of bits (low to high) currently used in the last byte
}


// --- ENTRY FUNCTIONS AND OVERVIEW METHODS ---
// Create a new bit buffer.
func New() *Buffer {
	return new(Buffer)
}

// Length of internal buffer, or -1 on invalid object.
func (b *Buffer) Len() int {
	if b == nil {
		return -1
	}

	return len(b.bits)
}

// Capacity of internal buffer, or -1 on invalid object.
func (b *Buffer) Cap() int {
	if b == nil {
		return -1
	}

	return cap(b.bits)
}

// Number of bits in buffer, or -1 on invalid object.
func (b *Buffer) Bits() int {
	if b == nil {
		return -1
	}

	// If we don't have any bits in the first byte, then the buffer is empty.
	if b.begin_bits == 0 {
		return 0
	}

	// Start the calculation by assuming the entire range of bits is currently used.
	bits := b.end_index - b.begin_index

	// If the start of our buffer is shifted over (like can happen after an Advance() operation), then we won't be using
	// the entire first byte. In this case, we need to subtract the unused bits from the total count.
	bits -= (8 - b.begin_bits)

	// If we aren't using all of the last byte, then we need to subtract the unused bits from the total count.
	bits -= (8 - b.end_bits)

	return bits
}
