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
