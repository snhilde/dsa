// Package hconvert performs base conversion to and from any arbitrary character set (up to
// MaxNumChars runes in length).
package hconvert

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/snhilde/dsa/data_structures/hstack"
)

var (
	errBadConverter = fmt.Errorf("invalid converter")
	errBadCharSet   = fmt.Errorf("invalid character set")
	errNoCharSet    = fmt.Errorf("no character set provided")
)

// Converter holds information about the conversion process, including the specified character sets.
type Converter struct {
	// Character sets to use for decoding and encoding.
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

// SetDecodeCharSet sets the character set to use for decoding.
func (c *Converter) SetDecodeCharSet(charSet CharSet) error {
	if c == nil {
		return errBadConverter
	} else if len(charSet.charSet) == 0 {
		return errNoCharSet
	}

	c.decCharSet = charSet

	return nil
}

// SetEncodeCharSet sets the character set to use for encoding.
func (c *Converter) SetEncodeCharSet(charSet CharSet) error {
	if c == nil {
		return errBadConverter
	} else if len(charSet.charSet) == 0 {
		return errNoCharSet
	}

	c.encCharSet = charSet

	return nil
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

// Decode decodes s using the decoding character set and returns the binary data or any error
// encountered.
func (c *Converter) Decode(s string) ([]byte, error) {
	if c == nil {
		return nil, errBadConverter
	}

	return DecodeWith(s, c.decCharSet)
}

// DecodeFrom reads encoded data from r until EOF, decodes the data using the decoding character
// set, and returns the binary data or any error encountered.
func (c *Converter) DecodeFrom(r io.Reader) ([]byte, error) {
	if c == nil {
		return nil, errBadConverter
	}

	encoded, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return c.Decode(string(encoded))
}

// DecodeWith decodes s with the provided character set and returns the binary data or any error
// encountered.
func DecodeWith(s string, charSet CharSet) ([]byte, error) {
	if s == "" {
		return []byte{}, nil
	} else if charSet.Len() == 0 {
		return nil, errNoCharSet
	}

	// Get the rune->int mapping for this character set.
	decMap := charSet.mapDecode()

	// Binary data that we will decode to.
	binary := new(big.Int)

	// Numerical base of this character set, for calculating the value of each character at its
	// place in the string.
	base := big.NewInt(int64(charSet.Len()))

	// We'll use this to calculate the value of each character at its place in the string. Because
	// range reads the string from left to right, we have to start with the highest place and move
	// down to 1. For example, if we have a string of "489" representing a base10 number, then the
	// starting place is 100, the next place is 10, and the last place is 1. This givs a total value
	// of (100 * 4) + (10 * 8) + (1 * 9) = 489.
	place := new(big.Int)
	place.Exp(base, big.NewInt(int64(len(s)-1)), nil)

	padding := charSet.Padding()
	for _, r := range s {
		if padding != "" && string(r) == padding {
			continue
		}

		// Figure out the value of this character in the character set.
		v, ok := decMap[r]
		if !ok {
			return nil, errBadCharSet
		}

		// Add this value to the total sum according to its overall place in the string.
		value := big.NewInt(int64(v))
		value.Mul(value, place)
		binary.Add(binary, value)

		// Move down to the next position.
		place.Div(place, base)
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

// EncodeTo encodes p using the encoding character set and writes it to w.
func (c *Converter) EncodeTo(p []byte, w io.Writer) error {
	if c == nil {
		return errBadConverter
	}

	encoded, err := c.Encode(p)
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
	// Basic strategy:
	// 1. Set up all the values we need.
	// 2. While there is data in the buffer:
	//     a. Calculate the modulus of the current buffer.
	//     b. Divide the buffer by the base. This effectively "pops" the modulus from the buffer.
	//     c. Add the modulus to the stack, according to its rune mapping.
	// 3. Because the modulo operation removes the last character from the buffer, the string is
	//    going to be reversed. To solve this, we'll pop the values from the stack one-by-one and
	//    add them to the output buffer, which will reverse the string back to the correct order.
	if len(p) == 0 {
		return "", nil
	} else if charSet.Len() == 0 {
		return "", errNoCharSet
	}

	// int->rune mapping for this character set.
	decMap := charSet.mapEncode()

	// Binary data that we will encode.
	binary := new(big.Int)
	binary.SetBytes(p)

	// Numerical base of this character set, for determining the appropriate character at each place
	// in the output string.
	base := big.NewInt(int64(charSet.Len()))

	// Container to hold the mod value as it's calculated.
	mod := new(big.Int)

	// Stack to hold the string as it's created in reverse order.
	stack := hstack.New()

	zero := big.NewInt(0)
	for binary.Cmp(zero) > 0 {
		// Calculate this position's value (mod) and move down to the next position (divide).
		binary.DivMod(binary, base, mod)

		// Add the character for this position.
		// Note: mod can never be greater than base, which is the length of the character set.
		value := int(mod.Int64())
		stack.Add(decMap[value])
	}

	// Write out the string in stack order to reverse it back to the correct order.
	out := new(strings.Builder)
	for stack.Count() > 0 {
		out.WriteRune(stack.Pop().(rune))
	}

	return out.String(), nil
}
