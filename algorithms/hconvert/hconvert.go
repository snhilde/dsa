// Package hconvert performs base conversion to and from any arbitrary character set (up to
// MaxNumChars runes in length).
package hconvert

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"strings"
	"unicode"

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

	// The original encoded string.
	input string

	// Internal buffer for storing decoded data.
	num *big.Int

	// The converter encoded string.
	output string
}

// NewConverter creates a new Converter object with the provided character sets.
func NewConverter(decode CharSet, encode CharSet) Converter {
	var c Converter

	c.decCharSet = decode
	c.encCharSet = encode

	return c
}

// SetDecodeCharSet sets the character set to use for decoding.
func (c *Converter) SetDecodeCharSet(charSet CharSet) error {
	if c == nil {
		return errBadConverter
	} else if charSet.isEmpty() {
		return errNoCharSet
	}

	c.decCharSet = charSet

	return nil
}

// SetEncodeCharSet sets the character set to use for encoding.
func (c *Converter) SetEncodeCharSet(charSet CharSet) error {
	if c == nil {
		return errBadConverter
	} else if charSet.isEmpty() {
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
	if c == nil {
		return "", errBadConverter
	}

	c.input = s

	// Decode the data to binary.
	if err := c.decode(); err != nil {
		return "", err
	}

	// Encode the binary to the converted string.
	if err := c.encode(); err != nil {
		return "", err
	}

	// If there wasn't any data in the buffer, then we can just return the zero-value character.
	if c.output == "" {
		c.output = string(c.encCharSet.Characters()[0])
	}

	return c.output, nil
}

// ConvertFrom reads an encoded string from r until EOF and converts it from the decode character
// set to the encode character set.
func (c *Converter) ConvertFrom(r io.Reader) (string, error) {
	encoded, err := ioutil.ReadAll(r)
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}

	return c.Convert(string(encoded))
}

func (c *Converter) decode() error {
	// Basic strategy:
	// 1. Set up all the values we need.
	// 2. Iterate through every character in the encoded input string and perform these actions:
	//     a. Get the value of the character according to its place in the character set.
	//     b. Multiply that by the position of the character in the string (its significance).
	//     c. Add that value to the buffer.
	// 3. The buffer will now hold the value of the encode input string.
	if c == nil {
		return errBadConverter
	}

	c.num = nil
	if c.decCharSet.isEmpty() {
		return fmt.Errorf("no decode character set provided")
	} else if c.input == "" {
		return nil
	} else if !isPrintable(c.input) {
		return fmt.Errorf("not printable text")
	}

	// Make a new big number to hold the binary data as we decode.
	binary := new(big.Int)

	// Figure out the umerical base of this character set, which we'll use for calculating the value
	// of each character at its position in the string.
	base := big.NewInt(int64(c.decCharSet.Len()))

	// Get the rune->int mapping for this character set.
	decMap := c.decCharSet.mapDecode()

	// Figure out the significance of the first character in the string. We'll use this to calculate
	// the value of each character at its position in the string. Because range reads the string
	// from left to right, we have to start with the highest position and move down to 1. For
	// example, if we have a string of "489" representing a base10 number, then the starting
	// significance is 100, the next significance is 10, and the last significance is 1. This gives
	// a total value of (100 * 4) + (10 * 8) + (1 * 9) = 489.
	significance := new(big.Int)
	significance.Exp(base, big.NewInt(int64(len(c.input)-1)), nil)

	padding := c.decCharSet.Padding()
	for _, b := range c.input {
		if padding != "" && string(b) == padding {
			continue
		}

		// Figure out the value of this character in the character set.
		v, ok := decMap[b]
		if !ok {
			return errBadCharSet
		}

		// Add this value to the total sum according to its overall significance in the string.
		value := big.NewInt(int64(v))
		value.Mul(value, significance)
		binary.Add(binary, value)

		// Move down to the next position.
		significance.Div(significance, base)
	}

	// Store the raw data for encoding.
	c.num = binary

	return nil
}

func (c *Converter) encode() error {
	// Basic strategy:
	// 1. Set up all the values we need.
	// 2. While there is data in the buffer:
	//     a. Calculate the modulus of the current buffer according to our base. This tells us the
	//        value of the next encoding character.
	//     b. Divide the buffer by the base. This effectively "pops" the character from the buffer.
	//     c. Add the value to the stack, according to its rune mapping.
	// 3. If the original string began with one or more characters that have a 0-value (according to
	//    their place in the character set), then we need to add an equivalent amount of 0-value
	//    characters in the new character set to the beginning of the string.
	// 4. Because the modulo operation calculates the characters from smallest to largest, the
	//    string is going to be reversed. To solve this, we'll pop the values from the stack
	//    one-by-one and add them to the output buffer, which will reverse the string back to the
	//    correct order.
	if c == nil {
		return errBadConverter
	}

	c.output = ""
	if c.num == nil || len(c.num.Bytes()) == 0 {
		return nil
	} else if c.encCharSet.isEmpty() {
		return fmt.Errorf("no encode character set provided")
	} else if c.decCharSet.isEmpty() {
		return fmt.Errorf("no decode character set provided")
	}

	// Grab the binary data that we will encode.
	binary := c.num

	// Calculate the numerical base of this character set, which we'll use to determine the
	// appropriate character at each position in the output string.
	base := big.NewInt(int64(c.encCharSet.Len()))

	// Get the int->rune mapping for this character set.
	encMap := c.encCharSet.mapEncode()

	// Make a container to hold the mod value as it's calculated.
	mod := new(big.Int)

	// Make a stack to hold the string as it's created in reverse order.
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

	// Calculate the number of leading 0-value characters in the original string.
	zeroChar := c.decCharSet.Characters()[0]
	leadingZeroes := 0
	for _, char := range c.input {
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

	c.output = out.String()

	return nil
}

// isPrintable determines whether or not the string has only printable runes.
func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsGraphic(r) {
			return false
		}
	}

	return true
}
