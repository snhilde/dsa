package hbit

import (
	"testing"
)


// TESTS
func TestNew(t *testing.T) {
	buf := New()
	if buf == nil {
		t.Error("unexpectedly failed New() test")
	}

	if buf.buf == nil {
		t.Error("Internal buffer was not created")
	}

	if buf.index_begin != 0 {
		t.Error("index_begin is not initialized")
		t.Log("\tExpected 0")
		t.Log("\tActually,", buf.index_begin)
	}

	if buf.index_end != 0 {
		t.Error("index_end is not initialized")
		t.Log("\tExpected 0")
		t.Log("\tActually,", buf.index_end)
	}

	if buf.offset != 0 {
		t.Error("offset is not initialized")
		t.Log("\tExpected 0")
		t.Log("\tActually,", buf.offset)
	}

	if buf.end_bits != 0 {
		t.Error("end_bits is not initialized")
		t.Log("\tExpected 0")
		t.Log("\tActually,", buf.end_bits)
	}

	if buf.Len() != 1 {
		t.Error("Incorrect length of inital byte slice")
		t.Log("\tExpected 0")
		t.Log("\tReceived,", len(buf.buf))
	}

	if buf.Cap() != 1 {
		t.Error("Incorrect capacity of inital byte slice")
		t.Log("\tExpected 0")
		t.Log("\tReceived,", cap(buf.buf))
	}
}

func TestBadPtr(t *testing.T) {
	// Make sure that a Buffer not created with New() is handled properly.
	var buf *Buffer

	// Test Len().
	if num := buf.Len(); num != -1 {
		t.Error("Incorrect result from bad Buffer test for Len()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Cap().
	if num := buf.Cap(); num != -1 {
		t.Error("Incorrect result from bad Buffer test for Cap()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Bits().
	if num := buf.Bits(); num != -1 {
		t.Error("Incorrect result from bad Buffer test for Bits()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Reset().
	if err := buf.Reset(); err == nil {
		t.Error("unexpectedly passed bad Buffer test for Reset()")
	}

	// Test String().
	if s := buf.String(); s != "<nil>" {
		t.Error("Incorrect result from bad Buffer test for String()")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}

	// Test AddBit().
	if err := buf.AddBit(1); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBit()")
	}

	// Test AddByte().
	if err := buf.AddByte(0x0C); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddByte()")
	}

	// Test AddBytes().
	if err := buf.AddBytes([]byte{0x0C, 0xFF}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBytes()")
	}

	// Test Advance().
	if err := buf.Advance(10); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Advance()")
	}
}

func TestBadStruct(t *testing.T) {
	// Make sure that a Buffer not created with New() is handled properly.
	var buf Buffer

	// Test Len().
	if num := buf.Len(); num != -1 {
		t.Error("Incorrect result from bad Buffer test for Len()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Cap().
	if num := buf.Cap(); num != -1 {
		t.Error("Incorrect result from bad Buffer test for Cap()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Bits().
	if num := buf.Bits(); num != -1 {
		t.Error("Incorrect result from bad Buffer test for Bits()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Reset().
	if err := buf.Reset(); err == nil {
		t.Error("unexpectedly passed bad Buffer test for Reset()")
	}

	// Test String().
	if s := buf.String(); s != "<nil>" {
		t.Error("Incorrect result from bad Buffer test for String()")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}

	// Test AddBit().
	if err := buf.AddBit(1); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBit()")
	}

	// Test AddByte().
	if err := buf.AddByte(0x0C); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddByte()")
	}

	// Test AddBytes().
	if err := buf.AddBytes([]byte{0x0C, 0xFF}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBytes()")
	}

	// Test Advance().
	if err := buf.Advance(10); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Advance()")
	}
}

func TestAddBit(t *testing.T) {
	buf := New()
	checkString(t, buf, "<empty>")
	checkBits(t, buf, 0)

	// Add one bit.
	if err := buf.AddBit(1); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "1")
	checkBits(t, buf, 1)

	// Add enough bits to fill the first byte.
	if err := buf.AddBit(1); err != nil {
		t.Error(err)
	}
	if err := buf.AddBit(0); err != nil {
		t.Error(err)
	}
	if err := buf.AddBit(1); err != nil {
		t.Error(err)
	}
	if err := buf.AddBit(0); err != nil {
		t.Error(err)
	}
	if err := buf.AddBit(1); err != nil {
		t.Error(err)
	}
	if err := buf.AddBit(1); err != nil {
		t.Error(err)
	}
	if err := buf.AddBit(1); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "11010111")
	checkBits(t, buf, 8)

	// Overflow the first byte.
	if err := buf.AddBit(0); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "110101110")
	checkBits(t, buf, 9)
}

func TestAddByte(t *testing.T) {
	buf := New()
	checkString(t, buf, "<empty>")
	checkBits(t, buf, 0)

	// Add a byte within a byte boundary.
	if err := buf.AddByte(0xF0); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "00001111")
	checkBits(t, buf, 8)

	// Add a byte across a byte boundary.
	buf = New()
	checkString(t, buf, "<empty>")
	checkBits(t, buf, 0)

	buf.AddBit(1)
	buf.AddBit(1)
	buf.AddByte(0xF0)
	checkString(t, buf, "1100001111")
	checkBits(t, buf, 10)
}

func TestAddBytes(t *testing.T) {
	buf := New()
	checkString(t, buf, "<empty>")
	checkBits(t, buf, 0)

	// Add some bytes within a byte boundary.
	if err := buf.AddBytes([]byte{0xF0, 0x0F}); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "0000111111110000")
	checkBits(t, buf, 16)

	// Add some bytes across a byte boundary.
	buf = New()
	buf.AddBit(1)
	buf.AddBit(1)
	if err := buf.AddBytes([]byte{0xF0, 0x0F}); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "110000111111110000")
	checkBits(t, buf, 18)
}

