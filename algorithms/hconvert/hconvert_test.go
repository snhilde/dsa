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
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	if converter := hconvert.NewConverter(dec, enc); !reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create empty converter")
	}

	dec = hconvert.CharSet{}
	enc = hconvert.Base16CharSet()
	if converter := hconvert.NewConverter(dec, enc); reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create converter with only encoding character set")
	}

	dec = hconvert.Base32CharSet()
	enc = hconvert.CharSet{}
	if converter := hconvert.NewConverter(dec, enc); reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create converter with only decoding character set")
	}

	dec = hconvert.Base32CharSet()
	enc = hconvert.Base16CharSet()
	if converter := hconvert.NewConverter(dec, enc); reflect.DeepEqual(converter, hconvert.Converter{}) {
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
