package hconvert

import (
	"reflect"
	"testing"
)

// Test creating a new CharSet.
func TestNewCharSet(t *testing.T) {
	set := []rune{'a', 'b', 'c'}
	charSet, err := NewCharSet(set)
	if err != nil {
		t.Error(err)
		return
	}

	// Test that the new CharSet is equivalent to the specified set.
	newSet := charSet.Characters()
	if len(set) != len(newSet) {
		t.Error("Lengths differ:", len(set), "!=", len(newSet))
	}

	for i, v := range set {
		if v != newSet[i] {
			t.Error("Difference at index", i, ":", v, "!=", newSet[i])
		}
	}

	// Test that we can change the source set and not affect the new CharSet.
	origSet := make([]rune, len(set))
	copy(origSet, set)
	set[0] = 'd'
	newSet = charSet.Characters()
	for i, v := range origSet {
		if v != newSet[i] {
			t.Error("Difference at index", i, ":", v, "!=", newSet[i])
		}
	}
}

// Test creating a character set that is longer than the max.
func TestCharSetLength(t *testing.T) {
	set := make([]rune, maxChars + 1)
	for i := range set {
		set[i] = rune(i)
	}

	_, err := NewCharSet(set)
	if err == nil {
		t.Error("Exceeded maximum CharSet length")
	}
}

// Test setting and getting the padding character.
func TestPadding(t *testing.T) {
	charSet := Base64CharSet()

	// Base64 defaults to a padding character of '='.
	if charSet.Padding() != '=' {
		t.Error("Unset padding character wrong:", string(charSet.Padding()))
	}

	// Let's set it to '&' and try again.
	charSet.SetPadding('&')
	if charSet.Padding() != '&' {
		t.Error("New padding character wrong:", string(charSet.Padding()))
	}

	// Let's try setting it to an emoji character.
	charSet.SetPadding('😁')
	if charSet.Padding() != '😁' {
		t.Error("Emoji padding character wrong:", string(charSet.Padding()))
	}
}

// Test that the length is being reported correctly.
func TestLength(t *testing.T) {
	// Test a regular set of 36 characters.
	set := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
		'W', 'X', 'Y', 'Z',
	}
	charSet, _ := NewCharSet(set)
	if charSet.Length() != len(set) {
		t.Error("Incorrect length (ASCII set)")
	}

	// Test a set of non-ASCII characters.
	set = []rune{
		'😂', '😇', '😌', '😗', '😜', 'ė', 'ɦ', 'Ͷ', 'ֆ', 'ࢷ', '௵', 'ລ', 'ጧ', 'ᓶ', 'ᡊ',
		'ᨖ', 'ᮗ', '₽', '℅', 'Ⅷ', 'ↇ', '⏏', '⏹', '▙', '☪', '☶',
	}
	charSet, _ = NewCharSet(set)
	if charSet.Length() != len(set) {
		t.Error("Incorrect length (unicode set)")
	}

	// Test a set of mixed ASCII and non-ASCII characters.
	set = []rune{
		'l', '☛', 'i', 'k', 's', '⥺', 'q', '⮶', '$', 'ⴵ', 'i', 'f', 'k', 'd', 'j', 'A',
		'S', '㙇', 'r', '#', '🌼', 'r', '8', 'ㆩ', '☫', '4', '✚', ' ', '❸', '⡇', 'o', '⪴',
		'3', 'J', '⽗', 'f', '㏙', 'a', '㤪', '䷷', '🌎', 'd', '🌱', '🦉', '🌗',
	}
	charSet, _ = NewCharSet(set)
	if charSet.Length() != len(set) {
		t.Error("Incorrect length (mixed set)")
	}
}

// Test that the correct character is returned.
func TestCharacters(t *testing.T) {
	// Test a regular set of 36 characters.
	set := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
		'W', 'X', 'Y', 'Z',
	}
	charSet, _ := NewCharSet(set)
	if !reflect.DeepEqual(charSet.Characters(), set) {
		t.Error("ASCII character set not returned")
	}

	// Test a set of non-ASCII characters.
	set = []rune{
		'😂', '😇', '😌', '😗', '😜', 'ė', 'ɦ', 'Ͷ', 'ֆ', 'ࢷ', '௵', 'ລ', 'ጧ', 'ᓶ', 'ᡊ',
		'ᨖ', 'ᮗ', '₽', '℅', 'Ⅷ', 'ↇ', '⏏', '⏹', '▙', '☪', '☶',
	}
	charSet, _ = NewCharSet(set)
	if !reflect.DeepEqual(charSet.Characters(), set) {
		t.Error("Non-ASCII character set not returned")
	}

	// Test a set of mixed ASCII and non-ASCII characters.
	set = []rune{
		'l', '☛', 'i', 'k', 's', '⥺', 'q', '⮶', '$', 'ⴵ', 'i', 'f', 'k', 'd', 'j', 'A',
		'S', '㙇', 'r', '#', '🌼', 'r', '8', 'ㆩ', '☫', '4', '✚', ' ', '❸', '⡇', 'o', '⪴',
		'3', 'J', '⽗', 'f', '㏙', 'a', '㤪', '䷷', '🌎', 'd', '🌱', '🦉', '🌗',
	}
	charSet, _ = NewCharSet(set)
	if !reflect.DeepEqual(charSet.Characters(), set) {
		t.Error("Mixed character set not returned")
	}
}

