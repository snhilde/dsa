package hbit

import (
	"testing"
)


// TESTS
func TestNew(t *testing.T) {
	b := New()
	if b == nil {
		t.Error("Unexpectedly failed New() test")
	}

	if b.n != nil {
		t.Error("Somehow created a node")
	}

	if num := b.Bits(); num != 0 {
		t.Error("Incorrect result from new buffer Bits() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", num)
	}

	if num := b.Offset(); num != 0 {
		t.Error("Incorrect result from new buffer Offset() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", num)
	}
}

func TestBadPtr(t *testing.T) {
	// Make sure that a Buffer not created with New() is handled properly.
	var b *Buffer

	// Test Bit().
	if val := b.Bit(8); val != false {
		t.Error("Incorrect result from bad Buffer test for Bit()")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", val)
	}

	// Test Bits().
	if num := b.Bits(); num != -1 {
		t.Error("Incorrect result from bad Buffer test for Bits()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Offset().
	if num := b.Offset(); num != -1 {
		t.Error("Incorrect result from bad Buffer test for Offset()")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", num)
	}

	// Test Copy().
	if nb := b.Copy(20); nb != nil {
		t.Error("Incorrect result from bad Buffer test for Copy()")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", nb)
	}

	// Test Recalibrate().
	if err := b.Recalibrate(); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Recalibrate()")
	}

	// Test Reset().
	b.Reset()

	// Test String().
	if s := b.String(); s != "<nil>" {
		t.Error("Incorrect result from bad Buffer test for String()")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}

	// Test Display().
	if s := b.Display(); s != "<nil>" {
		t.Error("Incorrect result from bad Buffer test for Display()")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}

	// Test AddBit().
	if err := b.AddBit(true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBit()")
	}

	// Test AddByte().
	if err := b.AddByte(0x0C); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddByte()")
	}

	// Test AddBytes().
	if err := b.AddBytes([]byte{0x0C, 0xFF}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for AddBytes()")
	}

	// Test Advance().
	if err := b.Advance(10); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Advance()")
	}

	// Test Reverse().
	if err := b.Reverse(10); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Reverse()")
	}

	// Test ANDBit().
	if err := b.ANDBit(10, true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or ANDBit()")
	}

	// Test ANDBytes().
	if err := b.ANDBytes([]byte{0xFF, 0xEE}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or ANDBytes()")
	}

	// Test ORBit().
	if err := b.ORBit(10, true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or ORBit()")
	}

	// Test ORBytes().
	if err := b.ORBytes([]byte{0xFF, 0xEE}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or ORBytes()")
	}

	// Test XORBit().
	if err := b.XORBit(10, true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or XORBit()")
	}

	// Test XORBytes().
	if err := b.XORBytes([]byte{0xFF, 0xEE}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or XORBytes()")
	}

	// Test NOTBit().
	if err := b.NOTBit(10); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or NOTBit()")
	}

	// Test NOTBytes().
	if err := b.NOTBytes(20); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or NOTBytes()")
	}

	// Test ShiftLeft().
	if err := b.ShiftLeft(5); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or ShiftLeft()")
	}

	// Test ShiftRight().
	if err := b.ShiftRight(6); err == nil {
		t.Error("Unexpectedly passed bad Buffer test or ShiftRight()")
	}
}
