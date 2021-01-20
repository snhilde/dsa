// This file provides some widely used character sets.

package hconvert

// ASCIICharSet creates a new CharSet with all 128 ASCII characters.
func ASCIICharSet() CharSet {
	c := make([]rune, 128)
	for i := range c {
		c[i] = rune(i)
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Base2CharSet creates a new CharSet with the base2 (binary) character set.
func Base2CharSet() CharSet {
	c := []rune{
		'0', '1',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Base8CharSet creates a new CharSet with the base8 (octal) character set.
func Base8CharSet() CharSet {
	c := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Base10CharSet creates a new CharSet with the base10 (decimal) character set.
func Base10CharSet() CharSet {
	c := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Base16CharSet creates a new CharSet with the base16 (hexadecimal) character set.
func Base16CharSet() CharSet {
	c := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Base32CharSet creates a new CharSet with the base32 character set.
func Base32CharSet() CharSet {
	c := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '2', '3', '4', '5', '6', '7',
	}

	charSet, _ := NewCharSet(c)
	charSet.SetPadding('=')

	return charSet
}

// Base36CharSet creates a new CharSet with the base36 character set.
func Base36CharSet() CharSet {
	c := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
		'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
		'W', 'X', 'Y', 'Z',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Base58CharSet creates a new CharSet with the base58 (bitcoin) character set.
func Base58CharSet() CharSet {
	c := []rune{
		'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G',
		'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y',
		'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Base62CharSet creates a new CharSet with the base62 (alphanumeric) character set.
func Base62CharSet() CharSet {
	c := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Base64CharSet creates a new CharSet with the base64 character set.
func Base64CharSet() CharSet {
	c := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/',
	}

	charSet, _ := NewCharSet(c)
	charSet.SetPadding('=')

	return charSet
}

// Base64URLCharSet creates a new CharSet with the base64url character set.
func Base64URLCharSet() CharSet {
	c := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '_',
	}

	charSet, _ := NewCharSet(c)
	charSet.SetPadding('=')

	return charSet
}

// ASCII85CharSet creates a new CharSet with the ASCII85 character set.
func ASCII85CharSet() CharSet {
	c := []rune{
		'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', '0',
		'1', '2', '3', '4', '5', '6', '7', '8', '9', ':', ';', '<', '=', '>', '?', '@',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '[', '\\', ']', '^', '_', '`',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}

// Z85CharSet creates a new CharSet with the Z85 character set.
func Z85CharSet() CharSet {
	c := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '.', '-',
		':', '+', '=', '^', '!', '/', '*', '?', '&', '<', '>', '(', ')', '[', ']', '{',
		'}', '@', '%', '$', '#',
	}

	charSet, _ := NewCharSet(c)

	return charSet
}