func TestAdvance(t *testing.T) {
	buf := New()
	checkString(t, buf, "<empty>")
	checkBits(t, buf, 0)

	// Test a small step within a byte first.
	buf = New()
	buf.AddBit(1)
	buf.AddBit(1)
	checkString(t, buf, "11")
	checkBits(t, buf, 2)

	if err := buf.Advance(1); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "1")
	checkBits(t, buf, 1)

	// Test a small step that overflows a byte.
	buf = New()
	buf.AddByte(0x00)
	buf.AddBit(1)
	buf.AddBit(1)
	checkString(t, buf, "0000000011")
	checkBits(t, buf, 10)

	if err := buf.Advance(9); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "1")
	checkBits(t, buf, 1)

	// Test a large step now.
	buf = New()
	buf.AddBytes([]byte{0xF0, 0x0F})
	checkString(t, buf, "0000111111110000")
	checkBits(t, buf, 16)

	if err := buf.Advance(8); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "11110000")
	checkBits(t, buf, 8)

	// Test a large and a small step.
	buf = New()
	buf.AddBytes([]byte{0xF0, 0x0F})
	buf.AddBytes([]byte{0xFA, 0x0A})
	buf.AddBit(1)
	buf.AddBit(1)
	checkString(t, buf, "0000111111110000010111110101000011")
	checkBits(t, buf, 34)

	if err := buf.Advance(12); err != nil {
		t.Error(err)
	}
	checkString(t, buf, "0000010111110101000011")
	checkBits(t, buf, 22)
}


// HELPERS
func checkString(t *testing.T, buf *Buffer, want string) {
	s := buf.String()
	if s != want {
		t.Error("Incorrect string")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkBits(t *testing.T, buf *Buffer, want int) {
	n := buf.Bits()
	if n != want {
		t.Error("Incorrect number of bits")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", n)
	}
}
