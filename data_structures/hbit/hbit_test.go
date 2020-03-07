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

	if b.node != nil {
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
		t.Log("\tExpected: <nil>")
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
	if n, err := b.Advance(10); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Advance()")
	}

	// Test Rewind().
	if n, err := b.Rewind(10); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Rewind()")
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

func TestInvalidArgs(t *testing.T) {
	// TODO
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
	b.Rewind(3)
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
	if nb == nil {
		t.Error("Failed to create copy of buffer")
	}
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
	if nnb == nil {
		t.Error("Failed to create copy of buffer")
	}
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
	if n, err := b.Rewind(10); n != 0 || err != nil {
		t.Error(err)
		t.Error("Rewind() results unexpected")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}
}

func TestReset(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bits and then resetting the buffer.
	b.AddBit(false)
	b.AddBit(false)
	checkBits(t, b, 2)
	if err := b.Reset(); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)

	// Test adding some bits, advancing, and then resetting the buffer.
	b.AddByte(0xFF)
	checkBits(t, b, 8)
	b.Advance(5)
	checkBits(t, b, 3)
	if err := b.Reset(); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)

	// Test adding some bits, advancing, reversing, and then resetting the buffer.
	b.AddByte(0xEE)
	checkBits(t, b, 8)
	b.Advance(4)
	checkBits(t, b, 4)
	b.Rewind(1)
	checkBits(t, b, 5)
	if err := b.Reset(); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)
}

func TestStringAndDisplay(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddBit(true)
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	b.AddBit(false)
	checkBits(t, b, 2)
	checkString(t, b, "10")
	checkDisplay(t, b, "10")

	b.AddByte(0xFF)
	checkBits(t, b, 10)
	checkString(t, b, "1011111111")
	checkDisplay(t, b, "1011 1111  11")

	// Make sure that a single byte does not have any spaces around it.
	b = New()
	b.AddByte(0x00)
	checkBits(t, b, 8)
	checkString(t, b, "00000000")
	checkDisplay(t, b, "0000 0000")

	// Test out multiple bytes and bit order.
	b = New()
	b.AddBytes([]byte{0x1A, 0x2B, 0x3C})
	checkBits(t, b, 24)
	checkString(t, b, "010110001101010000111100")
	checkDisplay(t, b, "0101 1000  1101 0100  0011 1100")

	// Test out advancing.
	b.Advance(10)
	checkBits(t, b, 14)
	checkString(t, b, "01010000111100")
	checkDisplay(t, b, "0101 0000  1111 00")

	// Test out reversing.
	b.Rewind(5)
	checkBits(t, b, 19)
	checkString(t, b, "0001101010000111100")
	checkDisplay(t, b, "0001 1010  1000 0111  100")

	// Make sure recalibrating doesn't affect anything.
	b.Recalibrate()
	checkBits(t, b, 19)
	checkString(t, b, "0001101010000111100")
	checkDisplay(t, b, "0001 1010  1000 0111  100")

	// Test out resetting the buffer.
	b.Reset()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")
}

func TestAddBit(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bits.
	if err := b.AddBit(true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	if err := b.AddBit(true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 2)
	checkString(t, b, "11")
	checkDisplay(t, b, "11")

	if err := b.AddBit(false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 3)
	checkString(t, b, "110")
	checkDisplay(t, b, "110")

	// Test advancing and adding a bit.
	b.Advance(2)
	if err := b.AddBit(true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 2)
	checkString(t, b, "01")
	checkDisplay(t, b, "01")

	// Test reversing and adding a bit.
	b.Rewind(1)
	if err := b.AddBit(false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 4)
	checkString(t, b, "1010")
	checkDisplay(t, b, "1010")

	// Test resetting and adding a bit.
	b.Reset()
	if err := b.AddBit(false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "0")
	checkDisplay(t, b, "0")
}

func TestAddByte(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bytes.
	if err := b.AddByte(0xF0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "00001111")
	checkDisplay(t, b, "0000 1111")

	if err := b.AddByte(0x88); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 16)
	checkString(t, b, "0000111100010001")
	checkDisplay(t, b, "0000 1111  0001 0001")

	// Test advancing and adding a byte.
	b.Advance(10)
	if err := b.AddByte(0x14); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 14)
	checkString(t, b, "01000100101000")
	checkDisplay(t, b, "0100 0100  1010 00")

	// Test reversing and adding a byte.
	b.Rewind(3)
	if err := b.AddByte(0xA0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 25)
	checkString(t, b, "1000100010010100000000101")
	checkDisplay(t, b, "1000 1000  1001 0100  0000 0010  1")

	// Test resetting and adding a byte.
	b.Reset()
	if err := b.AddByte(0x44); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "00100010")
	checkDisplay(t, b, "0010 0010")
}

