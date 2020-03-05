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
	if n, err := b.Reverse(10); n != 0 || err == nil {
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

func TestBit(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test an empty index.
	if val := b.Bit(0); val != false {
		t.Error("Incorrect result from Bit() test")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", val)
	}

	// Test an index out of range.
	if val := b.Bit(10); val != false {
		t.Error("Incorrect result from Bit() test")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", val)
	}

	b.AddByte(0x0F)
	if val := b.Bit(2); val != true {
		t.Error("Incorrect result from Bit() test")
		t.Log("\tExpected: true")
		t.Log("\tReceived:", val)
	}

	b.AddByte(0x0F)
	if val := b.Bit(6); val != false {
		t.Error("Incorrect result from Bit() test")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", val)
	}
}

func TestBits(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test before adding anything.
	if n := b.Bits(); n != 0 {
		t.Error("Incorrect result from Bits() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}

	// Test after adding some bits.
	b.AddBit(true)
	b.AddBit(false)
	b.AddBit(true)
	if n := b.Bits(); n != 3 {
		t.Error("Incorrect result from Bits() test")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", n)
	}

	// Test after adding a byte.
	b.AddByte(0x00)
	if n := b.Bits(); n != 11 {
		t.Error("Incorrect result from Bits() test")
		t.Log("\tExpected: 11")
		t.Log("\tReceived:", n)
	}

	// Test after adding some bytes.
	b.AddBytes([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06})
	if n := b.Bits(); n != 59 {
		t.Error("Incorrect result from Bits() test")
		t.Log("\tExpected: 59")
		t.Log("\tReceived:", n)
	}
}

func TestOffset(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test before adding anything.
	if n := b.Offset(); n != 0 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}

	// Test after adding some bits.
	b.AddBytes([]byte{0xFF, 0xFB})
	if n := b.Offset(); n != 0 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}

	// Test after moving the buffer forward some.
	b.Advance(5)
	if n := b.Offset(); n != 5 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 5")
		t.Log("\tReceived:", n)
	}

	// Test after moving the buffer back some.
	b.Reverse(3)
	if n := b.Offset(); n != 2 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 2")
		t.Log("\tReceived:", n)
	}

	// Test after wiping previous bits.
	b.Recalibrate()
	if n := b.Offset(); n != 0 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}
}

func TestCopy(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test copying an entire buffer.
	b.AddByte(0x0A)
	checkBits(t, b, 8)
	nb := b.Copy(b.Bits())
	checkBits(t, nb, 8)

	// Make sure that modifying one buffer does not modify the copy.
	b.AddBit(true)
	b.AddBit(false)
	checkBits(t, b, 10)
	checkBits(t, nb, 8)

	b.Advance(5)
	checkBits(t, b, 5)
	checkBits(t, nb, 8)

	nb.AddBytes([]byte{0xAA, 0xBB, 0xCC})
	checkBits(t, b, 5)
	checkBits(t, nb, 32)

	// Test copying part of a buffer.
	nnb := nb.Copy(10)
	checkBits(t, nb, 32)
	checkBits(t, nnb, 10)
}

func TestRecalibrate(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bytes, advancing the buffer, and setting it back.
	b.AddBytes([]byte{0x50, 0x60, 0x70})
	checkBits(t, b, 24)
	if n := b.Offset(); n != 0 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}

	b.Advance(20)
	checkBits(t, b, 4)
	if n := b.Offset(); n != 20 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 20")
		t.Log("\tReceived:", n)
	}

	if err := b.Recalibrate(); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 4)
	if n := b.Offset(); n != 0 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}

	// Test that we can no longer reverse the buffer.
	if n, err := b.Reverse(10); n != 0 || err != nil {
		t.Error("Reverse() results unexpected")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}
}


// HELPERS
func checkBits(t *testing.T, b *Buffer, want int) {
	if n := b.Bits(); n != want {
		t.Error("Incorrect number of bits")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", n)
	}
}

func checkString(t *testing.T, b *Buffer, want string) {
	if s := b.String(); s != want {
		t.Error("Incorrect string")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkDisplay(t *testing.T, b *Buffer, want string) {
	if s := b.Display(); s != want {
		t.Error("Incorrect display")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}
