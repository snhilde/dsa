// Package hconvert performs base conversion to and from any arbitrary character set (up to 256 runes in length).
package hconvert

import (
	"bytes"
	"errors"
	"math/big"
)

var (
	ErrBadConverter error = errors.New("Bad Converter")
)

// Converter holds information about the conversion process, including the specified character sets.
type Converter struct {
	// binary representation of the number
	buf *bytes.Buffer

	// character sets to use for decoding and encoding
	decCharSet CharSet
	encCharSet CharSet
}

// NewConverter creates a new Converter object with the provided character sets. If a character set is not needed (such
// as when doing only encoding or only decoding), pass `CharSet{}`.
func NewConverter(decode CharSet, encode CharSet) Converter {
	c := new(Converter)

	c.buf = new(bytes.Buffer)
	c.decCharSet = decode
	c.encCharSet = encode

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

// DecodeCharSet returns the CharSet used for decoding.
func (c *Converter) DecodeCharSet() CharSet {
	if c == nil {
		return CharSet{}
	}

	return c.decCharSet
}

// EncodeCharSet returns the CharSet used for encoding.
func (c *Converter) EncodeCharSet() CharSet {
	if c == nil {
		return CharSet{}
	}

	return c.encCharSet
}
