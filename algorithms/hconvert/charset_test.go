package hconvert_test

import (
	"reflect"
	"testing"

	"github.com/snhilde/dsa/algorithms/hconvert"
)

// Test creating a new CharSet.
func TestNewCharSet(t *testing.T) {
	t.Parallel()

	set := []rune{'a', 'b', 'c'}
	charSet, err := hconvert.NewCharSet(set)
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

	// Test creating a character set that is longer than the max.
	longSet := make([]rune, hconvert.MaxNumChars+1)
	for i := range longSet {
		longSet[i] = rune(i)
	}

	if _, err := hconvert.NewCharSet(longSet); err == nil {
		t.Error("Exceeded maximum CharSet length")
	}

	// Test that duplicate characters are not allowed for ASCII-only character sets.
	dupASCIISet := []rune{'a', 'b', 'c', 'a', 'q', 'w', 'e'}
	if _, err := hconvert.NewCharSet(dupASCIISet); err == nil {
		t.Error("Duplicate ASCII characters allowed in new character set")
	}

	// Test that duplicate characters are not allowed for Unicode character sets.
	dupUncideSet := []rune{
		'☛', '⥺', '⮶', '$', 'ⴵ', '㙇', '✚', '🌗', '🌼', '❸', '⡇', '⽗', '🌎', '㏙', '🌗', '㤪',
		'䷷', '🌱', '☫', '🦉', 'ㆩ', '䷷', '🌱', '☫', '🦉', 'ㆩ',
	}
	if _, err := hconvert.NewCharSet(dupUncideSet); err == nil {
		t.Error("Duplicate Unicode characters allowed in new character set")
	}
}

// Test setting and getting the padding character.
func TestPadding(t *testing.T) {
	t.Parallel()

	charSet := hconvert.Base64CharSet()

	// Base64 defaults to a padding character of '='.
	if charSet.Padding() != "=" {
		t.Error("Unset padding character wrong:", charSet.Padding())
	}

	// Let's set it to '&' and try again.
	charSet.SetPadding("&")
	if charSet.Padding() != "&" {
		t.Error("New padding character wrong:", charSet.Padding())
	}

	// Let's try setting it to an emoji character.
	charSet.SetPadding("😁")
	if charSet.Padding() != "😁" {
		t.Error("Emoji padding character wrong:", charSet.Padding())
	}
}

