package hconvert_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/snhilde/dsa/algorithms/hconvert"
)

// testSet holds the information for running decoding/encoding tests.
type testSet struct {
	// Character set to use for the conversion.
	charSet hconvert.CharSet

	// Name of the character set.
	setName string

	// String to decode.
	decodeMe string

	// Data to encode.
	encodeMe []byte
}

var (
	// This is the binary equivalent of base10 1,000,000,000. We're going to use this as a common
	// check for encoding and decoding tests.
	billion = []byte{59, 154, 202, 0}

	// These are the common tests for decoding/encoding using the const characer sets.
	constTests = []testSet{
		{hconvert.ASCIICharSet(), "ASCII", string([]byte{3, 92, 107, 20, 0}), billion},
		{hconvert.Base2CharSet(), "Base2", "111011100110101100101000000000", billion},
		{hconvert.Base4CharSet(), "Base4", "323212230220000", billion},
		{hconvert.Base8CharSet(), "Base8", "7346545000", billion},
		{hconvert.Base10CharSet(), "Base10", "1000000000", billion},
		{hconvert.Base16CharSet(), "Base16", "3B9ACA00", billion},
		{hconvert.Base32CharSet(), "Base32", "5ZVSQA", billion},
		{hconvert.Base36CharSet(), "Base36", "GJDGXS", billion},
		{hconvert.Base58CharSet(), "Base58", "2XNGAK", billion},
		{hconvert.Base62CharSet(), "Base62", "BFp3qQ", billion},
		{hconvert.Base64CharSet(), "Base64", "7msoA", billion},
		{hconvert.Base64URLCharSet(), "Base64URL", "7msoA", billion},
		{hconvert.ASCII85CharSet(), "ASCII85", "4.=:l", billion},
		{hconvert.Z85CharSet(), "Z85", "jdsp(", billion},
	}
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

func TestBad(t *testing.T) {
	t.Parallel()

	var converter *hconvert.Converter

	if err := converter.SetDecodeCharSet(hconvert.Base10CharSet()); err == nil {
		t.Error("Failed bad object test for SetDecodeCharSet")
	}

	if err := converter.SetEncodeCharSet(hconvert.Base10CharSet()); err == nil {
		t.Error("Failed bad object test for SetEncodeCharSet")
	}

	if b, err := converter.Decode("abc"); b != nil || err == nil {
		t.Error("Failed bad object test for Decode")
	}

	if b, err := converter.DecodeFrom(strings.NewReader("1234")); b != nil || err == nil {
		t.Error("Failed bad object test for DecodeFrom")
	}

	if s, err := converter.Encode([]byte{0x01, 0x02}); s != "" || err == nil {
		t.Error("Failed bad object test for Encode")
	}

	if err := converter.EncodeTo(new(strings.Builder)); err == nil {
		t.Error("Failed bad object test for EncodeTo")
	}
}

func TestSetDecodeCharSet(t *testing.T) {
	t.Parallel()

	// Test that you can't set an empty set.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	converter := hconvert.NewConverter(dec, enc)
	if err := converter.SetDecodeCharSet(hconvert.CharSet{}); err == nil {
		t.Error("Passed setting empty decoding character set")
	}

	// Test setting a decode char set when one is not already set.
	dec = hconvert.CharSet{}
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetDecodeCharSet(hconvert.Base16CharSet()); err != nil {
		t.Error(err)
	} else if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base16CharSet()) {
		t.Error("Failed to set decoding character set when one is not already set")
	}

	// Test setting a decode char set when one is already set.
	dec = hconvert.Base2CharSet()
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetDecodeCharSet(hconvert.Base64URLCharSet()); err != nil {
		t.Error(err)
	} else if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base64URLCharSet()) {
		t.Error("Failed to set decoding character set when one is already set")
	}

	// Test setting a decode char set when both are set.
	dec = hconvert.Base10CharSet()
	enc = hconvert.Base4CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetDecodeCharSet(hconvert.Base10CharSet()); err != nil {
		t.Error(err)
	} else if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base10CharSet()) {
		t.Error("Failed to set decoding character set when both are set")
	}
}

func TestSetEncodeCharSet(t *testing.T) {
	t.Parallel()

	// Test that you can't set an empty set.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	converter := hconvert.NewConverter(dec, enc)
	if err := converter.SetEncodeCharSet(hconvert.CharSet{}); err == nil {
		t.Error("Passed setting empty encoding character set")
	}

	// Test setting an encode char set when one is not already set.
	dec = hconvert.CharSet{}
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetEncodeCharSet(hconvert.Base58CharSet()); err != nil {
		t.Error(err)
	} else if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base58CharSet()) {
		t.Error("Failed to set encoding character set when one is not already set")
	}

	// Test setting an encode char set when one is already set.
	dec = hconvert.CharSet{}
	enc = hconvert.ASCIICharSet()
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetEncodeCharSet(hconvert.Base10CharSet()); err != nil {
		t.Error(err)
	} else if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base10CharSet()) {
		t.Error("Failed to set encoding character set when one is already set")
	}

	// Test setting an encode char set when both are set.
	dec = hconvert.Base4CharSet()
	enc = hconvert.Base10CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetEncodeCharSet(hconvert.Base64CharSet()); err != nil {
		t.Error(err)
	} else if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base64CharSet()) {
		t.Error("Failed to set encoding character set when both are set")
	}
}

