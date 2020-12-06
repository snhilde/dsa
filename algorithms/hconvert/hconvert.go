// Package hconvert performs base conversion to and from any arbitrary character set (up to 256 runes in length).
package hconvert

import (
	"bytes"
	"errors"
	"io/ioutil"
	"math/big"
)

var (
	ErrBadConverter error = errors.New("Bad Converter")
)

// Converter holds information about the conversion process, including the specified character sets.
type Converter struct {
	// binary representation of the data
	buf []byte

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

// DecodeFrom reads encoded data from r until EOF, decodes the data using the decoding character set, and stores it
// internally.
func (c *Converter) DecodeFrom(r io.Reader) error {
	if c == nil {
		return ErrBadConverter
	}

	encoded, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	binary, err := c.Decode(encoded)
	if err != nil {
		return err
	}

	c.buf = binary

	return nil
}

// EncodeTo encodes the internally stored data using the encoding character set and writes it to w.
func (c *Converter) EncodeTo(r io.Reader) error {
	if c == nil {
		return ErrBadConverter
	}

	if len(c.buf) == 0 {
		return nil
	}

	encoded, err := c.Encode(c.buf)
	if err != nil {
		return err
	}

	if n, err := r.Write(encoded); err != nil {
		return err
	} else if n != len(encoded) {
		return ErrShortWrite
	}

	return nil
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
