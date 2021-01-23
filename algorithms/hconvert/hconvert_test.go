package hconvert_test

import (
	"reflect"
	"testing"

	"github.com/snhilde/dsa/algorithms/hconvert"
)

func TestNewConverter(t *testing.T) {
	// We should be able to create a Converter with any combination of valid or invalid character
	// sets.
	enc := hconvert.CharSet{}
	dec := hconvert.CharSet{}
	if converter := hconvert.NewConverter(enc, dec); !reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create empty converter")
	}

	enc = hconvert.Base16CharSet()
	dec = hconvert.CharSet{}
	if converter := hconvert.NewConverter(enc, dec); reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create converter with only encoding character set")
	}

	enc = hconvert.CharSet{}
	dec = hconvert.Base32CharSet()
	if converter := hconvert.NewConverter(enc, dec); reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create converter with only decoding character set")
	}

	enc = hconvert.Base16CharSet()
	dec = hconvert.Base32CharSet()
	if converter := hconvert.NewConverter(enc, dec); reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create converter with both encoding and decoding character sets")
	}
}

func TestSetDecodeCharSet(t *testing.T) {
}

func TestSetEncodeCharSet(t *testing.T) {
}

func TestDecodeCharSet(t *testing.T) {
}

func TestEncodeCharSet(t *testing.T) {
}

func TestDecode(t *testing.T) {
}

func TestDecodeFrom(t *testing.T) {
}

func TestDecodeWith(t *testing.T) {
}

func TestEncode(t *testing.T) {
}

func TestEncodeTo(t *testing.T) {
}

func TestEncodeWith(t *testing.T) {
}
