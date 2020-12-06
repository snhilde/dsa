package hconvert

import (
	"errors"
)

const (
	// We are currently only accepting character sets up to 256 in length.
	maxChars int = 256
)

// CharSet holds a character set for encoding/decoding.
type CharSet struct {
	charSet []rune
	padding rune
}

// NewCharSet creates a new CharSet based on the provided character set. The values in the provided set should be
// organized in ascending order, so that the lowest value is first and the highest value is last. For example, a
// hexadecimal set would start with '0' and end with 'F'.
func NewCharSet(set []rune) (CharSet, error) {
	length := len(set)
	if length == 0 {
		return CharSet{}, errors.New("Missing character set")
	} else if length > maxChars {
		return CharSet{}, errors.New("Maximum of 256 characters allowed")
	}

	charSet := make([]rune, length)
	copy(charSet, set)

	return CharSet{charSet: charSet}, nil
}

// Padding returns the current padding character.
func (c *CharSet) Padding() rune {
	var pad rune
	if c != nil {
		pad = c.padding
	}
	return pad
}

// SetPadding sets the padding character to pad.
func (c *CharSet) SetPadding(pad rune) {
	if c != nil {
		c.padding = pad
	}
}

// Length returns the number of runes in the character set.
func (c *CharSet) Length() int {
	if c == nil {
		return -1
	}
	return len(c.charSet)
}

// Characters returns the characters in the character set.
func (c *CharSet) Characters() []rune {
	if c == nil {
		return nil
	}
	return c.charSet
}

// String returns the string representation of the concatenated runes in the character set.
func (c *CharSet) String() string {
	if c == nil {
		return ""
	}
	return string(c.charSet)
}