func TestAddBytes(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bytes.
	if err := b.AddBytes([]byte{0x54, 1}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 16)
	checkString(t, b, "0010101010000000")
	checkDisplay(t, b, "0010 1010  1000 0000")

	if err := b.AddBytes([]byte{0xAA, 0xBB, 0xCC}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 40)
	checkString(t, b, "0010101010000000010101011101110100110011")
	checkDisplay(t, b, "0010 1010  1000 0000  0101 0101  1101 1101  0011 0011")

	// Test advancing and adding some bytes.
	b.Advance(30)
	if err := b.AddBytes([]byte{0x01, 0x02}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 26)
	checkString(t, b, "01001100111000000001000000")
	checkDisplay(t, b, "0100 1100  1110 0000  0001 0000  00")

	// Test reversing and adding a byte.
	b.Rewind(5)
	if err := b.AddBytes([]byte{0x08}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 39)
	checkString(t, b, "101110100110011100000000100000000010000")
	checkDisplay(t, b, "1011 1010  0110 0111  0000 0000  1000 0000  0010 000")

	// Test resetting and adding some bytes.
	b.Reset()
	if err := b.AddBytes([]byte{0x98, 0x76}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 16)
	checkString(t, b, "0001100101101110")
	checkDisplay(t, b, "0001 1001  0110 1110")
}

func TestRemoveBit(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddByte(0x50)
	checkBits(t, b, 8)
	checkString(t, b, "00001010")
	checkDisplay(t, b, "0000 1010")

	// Test removing a bit.
	if err := b.RemoveBit(4); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 7)
	checkString(t, b, "0000010")
	checkDisplay(t, b, "0000 010")

	// Test removing the bit at the beginning of the buffer.
	b.Reset()
	b.AddByte(0x51)
	if err := b.RemoveBit(0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 7)
	checkString(t, b, "0001010")
	checkDisplay(t, b, "0001 010")

	// Test removing the bit at the end of the buffer.
	b.Reset()
	b.AddByte(0x51)
	if err := b.RemoveBit(b.Bits()-1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 7)
	checkString(t, b, "1000101")
	checkDisplay(t, b, "1000 101")

	// Test removing a bit when there's only 1 bit in the buffer.
	b.Reset()
	b.AddBit(true)
	if err := b.RemoveBit(0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test removing the first bit when there are only 2 bits in the buffer.
	b.Reset()
	b.AddBit(true)
	b.AddBit(false)
	if err := b.RemoveBit(0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "0")
	checkDisplay(t, b, "0")

	// Test removing the second bit when there are only 2 bits in the buffer.
	b.Reset()
	b.AddBit(true)
	b.AddBit(false)
	if err := b.RemoveBit(1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	// Test advancing and removing the first bit.
	b.Reset()
	b.AddByte(0x51)
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	if err := b.RemoveBit(0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 4)
	checkString(t, b, "1010")
	checkDisplay(t, b, "1010")

	b.Rewind(3)
	checkBits(t, b, 7)
	checkString(t, b, "1001010")
	checkDisplay(t, b, "1001 010")

	// Test advancing and removing the last bit.
	b.Reset()
	b.AddByte(0x51)
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	if err := b.RemoveBit(b.Bits()-1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 4)
	checkString(t, b, "0101")
	checkDisplay(t, b, "0101")

	b.Rewind(3)
	checkBits(t, b, 7)
	checkString(t, b, "1000101")
	checkDisplay(t, b, "1000 101")
}

func TestRemoveBits(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddByte(0x50)
	checkBits(t, b, 8)
	checkString(t, b, "00001010")
	checkDisplay(t, b, "0000 1010")

	// Test removing a bit.
	if err := b.RemoveBits(1, 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 7)
	checkString(t, b, "0001010")
	checkDisplay(t, b, "0001 010")

	// Test removing some bits at the beginning of the buffer.
	b.Reset()
	b.AddByte(0x51)
	if err := b.RemoveBits(0, 3); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	// Test removing some bits at the end of the buffer.
	b.Reset()
	b.AddByte(0x51)
	if err := b.RemoveBits(b.Bits()-3, 3); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 5)
	checkString(t, b, "10001")
	checkDisplay(t, b, "1000 1")

	// Test removing a bit when there's only 1 bit in the buffer.
	b.Reset()
	b.AddBit(true)
	if err := b.RemoveBits(0, 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test removing the first bit when there are only 2 bits in the buffer.
	b.Reset()
	b.AddBit(true)
	b.AddBit(false)
	if err := b.RemoveBits(0, 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "0")
	checkDisplay(t, b, "0")

	// Test removing the second bit when there are only 2 bits in the buffer.
	b.Reset()
	b.AddBit(true)
	b.AddBit(false)
	if err := b.RemoveBits(1, 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	// Test advancing and removing some bits.
	b.Reset()
	b.AddByte(0x51)
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	if err := b.RemoveBits(0, 2); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 3)
	checkString(t, b, "010")
	checkDisplay(t, b, "010")

	b.Rewind(3)
	checkBits(t, b, 6)
	checkString(t, b, "100010")
	checkDisplay(t, b, "1000 10")

	// Test advancing and removing some bits at the end.
	b.Reset()
	b.AddByte(0x51)
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	if err := b.RemoveBits(b.Bits()-3, 2); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 3)
	checkString(t, b, "010")
	checkDisplay(t, b, "010")

	b.Rewind(3)
	checkBits(t, b, 6)
	checkString(t, b, "100010")
	checkDisplay(t, b, "1000 10")

	// Test advancing, removing all bits, and reversing.
	b.Reset()
	b.AddByte(0x51)
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")
	if n := b.Offset(); n != 3 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", n)
	}

	if err := b.RemoveBits(0, b.Bits()); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.Rewind(3)
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")
	if n := b.Offset(); n != 0 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}

	// Test removing more bits than currently exist in the buffer.
	b.Reset()
	b.AddByte(0x51)
	if err := b.RemoveBits(2, 10); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 2)
	checkString(t, b, "10")
	checkDisplay(t, b, "10")
}

