// Package hconvert performs base conversion to and from any arbitrary character set (up to 256 runes in length).
package hconvert

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
)

var (
	errBadConverter = fmt.Errorf("bad Converter")
	errNoCharSet    = fmt.Errorf("no character set provided")
	errBadCharSet   = fmt.Errorf("character set mismatch")
)

// Converter holds information about the conversion process, including the specified character sets.
type Converter struct {
	// binary representation of the data
	buf []byte

	// character sets to use for decoding and encoding
	decCharSet CharSet
	encCharSet CharSet
}

// NewConverter creates a new Converter object with the provided character sets. If a character set
// is not needed (such as when doing only encoding or only decoding), pass `CharSet{}`.
func NewConverter(decode CharSet, encode CharSet) Converter {
	c := new(Converter)

	c.decCharSet = decode
	c.encCharSet = encode

	return *c
}

// Decode decodes s using the decoding character set and returns the binary data or any error
// encountered.
func (c *Converter) Decode(s string) ([]byte, error) {
	if c == nil {
		return nil, errBadConverter
	}

	return DecodeWith(s, c.decCharSet)
}

// DecodeFrom reads encoded data from r until EOF, decodes the data using the decoding character
// set, and stores it internally.
func (c *Converter) DecodeFrom(r io.Reader) error {
	if c == nil {
		return errBadConverter
	}

	encoded, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	binary, err := c.Decode(string(encoded))
	if err != nil {
		return err
	}

	c.buf = binary

	return nil
}

// DecodeWith decodes s with the provided character set and returns the binary data or any error
// encountered.
func DecodeWith(s string, charSet CharSet) ([]byte, error) {
	if charSet.Len() <= 0 {
		return nil, errNoCharSet
	} else if s == "" {
		return []byte{}, nil
	}

	// Get the mapping from rune to int for this character set.
	decMap := charSet.mapDecode()

	// This is the binary data that we will decode to.
	binary := big.NewInt(0)

	// This is the base of this character set and the current index in the string.
	base := big.NewInt(int64(charSet.Len()))
	index := big.NewInt(1)
	index.Exp(base, big.NewInt(int64(len(s)-1)), big.NewInt(0))

	// This is the decoded value of the current digit.
	value := big.NewInt(0)

	padding := charSet.Padding()
	for _, r := range s {
		if r == padding {
			continue
		}

		v, ok := decMap[r]
		if !ok {
			return nil, errBadCharSet
		}

		// Add this value to the total sum according to its overall place in the string.
		value.SetInt64(int64(v))
		binary.Add(binary, value.Mul(value, index))

		// Move down to the next position.
		index.Mul(index, base)
	}

	return binary.Bytes(), nil
}

// Encode encodes p using the encoding character set and returns the encoded data or any error
// encountered.
func (c *Converter) Encode(p []byte) (string, error) {
	if c == nil {
		return "", errBadConverter
	}

	return EncodeWith(p, c.encCharSet)
}

// EncodeTo encodes the internally stored data using the encoding character set and writes it to w.
func (c *Converter) EncodeTo(w io.Writer) error {
	if c == nil {
		return errBadConverter
	}

	if len(c.buf) == 0 {
		return nil
	}

	encoded, err := c.Encode(c.buf)
	if err != nil {
		return err
	}

	if n, err := w.Write([]byte(encoded)); err != nil {
		return err
	} else if n != len(encoded) {
		return io.ErrShortWrite
	}

	return nil
}

// EncodeWith encodes p with the provided character set and returns the encoded data or any error
// encountered.
func EncodeWith(p []byte, charSet CharSet) (string, error) {
	if charSet.Len() <= 0 {
		return "", errNoCharSet
	} else if len(p) == 0 {
		return "", nil
	}

	return "", nil
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