// Test that the string representation of the character set is correct.
func TestString(t *testing.T) {
	// Test a regular set of 36 characters.
	set := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
		'W', 'X', 'Y', 'Z',
	}
	charSet, _ := NewCharSet(set)
	if charSet.String() != string(set) {
		t.Error("ASCII character set string incorrect")
	}

	// Test a set of non-ASCII characters.
	set = []rune{
		'😂', '😇', '😌', '😗', '😜', 'ė', 'ɦ', 'Ͷ', 'ֆ', 'ࢷ', '௵', 'ລ', 'ጧ', 'ᓶ', 'ᡊ',
		'ᨖ', 'ᮗ', '₽', '℅', 'Ⅷ', 'ↇ', '⏏', '⏹', '▙', '☪', '☶',
	}
	charSet, _ = NewCharSet(set)
	if charSet.String() != string(set) {
		t.Error("Non-ASCII character set string incorrect")
	}

	// Test a set of mixed ASCII and non-ASCII characters.
	set = []rune{
		'l', '☛', 'i', 'k', 's', '⥺', 'q', '⮶', '$', 'ⴵ', 'i', 'f', 'k', 'd', 'j', 'A',
		'S', '㙇', 'r', '#', '🌼', 'r', '8', 'ㆩ', '☫', '4', '✚', ' ', '❸', '⡇', 'o', '⪴',
		'3', 'J', '⽗', 'f', '㏙', 'a', '㤪', '䷷', '🌎', 'd', '🌱', '🦉', '🌗',
	}
	charSet, _ = NewCharSet(set)
	if charSet.String() != string(set) {
		t.Error("Mixed character set string incorrect")
	}
}

// Test that the constant character sets have not changed.
func TestConsts(t *testing.T) {
	ascii := AsciiCharSet()
	asciiTest := []rune{
		  0,   1,   2,   3,   4,   5,   6,   7,   8,   9,  10,  11,  12,  13,  14,  15,  16,  17,  18,  19,
		 20,  21,  22,  23,  24,  25,  26,  27,  28,  29,  30,  31,  32,  33,  34,  35,  36,  37,  38,  39,
		 40,  41,  42,  43,  44,  45,  46,  47,  48,  49,  50,  51,  52,  53,  54,  55,  56,  57,  58,  59,
		 60,  61,  62,  63,  64,  65,  66,  67,  68,  69,  70,  71,  72,  73,  74,  75,  76,  77,  78,  79,
		 80,  81,  82,  83,  84,  85,  86,  87,  88,  89,  90,  91,  92,  93,  94,  95,  96,  97,  98,  99,
		100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
		120, 121, 122, 123, 124, 125, 126, 127,
	}
	if !reflect.DeepEqual(ascii.Characters(), asciiTest) {
		t.Error("ascii character set has changed")
	}

	base2 := Base2CharSet()
	base2Test := []rune{
		'0', '1',
	}
	if !reflect.DeepEqual(base2.Characters(), base2Test) {
		t.Error("base2 character set has changed")
	}

	base8 := Base8CharSet()
	base8Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7',
	}
	if !reflect.DeepEqual(base8.Characters(), base8Test) {
		t.Error("base8 character set has changed")
	}

	base10 := Base10CharSet()
	base10Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	if !reflect.DeepEqual(base10.Characters(), base10Test) {
		t.Error("base10 character set has changed")
	}

	base16 := Base16CharSet()
	base16Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
	}
	if !reflect.DeepEqual(base16.Characters(), base16Test) {
		t.Error("base16 character set has changed")
	}

	base32 := Base32CharSet()
	base32Test := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '2', '3', '4', '5', '6', '7',
	}
	if !reflect.DeepEqual(base32.Characters(), base32Test) {
		t.Error("base32 character set has changed")
	}

	base36 := Base36CharSet()
	base36Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
		'W', 'X', 'Y', 'Z',
	}
	if !reflect.DeepEqual(base36.Characters(), base36Test) {
		t.Error("base36 character set has changed")
	}

	base58 := Base58CharSet()
	base58Test := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R',
		'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
		'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
		'y', 'z', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	if !reflect.DeepEqual(base58.Characters(), base58Test) {
		t.Error("base58 character set has changed")
	}

	base62 := Base62CharSet()
	base62Test := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	if !reflect.DeepEqual(base62.Characters(), base62Test) {
		t.Error("base62 character set has changed")
	}

	base64 := Base64CharSet()
	base64Test := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
	}
	if !reflect.DeepEqual(base64.Characters(), base64Test) {
		t.Error("base64 character set has changed")
	}

	ascii85 := ASCII85CharSet()
	ascii85Test := []rune{
		'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', '0',
		'1', '2', '3', '4', '5', '6', '7', '8', '9', ':', ';', '<', '=', '>', '?', '@',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '[', '\\', ']', '^', '_', '`',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u',
	}
	if !reflect.DeepEqual(ascii85.Characters(), ascii85Test) {
		t.Error("aSCII85 character set has changed")
	}

	z85 := Z85CharSet()
	z85Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '.', '-',
		':', '+', '=', '^', '!', '/', '*', '?', '&', '<', '>', '(', ')', '[', ']', '{',
		'}', '@', '%', '$', '#',
	}
	if !reflect.DeepEqual(z85.Characters(), z85Test) {
		t.Error("z85 character set has changed")
	}
}