func testSetBit(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddByte(0x50)
	checkBits(t, b, 8)
	checkString(t, b, "00001010")
	checkDisplay(t, b, "0000 1010")

	// Test setting the first bit.
	if err := b.SetBit(0, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10001010")
	checkDisplay(t, b, "1000 1010")

	// Test setting the last bit.
	b.Reset()
	b.AddByte(0x50)
	if err := b.SetBit(0, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "00001011")
	checkDisplay(t, b, "0000 1011")

	// Test setting a middle bit.
	b.Reset()
	b.AddByte(0x50)
	if err := b.SetBit(4, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "00000011")
	checkDisplay(t, b, "0000 0011")

	// Test advancing and setting a bit.
	b.Reset()
	b.AddByte(0x51)
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	if err := b.SetBit(0, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 4)
	checkString(t, b, "0010")
	checkDisplay(t, b, "0010")

	b.Rewind(3)
	checkBits(t, b, 7)
	checkString(t, b, "1000010")
	checkDisplay(t, b, "1000 010")
}

func testSetBytes(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddByte(0x50)
	checkBits(t, b, 8)
	checkString(t, b, "00001010")
	checkDisplay(t, b, "0000 1010")

	// Test setting the first byte.
	if err := b.SetBytes(0, []byte{0xFF}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11111111")
	checkDisplay(t, b, "1111 1111")

	// Test setting too many bits.
	b.Reset()
	b.AddByte(0x50)
	if err := b.SetBytes(0, []byte{0x4C, 0xFF}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "00110010")
	checkDisplay(t, b, "0011 0010")

	// Test advancing and setting a byte.
	b.Reset()
	b.AddByte(0x51)
	checkString(t, b, "10001010")
	checkDisplay(t, b, "1000 1010")
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	if err := b.SetBytes(0, []byte{0x00}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 4)
	checkString(t, b, "0000")
	checkDisplay(t, b, "0000")

	b.Rewind(5)
	checkBits(t, b, 8)
	checkString(t, b, "10000000")
	checkDisplay(t, b, "1000 0000")
}

func TestAdvance(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test a normal advance.
	b.AddByte(0xFF)
	if n, err := b.Advance(5); n != 5 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Advance() test")
		t.Log("\tExpected: 5, <nil>")
		t.Log("\tReceived:", n, err)
	}
	checkBits(t, b, 3)
	checkString(t, b, "111")
	checkDisplay(t, b, "111")

	// Test advancing past the buffer.
	if n, err := b.Advance(10); n != 2 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Advance() test")
		t.Log("\tExpected: 2, <nil>")
		t.Log("\tReceived:", n, err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	// Rewind the buffer to make sure that we didn't overrun the end.
	if n, err := b.Rewind(1); n != 1 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Rewind() test")
		t.Log("\tExpected: 1, <nil>")
		t.Log("\tReceived:", n, err)
	}
	checkBits(t, b, 2)
	checkString(t, b, "11")
	checkDisplay(t, b, "11")
}

func TestRewind(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test advancing and reversing back to the starting point.
	b.AddByte(0xFF)
	b.Advance(5)
	checkBits(t, b, 3)
	checkString(t, b, "111")
	checkDisplay(t, b, "111")

	if n, err := b.Rewind(5); n != 5 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Rewind() test")
		t.Log("\tExpected: 5, <nil>")
		t.Log("\tReceived:", n, err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11111111")
	checkDisplay(t, b, "1111 1111")

	// Test advancing and reversing past the starting point.
	b.Advance(4)
	checkBits(t, b, 4)
	checkString(t, b, "1111")
	checkDisplay(t, b, "1111")

	if n, err := b.Rewind(6); n != 4 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Rewind() test")
		t.Log("\tExpected: 4, <nil>")
		t.Log("\tReceived:", n, err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11111111")
	checkDisplay(t, b, "1111 1111")
}

func TestMerge(t *testing.T) {
	// TODO: test merging the same buffer onto itself.
}

func TestANDBit(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddByte(0x55)
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")

	// Test 0 & 0
	if err := b.ANDBit(1, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")

	// Test 0 & 1
	if err := b.ANDBit(1, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")

	// Test 1 & 0
	if err := b.ANDBit(2, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10001010")
	checkDisplay(t, b, "1000 1010")

	// Test 1 & 1
	if err := b.ANDBit(0, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10001010")
	checkDisplay(t, b, "1000 1010")
}

func TestORBit(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddByte(0x55)
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")

	// Test 0 | 0
	if err := b.ORBit(1, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")

	// Test 0 | 1
	if err := b.ORBit(1, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11101010")
	checkDisplay(t, b, "1110 1010")

	// Test 1 | 0
	if err := b.ORBit(1, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11101010")
	checkDisplay(t, b, "1110 1010")

	// Test 1 | 1
	if err := b.ORBit(1, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11101010")
	checkDisplay(t, b, "1110 1010")
}

func TestXORBit(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddByte(0x55)
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")

	// Test 0 ^ 0
	if err := b.XORBit(1, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")

	// Test 0 ^ 1
	if err := b.XORBit(1, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11101010")
	checkDisplay(t, b, "1110 1010")

	// Test 1 ^ 0
	if err := b.XORBit(1, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11101010")
	checkDisplay(t, b, "1110 1010")

	// Test 1 ^ 1
	if err := b.XORBit(1, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")
}

func TestNOTBit(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddByte(0x55)
	checkBits(t, b, 8)
	checkString(t, b, "10101010")
	checkDisplay(t, b, "1010 1010")

	// Test ~0
	if err := b.NOTBit(1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11101010")
	checkDisplay(t, b, "1110 1010")

	// Test ~1
	if err := b.XORBit(6, true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11101000")
	checkDisplay(t, b, "1110 1000")
}

func TestANDBytes(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddBytes([]byte{0xFF, 0x00, 0xFF, 0x00})
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if err := b.ANDBytes([]byte{0x0F, 0xFF}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "11110000000000001111111100000000")
	checkDisplay(t, b, "1111 0000  0000 0000  1111 1111  0000 0000")

	// Test advancing, ANDing, and reversing.
	b.Advance(5)
	if err := b.ANDBytes([]byte{0x0C, 0x0F}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 27)
	checkString(t, b, "000000000001000011100000000")
	checkDisplay(t, b, "0000 0000  0001 0000  1110 0000  000")

	b.Rewind(5)
	checkBits(t, b, 32)
	checkString(t, b, "11110000000000001000011100000000")
	checkDisplay(t, b, "1111 0000  0000 0000  1000 0111  0000 0000")
}

func TestORBytes(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddBytes([]byte{0xFF, 0x00, 0xFF, 0x00})
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if err := b.ORBytes([]byte{0x0F, 0x0F}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "11111111111100001111111100000000")
	checkDisplay(t, b, "1111 1111  1111 0000  1111 1111  0000 0000")

	// Test advancing, ORing, and reversing.
	b.Advance(5)
	if err := b.ORBytes([]byte{0x0C, 0x0F}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 27)
	checkString(t, b, "111111101111111111100000000")
	checkDisplay(t, b, "1111 1110  1111 1111  1110 0000  000")

	b.Rewind(5)
	checkBits(t, b, 32)
	checkString(t, b, "11111111111101111111111100000000")
	checkDisplay(t, b, "1111 1111  1111 0111  1111 1111  0000 0000")
}

func TestXORBytes(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddBytes([]byte{0xFF, 0x00, 0xFF, 0x00})
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if err := b.XORBytes([]byte{0x0D, 0xFF}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "01001111111111111111111100000000")
	checkDisplay(t, b, "0100 1111  1111 1111  1111 1111  0000 0000")

	// Test advancing, XORing, and reversing.
	b.Advance(5)
	if err := b.XORBytes([]byte{0x0C, 0x0F}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 27)
	checkString(t, b, "110011110000111111100000000")
	checkDisplay(t, b, "1100 1111  0000 1111  1110 0000  000")

	b.Rewind(5)
	checkBits(t, b, 32)
	checkString(t, b, "01001110011110000111111100000000")
	checkDisplay(t, b, "0100 1110  0111 1000  0111 1111  0000 0000")
}

func TestNOTBytes(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddBytes([]byte{0xFF, 0x00, 0xFF, 0x00})
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if err := b.NOTBytes(2); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "00000000111111111111111100000000")
	checkDisplay(t, b, "0000 0000  1111 1111  1111 1111  0000 0000")

	// Test advancing, NOTing, and reversing.
	b.Advance(5)
	if err := b.NOTBytes(1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 27)
	checkString(t, b, "111000001111111111100000000")
	checkDisplay(t, b, "1110 0000  1111 1111  1110 0000  000")

	b.Rewind(5)
	checkBits(t, b, 32)
	checkString(t, b, "00000111000001111111111100000000")
	checkDisplay(t, b, "0000 0111  0000 0111  1111 1111  0000 0000")

	if err := b.NOTBytes(10); err != nil {
		t.Error(err)
	}
	// Test overrunning the buffer.
	checkBits(t, b, 32)
	checkString(t, b, "11111000111110000000000011111111")
	checkDisplay(t, b, "1111 1000  1111 1000  0000 0000  1111 1111")
}

func TestShiftLeft(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddBytes([]byte{0xFF, 0x00, 0xFF, 0x00})
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if err := b.ShiftLeft(10); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "00000011111111000000000000000000")
	checkDisplay(t, b, "0000 0011  1111 1100  0000 0000  0000 0000")

	// Make sure that bits before the starting point are not affected.
	b.Reset()
	b.AddBytes([]byte{0xFF, 0x00, 0xFF, 0x00})
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if n, err := b.Advance(5); n != 5 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Advance() test")
		t.Log("\tExpected: 5, <nil>")
		t.Log("\tReceived:", n, err)
	}
	if err := b.ShiftLeft(10); err != nil {
		t.Error(err)
	}
	if n, err := b.Rewind(5); n != 5 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Rewind() test")
		t.Log("\tExpected: 5, <nil>")
		t.Log("\tReceived:", n, err)
	}

	checkBits(t, b, 32)
	checkString(t, b, "11111011111111000000000000000000")
	checkDisplay(t, b, "1111 1011  1111 1100  0000 0000  0000 0000")

	// Test shifting if there's only 1 bit in the buffer.
	b.Reset()
	b.AddBit(true)
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	if err := b.ShiftLeft(10); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "0")
	checkDisplay(t, b, "0")
}

func TestShiftRight(t *testing.T) {
	b := New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.AddBytes([]byte{0xFF, 0x00, 0xFF, 0x00})
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if err := b.ShiftRight(10); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "00000000001111111100000000111111")
	checkDisplay(t, b, "0000 0000  0011 1111  1100 0000  0011 1111")

	// Make sure that bits before the starting point are not affected.
	b.Reset()
	b.AddBytes([]byte{0xFF, 0x00, 0xFF, 0x00})
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if n, err := b.Advance(5); n != 5 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Advance() test")
		t.Log("\tExpected: 5, <nil>")
		t.Log("\tReceived:", n, err)
	}
	if err := b.ShiftRight(10); err != nil {
		t.Error(err)
	}
	if n, err := b.Rewind(5); n != 5 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Rewind() test")
		t.Log("\tExpected: 5, <nil>")
		t.Log("\tReceived:", n, err)
	}

	checkBits(t, b, 32)
	checkString(t, b, "11111000000000011100000000111111")
	checkDisplay(t, b, "1111 1000  0000 0001  1100 0000  0011 1111")

	// Test shifting if there's only 1 bit in the buffer.
	b.Reset()
	b.AddBit(true)
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	if err := b.ShiftRight(10); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "0")
	checkDisplay(t, b, "0")
}

func TestWriteInt(t *testing.T) {
	// TODO: write tests
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
