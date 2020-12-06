// Package hconvert performs base conversion to and from any arbitrary character set (up to 256 runes in length).
package hconvert

import (
)

// Type Converter holds information about the conversion process, including the specified character sets.
type Converter struct {
	// binary representation of the number
	bin []byte

	// character sets to use for encoding and decoding
	encCharSet CharSet
	decCharSet CharSet
}

// NewConverter creates a new Converter object with the provided character sets. If a character set is not needed (such
// as when doing only encoding or only decoding), pass `CharSet{}`.
func NewConverter(encode CharSet, decode CharSet) Converter {
	c := new(Converter)

	c.encCharSet = encode
	c.decCharSet = decode

	return *c
}

// SetDecodeCharSet sets the character set to use for decoding.
func (c *Converter) SetDecodeCharSet(charSet CharSet) {
	if c != nil {
		c.decCharSet = charSet
	}
}

// SetEncodeCharSet sets the character set to use for encoding.
func (c *Converter) SetEncodeCharSet(charSet CharSet) {
	if c != nil {
		c.encCharSet = charSet
	}
}