// Test that the length is being reported correctly.
func TestLength(t *testing.T) {
	t.Parallel()

	// Test a regular set of 64 characters.
	set := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if charSet.Len() != len(set) {
		t.Error("Incorrect length (ASCII set)")
	}

	// Test a set of non-ASCII characters.
	set = []rune{
		'😌', '😗', '😜', 'ė', '😂', '😇', 'ɦ', 'ࢷ', '௵', 'ລ', '₽', '℅', 'Ⅷ', 'ↇ', 'ጧ', 'ᓶ',
		'ᡊ', 'ᨖ', 'ᮗ', '⏏', '⏹', '▙', '☪', '☶', 'Ͷ', 'ֆ',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if charSet.Len() != len(set) {
		t.Error("Incorrect length (unicode set)")
	}

	// Test a set of mixed ASCII and non-ASCII characters.
	set = []rune{
		'S', '㙇', 'r', '#', '🌼', 'R', '8', 'ㆩ', '☫', '4', '✚', ' ', '❸', '⡇', 'o', '⪴',
		'l', '☛', 'i', 'k', 's', '⥺', 'q', '⮶', '$', 'ⴵ', 'Y', 'f', '(', 'd', 'j', 'A',
		'3', 'J', '⽗', 'Q', '㏙', 'a', '㤪', '䷷', '🌎', '!', '🌱', '🦉', '🌗',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if charSet.Len() != len(set) {
		t.Error("Incorrect length (mixed set)")
	}
}

// Test that the correct character is returned.
func TestCharacters(t *testing.T) {
	t.Parallel()

	// Test a regular set of 16 characters.
	set := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(charSet.Characters(), set) {
		t.Error("ASCII character set not returned")
	}

	// Test a set of non-ASCII characters.
	set = []rune{
		'ᨖ', 'ᮗ', '₽', '😜', 'ↇ', '⏏', '⏹', '▙', '℅', 'Ⅷ', '😂', '😇', '😌', '😗', '☪', '☶',
		'௵', 'ລ', 'ጧ', 'ᓶ', 'ᡊ', 'ė', 'ɦ', 'Ͷ', 'ֆ', 'ࢷ',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(charSet.Characters(), set) {
		t.Error("Non-ASCII character set not returned")
	}

	// Test a set of mixed ASCII and non-ASCII characters.
	set = []rune{
		'3', 'J', '⽗', '.', '㏙', 'a', '㤪', '䷷', '🌎', '%', '🌱', '🦉', '🌗', 'o', '⪴',
		'S', '㙇', 'r', '#', '🌼', '\'', '8', 'ㆩ', '☫', '4', '✚', ' ', '❸', '⡇', 'A',
		'l', '☛', 'i', 'k', 's', '⥺', 'q', '⮶', '$', 'ⴵ', ']', 'f', '@', 'd', 'j',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(charSet.Characters(), set) {
		t.Error("Mixed character set not returned")
	}
}

// Test that the string representation of the character set is correct.
func TestString(t *testing.T) {
	t.Parallel()

	// Test a regular set of 36 characters.
	set := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
		'W', 'X', 'Y', 'Z',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if charSet.String() != string(set) {
		t.Error("ASCII character set string incorrect")
	}

	// Test a set of non-ASCII characters.
	set = []rune{
		'😂', '😇', '😌', '😗', '😜', 'ė', 'ɦ', 'Ͷ', 'ֆ', 'ࢷ', '௵', 'ລ', 'ጧ', 'ᓶ', 'ᡊ',
		'ᨖ', 'ᮗ', '₽', '℅', 'Ⅷ', 'ↇ', '⏏', '⏹', '▙', '☪', '☶',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if charSet.String() != string(set) {
		t.Error("Non-ASCII character set string incorrect")
	}

	// Test a set of mixed ASCII and non-ASCII characters.
	set = []rune{
		'l', '☛', 'i', '_', 's', '⥺', 'q', '⮶', '$', 'ⴵ', 'z', 'f', 'k', '7', 'j', 'A',
		'S', '㙇', 'r', '#', '🌼', 'p', '8', 'ㆩ', '☫', '4', '✚', ' ', '❸', '⡇', 'o', '⪴',
		'3', 'J', '⽗', 'm', '㏙', 'a', '㤪', '䷷', '🌎', 'd', '🌱', '🦉', '🌗',
	}
	if charSet, err := hconvert.NewCharSet(set); err != nil {
		t.Error(err)
	} else if charSet.String() != string(set) {
		t.Error("Mixed character set string incorrect")
	}
}

// Test that the constant character sets have not changed.
func TestConsts(t *testing.T) {
	t.Parallel()

	ascii := hconvert.ASCIICharSet()
	asciiTest := []rune{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
		48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93,
		94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112,
		113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	}
	if !reflect.DeepEqual(ascii.Characters(), asciiTest) {
		t.Error("ascii character set has changed")
	}
	if ascii.Padding() != "" {
		t.Error("ascii padding has changed")
	}

	binary := hconvert.BinaryCharSet()
	binaryTest := []rune{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
		48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93,
		94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112,
		113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
		131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148,
		149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166,
		167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184,
		185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202,
		203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220,
		221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238,
		239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255,
	}
	if !reflect.DeepEqual(binary.Characters(), binaryTest) {
		t.Error("binary character set has changed")
	}
	if binary.Padding() != "" {
		t.Error("binary padding has changed")
	}

	base2 := hconvert.Base2CharSet()
	base2Test := []rune{
		'0', '1',
	}
	if !reflect.DeepEqual(base2.Characters(), base2Test) {
		t.Error("base2 character set has changed")
	}
	if base2.Padding() != "" {
		t.Error("base2 padding has changed")
	}

	base4 := hconvert.Base4CharSet()
	base4Test := []rune{
		'0', '1', '2', '3',
	}
	if !reflect.DeepEqual(base4.Characters(), base4Test) {
		t.Error("base4 character set has changed")
	}
	if base4.Padding() != "" {
		t.Error("base4 padding has changed")
	}

	base8 := hconvert.Base8CharSet()
	base8Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7',
	}
	if !reflect.DeepEqual(base8.Characters(), base8Test) {
		t.Error("base8 character set has changed")
	}
	if base8.Padding() != "" {
		t.Error("base8 padding has changed")
	}

	base10 := hconvert.Base10CharSet()
	base10Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	if !reflect.DeepEqual(base10.Characters(), base10Test) {
		t.Error("base10 character set has changed")
	}
	if base10.Padding() != "" {
		t.Error("base10 padding has changed")
	}

	base16 := hconvert.Base16CharSet()
	base16Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
	}
	if !reflect.DeepEqual(base16.Characters(), base16Test) {
		t.Error("base16 character set has changed")
	}
	if base16.Padding() != "" {
		t.Error("base16 padding has changed")
	}

	base32 := hconvert.Base32CharSet()
	base32Test := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '2', '3', '4', '5', '6', '7',
	}
	if !reflect.DeepEqual(base32.Characters(), base32Test) {
		t.Error("base32 character set has changed")
	}
	if base32.Padding() != "=" {
		t.Error("base32 padding has changed")
	}

	base36 := hconvert.Base36CharSet()
	base36Test := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
		'W', 'X', 'Y', 'Z',
	}
	if !reflect.DeepEqual(base36.Characters(), base36Test) {
		t.Error("base36 character set has changed")
	}
	if base36.Padding() != "" {
		t.Error("base36 padding has changed")
	}

	base58 := hconvert.Base58CharSet()
	base58Test := []rune{
		'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G',
		'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y',
		'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
	if !reflect.DeepEqual(base58.Characters(), base58Test) {
		t.Error("base58 character set has changed")
	}
	if base58.Padding() != "" {
		t.Error("base58 padding has changed")
	}

	base62 := hconvert.Base62CharSet()
	base62Test := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	if !reflect.DeepEqual(base62.Characters(), base62Test) {
		t.Error("base62 character set has changed")
	}
	if base62.Padding() != "" {
		t.Error("base62 padding has changed")
	}

	base64 := hconvert.Base64CharSet()
	base64Test := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
	}
	if !reflect.DeepEqual(base64.Characters(), base64Test) {
		t.Error("base64 character set has changed")
	}
	if base64.Padding() != "=" {
		t.Error("base64 padding has changed")
	}

	base64url := hconvert.Base64URLCharSet()
	base64urlTest := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '_',
	}
	if !reflect.DeepEqual(base64url.Characters(), base64urlTest) {
		t.Error("base64url character set has changed")
	}
	if base64url.Padding() != "=" {
		t.Error("base64url padding has changed")
	}

	ascii85 := hconvert.ASCII85CharSet()
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
	if ascii85.Padding() != "" {
		t.Error("ascii85 padding has changed")
	}

	z85 := hconvert.Z85CharSet()
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
	if z85.Padding() != "" {
		t.Error("z85 padding has changed")
	}

	codePage437 := hconvert.CodePage437CharSet()
	codePage437Test := []rune{
		0, '☺', '☻', '♥', '♦', '♣', '♠', '•', '◘', '○', '◙', '♂', '♀', '♪', '♫', '☼',
		'►', '◄', '↕', '‼', '¶', '§', '▬', '↨', '↑', '↓', '→', '←', '∟', '↔', '▲', '▼',
		' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ':', ';', '<', '=', '>', '?',
		'@', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
		'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '[', '\\', ']', '^', '_',
		'`', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '{', '|', '}', '~', '⌂',
		'Ç', 'ü', 'é', 'â', 'ä', 'à', 'å', 'ç', 'ê', 'ë', 'è', 'ï', 'î', 'ì', 'Ä', 'Å',
		'É', 'æ', 'Æ', 'ô', 'ö', 'ò', 'û', 'ù', 'ÿ', 'Ö', 'Ü', '¢', '£', '¥', '₧', 'ƒ',
		'á', 'í', 'ó', 'ú', 'ñ', 'Ñ', 'ª', 'º', '¿', '⌐', '¬', '½', '¼', '¡', '«', '»',
		'░', '▒', '▓', '│', '┤', '╡', '╢', '╖', '╕', '╣', '║', '╗', '╝', '╜', '╛', '┐',
		'└', '┴', '┬', '├', '─', '┼', '╞', '╟', '╚', '╔', '╩', '╦', '╠', '═', '╬', '╧',
		'╨', '╤', '╥', '╙', '╘', '╒', '╓', '╫', '╪', '┘', '┌', '█', '▄', '▌', '▐', '▀',
		'α', 'ß', 'Γ', 'π', 'Σ', 'σ', 'µ', 'τ', 'Φ', 'Θ', 'Ω', 'δ', '∞', 'φ', 'ε', '∩',
		'≡', '±', '≥', '≤', '⌠', '⌡', '÷', '≈', '°', '∙', '·', '√', 'ⁿ', '²', '■', '\u00A0',
	}
	if !reflect.DeepEqual(codePage437.Characters(), codePage437Test) {
		t.Error("codePage437 character set has changed")
	}
	if codePage437.Padding() != "" {
		t.Error("codePage437 padding has changed")
	}
}
