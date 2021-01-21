package hconvert

import (
	"fmt"
)

const (
	// MaxNumChars is the maximum number of characters in a character set.
	MaxNumChars int = 256
)

// CharSet holds a character set for encoding/decoding.
type CharSet struct {
	charSet []rune
	padding rune
}

// NewCharSet creates a new CharSet based on the provided character set. The values in the provided
// set should be organized in ascending order, so that the lowest value is first and the highest
// value is last. For example, a hexadecimal set would start with '0' and end with 'F'. Duplicate
// characters are not allowed.
func NewCharSet(set []rune) (CharSet, error) {
	length := len(set)
	if length == 0 {
		return CharSet{}, fmt.Errorf("missing character set")
	} else if length > MaxNumChars {
		return CharSet{}, fmt.Errorf("maximum of %d characters allowed", MaxNumChars)
	}

	// Make sure there are no duplicate characters.
	dupCheckMap := make(map[rune]bool)
	for _, c := range set {
		if dupCheckMap[c] {
			return CharSet{}, fmt.Errorf("duplicate character in set (%s)", string(c))
		}
		dupCheckMap[c] = true
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

// Len returns the number of runes in the character set.
func (c *CharSet) Len() int {
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

// mapEncode builds a map for converting from value to rune.
func (c *CharSet) mapEncode() map[int]rune {
	m := make(map[int]rune)

	if c != nil {
		for i, v := range c.Characters() {
			m[i] = v
		}
	}

	return m
}

// mapDecode builds a map for converting from rune to value.
func (c *CharSet) mapDecode() map[rune]int {
	m := make(map[rune]int, c.Len()*2)

	if c != nil {
		for i, v := range c.Characters() {
			m[v] = i
		}
	}

	return m
}