func TestDecodeCharSet(t *testing.T) {
	t.Parallel()

	// Test getting the decoding character set when nothing is set.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	converter := hconvert.NewConverter(dec, enc)
	if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.CharSet{}) {
		t.Error("Failed to retrieve decoding character set when nothing is set")
	}

	// Test getting the decoding character set when only that is set.
	dec = hconvert.Base2CharSet()
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base2CharSet()) {
		t.Error("Failed to retrieve decoding character set when only that is set")
	}

	// Test getting the decoding character set when only the encoding one is set.
	dec = hconvert.CharSet{}
	enc = hconvert.Base4CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.CharSet{}) {
		t.Error("Failed to retrieve decoding character set when only the encoding one is set")
	}

	// Test getting the decoding character set when both are set.
	dec = hconvert.Base10CharSet()
	enc = hconvert.Base8CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base10CharSet()) {
		t.Error("Failed to retrieve decoding character set when both are set")
	}
}

func TestEncodeCharSet(t *testing.T) {
	t.Parallel()

	// Test getting the encoding character set when nothing is set.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	converter := hconvert.NewConverter(dec, enc)
	if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.CharSet{}) {
		t.Error("Failed to retrieve encoding character set when nothing is set")
	}

	// Test getting the encoding character set when only that is set.
	dec = hconvert.CharSet{}
	enc = hconvert.Base58CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base58CharSet()) {
		t.Error("Failed to retrieve encoding character set when only that is set")
	}

	// Test getting the encoding character set when only the encoding one is set.
	dec = hconvert.Base64CharSet()
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.CharSet{}) {
		t.Error("Failed to retrieve encoding character set when only the encoding one is set")
	}

	// Test getting the encoding character set when both are set.
	dec = hconvert.Base32CharSet()
	enc = hconvert.Base2CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base2CharSet()) {
		t.Error("Failed to retrieve encoding character set when both are set")
	}
}

func TestDecode(t *testing.T) {
	t.Parallel()

	// Test that we can properly decode to the binary equivalent of one billion in base10 using a
	// converter with each of the const character sets.
	converter := new(hconvert.Converter)
	for _, test := range constTests {
		converter.SetDecodeCharSet(test.charSet)
		b, err := converter.Decode(test.decodeMe)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if !reflect.DeepEqual(b, test.encodeMe) {
			t.Error(test.setName, "- Decode failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", b)
		}
	}
}

func TestDecodeFrom(t *testing.T) {
	t.Parallel()

	// Test that we can properly read and decode data from the provided io.Reader using a converter
	// with each of the const character sets
	converter := new(hconvert.Converter)
	for _, test := range constTests {
		converter.SetDecodeCharSet(test.charSet)
		reader := strings.NewReader(test.decodeMe)
		b, err := converter.DecodeFrom(reader)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if !reflect.DeepEqual(b, test.encodeMe) {
			t.Error(test.setName, "- Decode failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", b)
		}
	}
}

func TestDecodeWith(t *testing.T) {
	t.Parallel()

	// Test that we can properly decode to the binary equivalent of one billion in base10 using the
	// const character sets.
	for _, test := range constTests {
		b, err := hconvert.DecodeWith(test.decodeMe, test.charSet)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if !reflect.DeepEqual(b, test.encodeMe) {
			t.Error(test.setName, "- Decode failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", b)
		}
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()

	// Test that we can properly encode the binary equivalent of one billion in base10 using a
	// converter with each of the const character sets.
	converter := new(hconvert.Converter)
	for _, test := range constTests {
		converter.SetEncodeCharSet(test.charSet)
		s, err := converter.Encode(test.encodeMe)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if s != test.decodeMe {
			t.Error(test.setName, "- Encode failed")
			t.Log("Expected:", test.decodeMe)
			t.Log("Received:", s)
		}
	}
}

func TestEncodeTo(t *testing.T) {
	t.Parallel()
}

func TestEncodeWith(t *testing.T) {
	t.Parallel()

	// Test that we can properly encode the binary equivalent of one billion in base10 using the
	// const character sets.
	for _, test := range constTests {
		s, err := hconvert.EncodeWith(test.encodeMe, test.charSet)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if s != test.decodeMe {
			t.Error(test.setName, "- Encode failed")
			t.Log("Expected:", test.decodeMe)
			t.Log("Received:", s)
		}
	}
}
