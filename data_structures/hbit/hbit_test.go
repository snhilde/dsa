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
	var num  int
	var err  error
	var s    string

	// Test Len()
	num = buf.Len()
	if num != -1 {
		t.Error("Incorrect result from bad Buffer test for Len()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Cap()
	num = buf.Cap()
	if num != -1 {
		t.Error("Incorrect result from bad Buffer test for Cap()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Bits()
	num = buf.Bits()
	if num != -1 {
		t.Error("Incorrect result from bad Buffer test for Bits()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Reset()
	err = buf.Reset()
	if err == nil {
		t.Error("unexpectedly passed bad Buffer test for Reset()")
	}

	// Test String()
	s = buf.String()
	if s != "<nil>" {
		t.Error("Incorrect result from bad Buffer test for String()")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}

	// Test AddBit()
	err = buf.AddBit(1)
	if err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBit()")
	}

	// Test AddByte()
	err = buf.AddByte(0x0C)
	if err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddByte()")
	}

	// Test AddBytes()
	err = buf.AddBytes([]byte{0x0C, 0xFF})
	if err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBytes()")
	}
}

func TestBadStruct(t *testing.T) {
	// Make sure that a Buffer not created with New() is handled properly.
	var buf Buffer
	var num int
	var err error
	var s   string

	// Test Len()
	num = buf.Len()
	if num != -1 {
		t.Error("Incorrect result from bad Buffer test for Len()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Cap()
	num = buf.Cap()
	if num != -1 {
		t.Error("Incorrect result from bad Buffer test for Cap()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Bits()
	num = buf.Bits()
	if num != -1 {
		t.Error("Incorrect result from bad Buffer test for Bits()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Reset()
	err = buf.Reset()
	if err == nil {
		t.Error("unexpectedly passed bad Buffer test for Reset()")
	}

	// Test String()
	s = buf.String()
	if s != "<nil>" {
		t.Error("Incorrect result from bad Buffer test for String()")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}

	// Test AddBit()
	err = buf.AddBit(1)
	if err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBit()")
	}

	// Test AddByte()
	err = buf.AddByte(0x0C)
	if err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddByte()")
	}

	// Test AddBytes()
	err = buf.AddBytes([]byte{0x0C, 0xFF})
	if err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBytes()")
	}
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
