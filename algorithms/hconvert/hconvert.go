// Package hconvert performs base conversion to and from any arbitrary character set (up to 256 runes in length).
package hconvert

import (
)

type Converter struct {
	// binary representation of the number
	bin []byte

	// character sets to use for encoding and decoding
	encCharSet CharSet
	decCharSet CharSet
}
