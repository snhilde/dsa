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

	// The original, encoded string.
	orig string

	// Internal buffer for storing decoded data.
	num *big.Int
}

// NewConverter creates a new Converter object with the provided character sets.
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

// DecodeCharSet returns the character set used for decoding.
func (c *Converter) DecodeCharSet() CharSet {
	if c == nil {
		return CharSet{}
	}

	return c.decCharSet
}

// EncodeCharSet returns the character set used for encoding.
func (c *Converter) EncodeCharSet() CharSet {
	if c == nil {
		return CharSet{}
	}

	return c.encCharSet
}

// Convert converts the string from the decode character set to the encode character set.
func (c *Converter) Convert(s string) (string, error) {
	switch {
	case c == nil:
		return "", errBadConverter
	case c.decCharSet.Len() == 0:
		return "", errNoCharSet
	case c.encCharSet.Len() == 0:
		return "", errNoCharSet
	}

	c.orig = s

	// Decode the data to binary.
	num, err := c.decode()
	if err != nil {
		return "", err
	}
	c.num = num

	// Encode the binary to the converted string.
	out, err := c.encode()
	if err != nil {
		return "", err
	}

	// If there wasn't any data in the buffer, then we can just return the zero-value character.
	if out == "" {
		out = string(c.encCharSet.Characters()[0])
	}

	return out, nil
}

// ConvertFrom reads an encoded string from r until EOF and converts it from the decode character
// set to the encode character set.
func (c *Converter) ConvertFrom(r io.Reader) (string, error) {
	encoded, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return c.Convert(string(encoded))
}

func (c *Converter) decode() (*big.Int, error) {
	// Get the rune->int mapping for this character set.
	decMap := c.decCharSet.mapDecode()

	// Binary data that we will decode to.
	binary := new(big.Int)

	// Numerical base of this character set, for calculating the value of each character at its
	// place in the string.
	base := big.NewInt(int64(c.decCharSet.Len()))

	// We'll use this to calculate the value of each character at its place in the string. Because
	// range reads the string from left to right, we have to start with the highest place and move
	// down to 1. For example, if we have a string of "489" representing a base10 number, then the
	// starting place is 100, the next place is 10, and the last place is 1. This gives a total
	// value of (100 * 4) + (10 * 8) + (1 * 9) = 489.
	significance := new(big.Int)
	significance.Exp(base, big.NewInt(int64(len(c.orig)-1)), nil)

	padding := c.decCharSet.Padding()
	for _, b := range c.orig {
		if padding != "" && string(b) == padding {
			continue
		}

		// Figure out the value of this character in the character set.
		v, ok := decMap[b]
		if !ok {
			return nil, errBadCharSet
		}

		// Add this value to the total sum according to its overall significance in the string.
		value := big.NewInt(int64(v))
		value.Mul(value, significance)
		binary.Add(binary, value)

		// Move down to the next position.
		significance.Div(significance, base)
	}

	return binary, nil
}

func (c *Converter) encode() (string, error) {
	// Basic strategy:
	// 1. Set up all the values we need.
	// 2. While there is data in the buffer:
	//     a. Calculate the modulus of the current buffer.
	//     b. Divide the buffer by the base. This effectively "pops" the modulus from the buffer.
	//     c. Add the modulus to the stack, according to its rune mapping.
	// 3. If the original string began with one or more characters that have a 0-value (according to
	//    their place in the character set), then we need to add an equivalent amount of 0-value
	//    characters in the new character set to the beginning of the string.
	// 4. Because the modulo operation removes the last character from the buffer, the string is
	//    going to be reversed. To solve this, we'll pop the values from the stack one-by-one and
	//    add them to the output buffer, which will reverse the string back to the correct order.
	if len(c.num.Bytes()) == 0 {
		return "", nil
	} else if c.encCharSet.Len() == 0 {
		return "", errNoCharSet
	}

	// int->rune mapping for this character set.
	encMap := c.encCharSet.mapEncode()

	// Binary data that we will encode.
	binary := c.num

	// Numerical base of this character set, for determining the appropriate character at each place
	// in the output string.
	base := big.NewInt(int64(c.encCharSet.Len()))

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
		stack.Add(encMap[value])
	}

	// Calculate how many leading 0-value characters the original string has.
	zeroChar := c.decCharSet.Characters()[0]
	leadingZeroes := 0
	for _, char := range c.orig {
		if char == zeroChar {
			leadingZeroes++
		} else {
			break
		}
	}

	// Add back in the appropriate number of zero-value characters for this character set.
	for i := leadingZeroes * c.decCharSet.Len(); i > 0; i /= c.encCharSet.Len() {
		stack.Add(encMap[0])
	}

	// Write out the string in stack order to reverse it back to the correct order.
	out := new(strings.Builder)
	for stack.Count() > 0 {
		out.WriteRune(stack.Pop().(rune))
	}

	return out.String(), nil
}
