package hconvert_test

import (
	"reflect"
	"testing"

	"github.com/snhilde/dsa/algorithms/hconvert"
)

func TestNewConverter(t *testing.T) {
	t.Parallel()

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
	t.Parallel()
}

func TestSetEncodeCharSet(t *testing.T) {
	t.Parallel()
}

func TestDecodeCharSet(t *testing.T) {
	t.Parallel()
}

func TestEncodeCharSet(t *testing.T) {
	t.Parallel()
}

func TestDecode(t *testing.T) {
	t.Parallel()
}

func TestDecodeFrom(t *testing.T) {
	t.Parallel()
}

func TestDecodeWith(t *testing.T) {
	t.Parallel()
}

func TestEncode(t *testing.T) {
	t.Parallel()
}

func TestEncodeTo(t *testing.T) {
	t.Parallel()
}

func TestEncodeWith(t *testing.T) {
	t.Parallel()
}
