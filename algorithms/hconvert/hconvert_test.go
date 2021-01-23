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

const (
	// This character is not in any of the standard character sets. Decoding with this should fail.
	invalidChar = "∞"
)

var (
	// These are the common tests for decoding/encoding using the standard characer sets.
	data = []byte{
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x00, 0x11, 0x22, 0x33,
		0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF,
	}
	standardTests = []testSet{
		{hconvert.ASCIICharSet(), "ASCII", string([]byte{4, 70, 69, 51, 98, 53, 60, 111, 60, 0, 17, 17, 12, 104, 69, 43, 25, 111, 8, 76, 106, 87, 60, 102, 119, 93, 127}), data},
		{hconvert.Base2CharSet(), "Base2", "10010001101000101011001111000100110101011110011011110111100000000000100010010001000110011010001000101010101100110011101111000100010011001101010101011101111001100110111011110111011111111", data},
		{hconvert.Base4CharSet(), "Base4", "102031011121320212223303132330000010102020303101011111212131320202121222223233030313132323333", data},
		{hconvert.Base8CharSet(), "Base8", "22150531704653633674000422106321052546357042315253571467367377", data},
		{hconvert.Base10CharSet(), "Base10", "27898229935051914141545580467023890909217686409246011135", data},
		{hconvert.Base16CharSet(), "Base16", "123456789ABCDEF00112233445566778899AABBCCDDEEFF", data},
		{hconvert.Base32CharSet(), "Base32", "SGRLHRGV433YACERDGRCVMZ3YRGNKXPGN33X7", data},
		{hconvert.Base36CharSet(), "Base36", "9FUQ0DSPUFM2T5124ZH3E11TSLYK2Z12BQ4F", data},
		{hconvert.Base58CharSet(), "Base58", "71vsZ3PsK9vVJkUU7G7iY1oJSVcoXxFg", data},
		{hconvert.Base62CharSet(), "Base62", "vLmUeMmpma4rHo3dGSdqrq39w4gK4D1", data},
		{hconvert.Base64CharSet(), "Base64", "SNFZ4mrze8AESIzRFVmd4iZqrvM3e7/", data},
		{hconvert.Base64URLCharSet(), "Base64URL", "SNFZ4mrze8AESIzRFVmd4iZqrvM3e7_", data},
		{hconvert.ASCII85CharSet(), "ASCII85", ";D6Y_Nm?3dN9r21pY8F'SiQu,n`^:", data},
		{hconvert.Z85CharSet(), "Z85", "qzlU.J)ui^Jo@hg{UnB6O&M#b[-Zp", data},
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

	if err := converter.EncodeTo([]byte{0x03, 0x04}, new(strings.Builder)); err == nil {
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

	// Test that we can properly decode the data using a converter with each of the standard
	// character sets.
	converter := new(hconvert.Converter)
	for _, test := range standardTests {
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

	// Test that we can't decode characters not in the character set.
	for _, test := range standardTests {
		converter.SetDecodeCharSet(test.charSet)
		if _, err := converter.Decode(invalidChar); err == nil {
			t.Error(test.setName, "- Decode invalid character passed")
		}
	}
}

func TestDecodeFrom(t *testing.T) {
	t.Parallel()

	// Test that we can properly read and decode data from the provided io.Reader using a converter
	// with each of the standard character sets
	converter := new(hconvert.Converter)
	for _, test := range standardTests {
		converter.SetDecodeCharSet(test.charSet)
		reader := strings.NewReader(test.decodeMe)
		b, err := converter.DecodeFrom(reader)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if !reflect.DeepEqual(b, test.encodeMe) {
			t.Error(test.setName, "- DecodeFrom failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", b)
		}
	}

	// Test that we can't decode characters not in the character set.
	for _, test := range standardTests {
		converter.SetDecodeCharSet(test.charSet)
		reader := strings.NewReader(invalidChar)
		if _, err := converter.DecodeFrom(reader); err == nil {
			t.Error(test.setName, "- DecodeFrom invalid character passed")
		}
	}
}

func TestDecodeWith(t *testing.T) {
	t.Parallel()

	// Test that we can properly decode the data using the standard character sets.
	for _, test := range standardTests {
		b, err := hconvert.DecodeWith(test.decodeMe, test.charSet)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if !reflect.DeepEqual(b, test.encodeMe) {
			t.Error(test.setName, "- DecodeWith failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", b)
		}
	}

	// Test that we can't decode characters not in the character set.
	for _, test := range standardTests {
		if _, err := hconvert.DecodeWith(invalidChar, test.charSet); err == nil {
			t.Error(test.setName, "- DecodeWith invalid character passed")
		}
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()

	// Test that we can properly encode the data using a converter with each of the standard
	// character sets.
	converter := new(hconvert.Converter)
	for _, test := range standardTests {
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

	// Test that we can properly encode data and write it to the provided io.Writer using a
	// converter with each of the standard character sets
	converter := new(hconvert.Converter)
	for _, test := range standardTests {
		converter.SetEncodeCharSet(test.charSet)
		writer := new(strings.Builder)
		if err := converter.EncodeTo(test.encodeMe, writer); err != nil {
			t.Error(test.setName, "-", err)
		}

		if writer.String() != test.decodeMe {
			t.Error(test.setName, "- EncodeTo failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", writer.String())
		}
	}
}

func TestEncodeWith(t *testing.T) {
	t.Parallel()

	// Test that we can properly encode the data using the standard character sets.
	for _, test := range standardTests {
		s, err := hconvert.EncodeWith(test.encodeMe, test.charSet)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if s != test.decodeMe {
			t.Error(test.setName, "- EncodeWith failed")
			t.Log("Expected:", test.decodeMe)
			t.Log("Received:", s)
		}
	}
}
