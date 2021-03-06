package hbit_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/snhilde/dsa/data_structures/hbit"
)

func TestNew(t *testing.T) {
	b := hbit.New()
	if b == nil {
		t.Error("Unexpectedly failed New() test")
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
	var b *hbit.Buffer

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
	if err := b.Reset(); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Reset()")
	}

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

	// Test Read().
	if n, err := b.Read([]byte{}); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Read()")
	}

	// Test ReadByte().
	if n, err := b.ReadByte(4); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ReadByte()")
	}

	// Test ReadInt().
	if n, err := b.ReadInt(6); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ReadInt()")
	}

	// Test ReadFrom().
	if n, err := b.ReadFrom(strings.NewReader("test")); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Read()")
	}

	// Test Write().
	if n, err := b.Write([]byte{}); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Write()")
	}

	// Test WriteBit().
	if err := b.WriteBit(true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for WriteBit()")
	}

	// Test WriteByte().
	if err := b.WriteByte(0x0C); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for WriteByte()")
	}

	// Test WriteBytes().
	if err := b.WriteBytes(0x0C, 0xFF); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for WriteBytes()")
	}

	// Test WriteString().
	if err := b.WriteString("test string"); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for WriteString()")
	}

	// Test SetBit().
	if err := b.SetBit(2, true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for SetBit()")
	}

	// Test SetBytes().
	if err := b.SetBytes(4, []byte{0xFF}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for SetBytes()")
	}

	// Test RemoveBit().
	if err := b.RemoveBit(5); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for RemoveBit()")
	}

	// Test RemoveBits().
	if err := b.RemoveBits(2, 5); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for RemoveBits()")
	}

	// Test Advance().
	if n, err := b.Advance(10); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Advance()")
	}

	// Test Rewind().
	if n, err := b.Rewind(10); n != 0 || err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Rewind()")
	}

	// Test Join().
	if err := b.Join(hbit.New()); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for Join()")
	}

	// Test ANDBit().
	if err := b.ANDBit(10, true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ANDBit()")
	}

	// Test ORBit().
	if err := b.ORBit(10, true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ORBit()")
	}

	// Test XORBit().
	if err := b.XORBit(10, true); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for XORBit()")
	}

	// Test ANDBytes().
	if err := b.ANDBytes([]byte{0xFF, 0xEE}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ANDBytes()")
	}

	// Test ORBytes().
	if err := b.ORBytes([]byte{0xFF, 0xEE}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ORBytes()")
	}

	// Test XORBytes().
	if err := b.XORBytes([]byte{0xFF, 0xEE}); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for XORBytes()")
	}

	// Test ANDBuffer().
	if err := b.ANDBuffer(hbit.New()); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ANDBuffer()")
	}

	// Test ORBuffer().
	if err := b.ORBuffer(hbit.New()); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ORBuffer()")
	}

	// Test XORBuffer().
	if err := b.XORBuffer(hbit.New()); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for XORBuffer()")
	}

	// Test ShiftLeft().
	if err := b.ShiftLeft(5); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ShiftLeft()")
	}

	// Test ShiftRight().
	if err := b.ShiftRight(6); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for ShiftRight()")
	}

	// Test NOTBit().
	if err := b.NOTBit(10); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for NOTBit()")
	}

	// Test NOTBits().
	if err := b.NOTBits(20); err == nil {
		t.Error("Unexpectedly passed bad Buffer test for NOTBits()")
	}
}

func TestInvalidArgs(t *testing.T) {
	// Make sure that every method is capable of handling bad arguments.
	b := hbit.New()
	b.WriteBytes(0xFF, 0xEE, 0xDD)

	// Test Bit() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if val := b.Bit(-1); val != false {
		t.Error("Incorrect result from negative index test for Bit()")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", val)
	}

	// Test Bit() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if val := b.Bit(100); val != false {
		t.Error("Incorrect result from out-of-range index test for Bit()")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", val)
	}

	// Test Copy() - negative value.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if nb := b.Copy(-1); nb.String() != "<empty>" {
		t.Error("Incorrect result from negative value test for Copy()")
		t.Log("\tExpected: <empty>")
		t.Log("\tReceived:", nb)
	}

	// Test Copy() - out-of-range value.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if nb := b.Copy(100); nb.String() != "111111110111011110111011" {
		t.Error("Incorrect result from out-of-range value test for Copy()")
		t.Log("\tExpected: 111111110111011110111011")
		t.Log("\tReceived:", nb)
	}

	// Test Read() - empty argument.
	if _, err := b.Read([]byte{}); err != nil {
		t.Error("Unexpectedly failed empty argument test for Read()")
		t.Error(err)
	}

	// Test ReadByte() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if n, err := b.ReadByte(-1); n != 0 || err == nil {
		t.Error("Unexpectedly passed negative index test for ReadByte()")
	}

	// Test ReadByte() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if n, err := b.ReadByte(100); n != 0 || err == nil {
		t.Error("Unexpectedly passed out-of-range index test for ReadByte()")
	}

	// Test ReadInt() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if n, err := b.ReadInt(-1); n != 0 || err == nil {
		t.Error("Unexpectedly passed negative index test for ReadInt()")
	}

	// Test ReadInt() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if n, err := b.ReadInt(100); n != 0 || err == nil {
		t.Error("Unexpectedly passed out-of-range index test for ReadInt()")
	}

	// Test ReadFrom() - empty argument.
	if _, err := b.ReadFrom(nil); err != io.EOF {
		t.Error("Unexpectedly passed empty argument test for ReadFrom()")
	}

	// Test Write() - empty argument.
	if _, err := b.Write([]byte{}); err != nil {
		t.Error("Unexpectedly failed empty argument test for Write()")
		t.Error(err)
	}

	// Test WriteBytes() - empty argument.
	if err := b.WriteBytes(); err != nil {
		t.Error("Unexpectedly failed empty argument test for WriteBytes()")
		t.Error(err)
	}

	// Test WriteString() - empty argument.
	if err := b.WriteString(""); err != nil {
		t.Error("Unexpectedly failed empty argument test for WriteString()")
		t.Error(err)
	}

	// Test SetBit() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.SetBit(-1, true); err == nil {
		t.Error("Unexpectedly passed negative index test for SetBit()")
	}

	// Test SetBit() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.SetBit(100, true); err == nil {
		t.Error("Unexpectedly passed out-of-range index test for SetBit()")
	}

	// Test SetBytes() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.SetBytes(-1, []byte{0xFF}); err == nil {
		t.Error("Unexpectedly passed negative index test for SetBytes()")
	}

	// Test SetBytes() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.SetBytes(100, []byte{0xFF}); err == nil {
		t.Error("Unexpectedly passed out-of-range index test for SetBytes()")
	}

	// Test SetBytes() - empty reference bytes.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.SetBytes(4, []byte{}); err != nil {
		t.Error("Unexpectedly failed empty bytes test for SetBytes()")
		t.Error(err)
	}

	// Test RemoveBit() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.RemoveBit(-1); err == nil {
		t.Error("Unexpectedly passed negative index test for RemoveBit()")
	}

	// Test RemoveBit() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.RemoveBit(100); err == nil {
		t.Error("Unexpectedly passed out-of-range index test for RemoveBit()")
	}

	// Test RemoveBits() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.RemoveBits(-1, 5); err == nil {
		t.Error("Unexpectedly passed negative index test for RemoveBits()")
	}

	// Test RemoveBits() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.RemoveBits(100, 5); err == nil {
		t.Error("Unexpectedly passed out-of-range index test for RemoveBits()")
	}

	// Test RemoveBits() - negative number.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.RemoveBits(1, -1); err != nil {
		t.Error("Unexpectedly failed negative number test for RemoveBits()")
		t.Error(err)
	}

	// Test RemoveBits() - out-of-range number.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.RemoveBits(0, 100); err != nil {
		t.Error("Unexpectedly failed out-of-range number test for RemoveBits()")
		t.Error(err)
	}

	// Test RemoveBits() - no numbers.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.RemoveBits(1, 0); err != nil {
		t.Error("Unexpectedly failed no number test for RemoveBits()")
		t.Error(err)
	}

	// Test Advance() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if n, err := b.Advance(-1); n != 0 || err == nil {
		t.Error("Unexpectedly passed negative index test for Advance()")
	}

	// Test Advance() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if n, err := b.Advance(100); n == 0 || err != nil {
		t.Error("Unexpectedly failed out-of-range index test for Advance()")
		t.Error(err)
	}

	// Test Rewind() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if n, err := b.Rewind(-1); n != 0 || err == nil {
		t.Error("Unexpectedly passed negative index test for Rewind()")
	}

	// Test Rewind() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	b.Advance(10)
	if n, err := b.Rewind(100); n != 10 || err != nil {
		t.Error("Unexpectedly failed out-of-range index test for Rewind()")
		t.Error(err)
	}

	// Test Join() - empty buffer.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.Join(nil); err != nil {
		t.Error("Unexpectedly failed empty buffer test for Join()")
		t.Error(err)
	}

	// Test ANDBit() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ANDBit(-1, true); err == nil {
		t.Error("Unexpectedly passed negative index test for ANDBit()")
	}

	// Test ANDBit() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ANDBit(100, true); err == nil {
		t.Error("Unexpectedly passed out-of-range index test for ANDBit()")
	}

	// Test ORBit() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ORBit(-1, true); err == nil {
		t.Error("Unexpectedly passed negative index test for ORBit()")
	}

	// Test ORBit() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ORBit(100, true); err == nil {
		t.Error("Unexpectedly passed out-of-range index test for ORBit()")
	}

	// Test XORBit() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.XORBit(-1, true); err == nil {
		t.Error("Unexpectedly passed negative index test for XORBit()")
	}

	// Test XORBit() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.XORBit(100, true); err == nil {
		t.Error("Unexpectedly passed out-of-range index test for XORBit()")
	}

	// Test ANDBytes() - empty reference bytes.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ANDBytes([]byte{}); err != nil {
		t.Error("Unexpectedly failed empty bytes test for ANDBytes()")
		t.Error(err)
	}

	// Test ORBytes() - empty reference bytes.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ORBytes([]byte{}); err != nil {
		t.Error("Unexpectedly failed empty bytes test for ORBytes()")
		t.Error(err)
	}

	// Test XORBytes() - empty reference bytes.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.XORBytes([]byte{}); err != nil {
		t.Error("Unexpectedly failed empty bytes test for XORBytes()")
		t.Error(err)
	}

	// Test ANDBuffer() - empty reference buffer.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ANDBuffer(nil); err == nil {
		t.Error("Unexpectedly passed empty bytes test for ANDBuffer()")
	}

	// Test ORBuffer() - empty reference buffer.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ORBuffer(nil); err == nil {
		t.Error("Unexpectedly passed empty bytes test for ORBuffer()")
	}

	// Test XORBuffer() - empty reference buffer.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.XORBuffer(nil); err == nil {
		t.Error("Unexpectedly passed empty bytes test for XORBuffer()")
	}

	// Test ShiftLeft() - negative number.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ShiftLeft(-1); err == nil {
		t.Error("Unexpectedly passed negative number test for ShiftLeft()")
		t.Error(err)
	}

	// Test ShiftLeft() - out-of-range number.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ShiftLeft(100); err != nil {
		t.Error("Unexpectedly failed out-of-range number test for ShiftLeft()")
		t.Error(err)
	}

	// Test ShiftRight() - negative number.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ShiftRight(-1); err == nil {
		t.Error("Unexpectedly passed negative number test for ShiftRight()")
		t.Error(err)
	}

	// Test ShiftRight() - out-of-range number.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.ShiftRight(100); err != nil {
		t.Error("Unexpectedly failed out-of-range number test for ShiftRight()")
		t.Error(err)
	}

	// Test NOTBit() - negative index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.NOTBit(-1); err == nil {
		t.Error("Unexpectedly passed negative index test for NOTBit()")
	}

	// Test NOTBit() - out-of-range index.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.NOTBit(100); err == nil {
		t.Error("Unexpectedly passed out-of-range index test for NOTBit()")
	}

	// Test NOTBits() - negative number.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.NOTBits(-1); err == nil {
		t.Error("Unexpectedly passed negative number test for NOTBits()")
	}

	// Test NOTBits() - out-of-range number.
	b.Reset()
	b.WriteBytes(0xFF, 0xEE, 0xDD)
	if err := b.NOTBits(100); err != nil {
		t.Error("Unexpectedly failed out-of-range number test for NOTBits()")
		t.Error(err)
	}
}

func TestBit(t *testing.T) {
	b := hbit.New()
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

	b.WriteByte(0x0F)
	if val := b.Bit(2); val != true {
		t.Error("Incorrect result from Bit() test")
		t.Log("\tExpected: true")
		t.Log("\tReceived:", val)
	}

	b.WriteByte(0x0F)
	if val := b.Bit(6); val != false {
		t.Error("Incorrect result from Bit() test")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", val)
	}
}

func TestBits(t *testing.T) {
	b := hbit.New()
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
	b.WriteBit(true)
	b.WriteBit(false)
	b.WriteBit(true)
	if n := b.Bits(); n != 3 {
		t.Error("Incorrect result from Bits() test")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", n)
	}

	// Test after adding a byte.
	b.WriteByte(0x00)
	if n := b.Bits(); n != 11 {
		t.Error("Incorrect result from Bits() test")
		t.Log("\tExpected: 11")
		t.Log("\tReceived:", n)
	}

	// Test after adding some bytes.
	b.WriteBytes(0x01, 0x02, 0x03, 0x04, 0x05, 0x06)
	if n := b.Bits(); n != 59 {
		t.Error("Incorrect result from Bits() test")
		t.Log("\tExpected: 59")
		t.Log("\tReceived:", n)
	}
}

func TestOffset(t *testing.T) {
	b := hbit.New()
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
	b.WriteBytes(0xFF, 0xFB)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test copying an entire buffer.
	b.WriteByte(0x0A)
	checkBits(t, b, 8)
	nb := b.Copy(b.Bits())
	if nb == nil {
		t.Error("Failed to create copy of buffer")
	}
	checkBits(t, nb, 8)

	// Make sure that modifying one buffer does not modify the copy.
	b.WriteBit(true)
	b.WriteBit(false)
	checkBits(t, b, 10)
	checkBits(t, nb, 8)

	b.Advance(5)
	checkBits(t, b, 5)
	checkBits(t, nb, 8)

	nb.WriteBytes(0xAA, 0xBB, 0xCC)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bytes, advancing the buffer, and setting it back.
	b.WriteBytes(0x50, 0x60, 0x70)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bits and then resetting the buffer.
	b.WriteBit(false)
	b.WriteBit(false)
	checkBits(t, b, 2)
	if err := b.Reset(); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)

	// Test adding some bits, advancing, and then resetting the buffer.
	b.WriteByte(0xFF)
	checkBits(t, b, 8)
	b.Advance(5)
	checkBits(t, b, 3)
	if err := b.Reset(); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)

	// Test adding some bits, advancing, reversing, and then resetting the buffer.
	b.WriteByte(0xEE)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBit(true)
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	b.WriteBit(false)
	checkBits(t, b, 2)
	checkString(t, b, "10")
	checkDisplay(t, b, "10")

	b.WriteByte(0xFF)
	checkBits(t, b, 10)
	checkString(t, b, "1011111111")
	checkDisplay(t, b, "1011 1111  11")

	// Make sure that a single byte does not have any spaces around it.
	b = hbit.New()
	b.WriteByte(0x00)
	checkBits(t, b, 8)
	checkString(t, b, "00000000")
	checkDisplay(t, b, "0000 0000")

	// Test out multiple bytes and bit order.
	b = hbit.New()
	b.WriteBytes(0x1A, 0x2B, 0x3C)
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

func TestRead(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test reading out one byte of all true bits. */
	b.WriteBytes(0xFF)
	checkBits(t, b, 8)
	checkString(t, b, "11111111")
	checkDisplay(t, b, "1111 1111")

	buf := make([]byte, 1)
	if n, err := b.Read(buf); n != 1 || err != nil {
		t.Error(err)
		t.Log(n)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	if buf[0] != 0xFF {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 0xFF")
		t.Log("\tReceived:", buf[0])
	}

	// Test reading out one byte of all false bits. */
	b.Reset()
	b.WriteBytes(0x00)
	checkBits(t, b, 8)
	checkString(t, b, "00000000")
	checkDisplay(t, b, "0000 0000")

	buf = make([]byte, 1)
	if n, err := b.Read(buf); n != 1 || err != nil {
		t.Error(err)
		t.Log(n)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	if buf[0] != 0x00 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 0x00")
		t.Log("\tReceived:", buf[0])
	}

	// Test reading out part of a byte.
	b.Reset()
	b.WriteBit(true)
	b.WriteBit(false)
	b.WriteBit(true)
	b.WriteBit(false)
	checkBits(t, b, 4)
	checkString(t, b, "1010")
	checkDisplay(t, b, "1010")

	buf = make([]byte, 1)
	if n, err := b.Read(buf); n != 1 || err != nil {
		t.Error(err)
		t.Log(n)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	if buf[0] != 0x05 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 0x05")
		t.Log("\tReceived:", buf[0])
	}

	// Test reading out a full byte with part of a byte remaining.
	b.Reset()
	b.WriteByte(0xC4)
	b.WriteBit(true)
	b.WriteBit(false)
	b.WriteBit(true)
	b.WriteBit(false)
	checkBits(t, b, 12)
	checkString(t, b, "001000111010")
	checkDisplay(t, b, "0010 0011  1010")

	buf = make([]byte, 1)
	if n, err := b.Read(buf); n != 1 || err != nil {
		t.Error(err)
		t.Log(n)
	}
	checkBits(t, b, 4)
	checkString(t, b, "1010")
	checkDisplay(t, b, "1010")

	if buf[0] != 0xC4 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 0xC4")
		t.Log("\tReceived:", buf[0])
	}

	// Test reading with a 2-byte arg on a 1-byte buffer.
	b.Reset()
	b.WriteBytes(0x33)
	checkBits(t, b, 8)
	checkString(t, b, "11001100")
	checkDisplay(t, b, "1100 1100")

	buf = make([]byte, 2)
	if n, err := b.Read(buf); n != 1 || err != nil {
		t.Error(err)
		t.Log(n)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	if buf[0] != 0x33 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 0x33")
		t.Log("\tReceived:", buf[0])
	}

	// Now make sure a subsequent read returns io.EOF.
	if n, err := b.Read(buf); n != 0 || err != io.EOF {
		t.Error(err)
		t.Log(n)
	}
}

func TestReadByte(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test out one bit.
	b.WriteBit(true)
	if n, _ := b.ReadByte(0); n != 0x01 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 1")
		t.Log("\tReceived:", n)
	}

	// Test out a simple byte.
	b.Reset()
	b.WriteBytes(0x05)
	if n, _ := b.ReadByte(0); n != 0x05 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 5")
		t.Log("\tReceived:", n)
	}

	// Test out 4 bytes.
	b.Reset()
	b.WriteBytes(0x0F, 0xF0, 0xFF, 0x08)
	if n, _ := b.ReadByte(0); n != 15 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 15")
		t.Log("\tReceived:", n)
	}

	if n, _ := b.ReadByte(16); n != 255 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 255")
		t.Log("\tReceived:", n)
	}

	// Test out advancing, reading out, and rewinding.
	b.Reset()
	b.WriteBytes(0x0F, 0xF0, 0xFF, 0x08)
	b.Advance(10)
	if n, _ := b.ReadByte(0); n != 252 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 252")
		t.Log("\tReceived:", n)
	}

	b.Rewind(10)
	if n, _ := b.ReadByte(1); n != 7 {
		t.Error("Incorrect result from ReadByte() test")
		t.Log("\tExpected: 7")
		t.Log("\tReceived:", n)
	}
}

func TestReadInt(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test out one bit.
	b.WriteBit(true)
	if n, _ := b.ReadInt(0); n != 1 {
		t.Error("Incorrect result from ReadInt() test")
		t.Log("\tExpected: 1")
		t.Log("\tReceived:", n)
	}

	// Test out a simple byte.
	b.Reset()
	b.WriteBytes(0x05)
	if n, _ := b.ReadInt(0); n != 5 {
		t.Error("Incorrect result from ReadInt() test")
		t.Log("\tExpected: 5")
		t.Log("\tReceived:", n)
	}

	// Test out 4 bytes.
	b.Reset()
	b.WriteBytes(0x0F, 0xFF, 0xFF, 0xFF)
	if n, _ := b.ReadInt(0); n != -241 {
		t.Error("Incorrect result from ReadInt() test")
		t.Log("\tExpected: -241")
		t.Log("\tReceived:", n)
	}

	// Test out a negative number.
	b.Reset()
	b.WriteBytes(0xF0, 0xFF, 0xFF, 0xFF)
	if n, _ := b.ReadInt(0); n != -16 {
		t.Error("Incorrect result from ReadInt() test")
		t.Log("\tExpected: -16")
		t.Log("\tReceived:", n)
	}
}

func TestReadFrom(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test reading from a strings reader.
	r := strings.NewReader("test string")
	if n, err := b.ReadFrom(r); n != 11 || err != nil {
		t.Error("Unexpectedly failed ReadFrom() test")
		t.Error(err)
	}
	checkBits(t, b, 88)
	checkString(t, b, "0010111010100110110011100010111000000100110011100010111001001110100101100111011011100110")
	checkDisplay(t, b, "0010 1110  1010 0110  1100 1110  0010 1110  0000 0100  1100 1110  0010 1110  0100 1110  1001 0110  0111 0110  1110 0110")

	// Test reading from a strings reader.
	b.Reset()
	r = strings.NewReader("A")
	if n, err := b.ReadFrom(r); n != 1 || err != nil {
		t.Error("Unexpectedly failed ReadFrom() test")
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "10000010")
	checkDisplay(t, b, "1000 0010")

	// Test reading from a bytes reader.
	b.Reset()
	br := bytes.NewReader([]byte{0x11, 0x22, 0x33})
	if n, err := b.ReadFrom(br); n != 3 || err != nil {
		t.Error("Unexpectedly failed ReadFrom() test")
		t.Error(err)
	}
	checkBits(t, b, 24)
	checkString(t, b, "100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100")
}

func TestWrite(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test writing hex bytes.
	p := []byte{0x11, 0x22, 0x33}
	if n, err := b.Write(p); n != 3 || err != nil {
		t.Error("Unexpectedly failed Write() test")
		t.Error(err)
	}
	checkBits(t, b, 24)
	checkString(t, b, "100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100")

	// Test writing characters.
	b.Reset()
	s := "test string"
	if n, err := b.Write([]byte(s)); n != 11 || err != nil {
		t.Error("Unexpectedly failed Write() test")
		t.Error(err)
	}
	checkBits(t, b, 88)
	checkString(t, b, "0010111010100110110011100010111000000100110011100010111001001110100101100111011011100110")
	checkDisplay(t, b, "0010 1110  1010 0110  1100 1110  0010 1110  0000 0100  1100 1110  0010 1110  0100 1110  1001 0110  0111 0110  1110 0110")
}

func TestWriteBit(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bits.
	if err := b.WriteBit(true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	if err := b.WriteBit(true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 2)
	checkString(t, b, "11")
	checkDisplay(t, b, "11")

	if err := b.WriteBit(false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 3)
	checkString(t, b, "110")
	checkDisplay(t, b, "110")

	// Test advancing and adding a bit.
	b.Advance(2)
	if err := b.WriteBit(true); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 2)
	checkString(t, b, "01")
	checkDisplay(t, b, "01")

	// Test reversing and adding a bit.
	b.Rewind(1)
	if err := b.WriteBit(false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 4)
	checkString(t, b, "1010")
	checkDisplay(t, b, "1010")

	// Test resetting and adding a bit.
	b.Reset()
	if err := b.WriteBit(false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "0")
	checkDisplay(t, b, "0")
}

func TestWriteByte(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bytes.
	if err := b.WriteByte(0xF0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "00001111")
	checkDisplay(t, b, "0000 1111")

	if err := b.WriteByte(0x88); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 16)
	checkString(t, b, "0000111100010001")
	checkDisplay(t, b, "0000 1111  0001 0001")

	// Test advancing and adding a byte.
	b.Advance(10)
	if err := b.WriteByte(0x14); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 14)
	checkString(t, b, "01000100101000")
	checkDisplay(t, b, "0100 0100  1010 00")

	// Test reversing and adding a byte.
	b.Rewind(3)
	if err := b.WriteByte(0xA0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 25)
	checkString(t, b, "1000100010010100000000101")
	checkDisplay(t, b, "1000 1000  1001 0100  0000 0010  1")

	// Test resetting and adding a byte.
	b.Reset()
	if err := b.WriteByte(0x44); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "00100010")
	checkDisplay(t, b, "0010 0010")
}

func TestWriteBytes(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test adding some bytes.
	if err := b.WriteBytes(0x54, 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 16)
	checkString(t, b, "0010101010000000")
	checkDisplay(t, b, "0010 1010  1000 0000")

	if err := b.WriteBytes(0xAA, 0xBB, 0xCC); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 40)
	checkString(t, b, "0010101010000000010101011101110100110011")
	checkDisplay(t, b, "0010 1010  1000 0000  0101 0101  1101 1101  0011 0011")

	// Test advancing and adding some bytes.
	b.Advance(30)
	if err := b.WriteBytes(0x01, 0x02); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 26)
	checkString(t, b, "01001100111000000001000000")
	checkDisplay(t, b, "0100 1100  1110 0000  0001 0000  00")

	// Test reversing and adding a byte.
	b.Rewind(5)
	if err := b.WriteBytes(0x08); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 39)
	checkString(t, b, "101110100110011100000000100000000010000")
	checkDisplay(t, b, "1011 1010  0110 0111  0000 0000  1000 0000  0010 000")

	// Test resetting and adding some bytes.
	b.Reset()
	if err := b.WriteBytes(0x98, 0x76); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 16)
	checkString(t, b, "0001100101101110")
	checkDisplay(t, b, "0001 1001  0110 1110")
}

func TestWriteString(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Add a string.
	b.Reset()
	s := "test string"
	if n, err := b.Write([]byte(s)); n != 11 || err != nil {
		t.Error("Unexpectedly failed Write() test")
		t.Error(err)
	}
	checkBits(t, b, 88)
	checkString(t, b, "0010111010100110110011100010111000000100110011100010111001001110100101100111011011100110")
	checkDisplay(t, b, "0010 1110  1010 0110  1100 1110  0010 1110  0000 0100  1100 1110  0010 1110  0100 1110  1001 0110  0111 0110  1110 0110")

	// Add another string after that.
	s = "ok"
	if n, err := b.Write([]byte(s)); n != 2 || err != nil {
		t.Error("Unexpectedly failed Write() test")
		t.Error(err)
	}
	checkBits(t, b, 104)
	checkString(t, b, "00101110101001101100111000101110000001001100111000101110010011101001011001110110111001101111011011010110")
	checkDisplay(t, b, "0010 1110  1010 0110  1100 1110  0010 1110  0000 0100  1100 1110  0010 1110  0100 1110  1001 0110  0111 0110  1110 0110  1111 0110  1101 0110")
}

func TestSetBit(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test setting the first bit.
	b.WriteByte(0xFF)
	if err := b.SetBit(0, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "01111111")
	checkDisplay(t, b, "0111 1111")

	// Test setting the last bit.
	b.Reset()
	b.WriteByte(0xFF)
	if err := b.SetBit(7, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11111110")
	checkDisplay(t, b, "1111 1110")

	// Test setting a middle bit.
	b.Reset()
	b.WriteByte(0xFF)
	if err := b.SetBit(3, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "11101111")
	checkDisplay(t, b, "1110 1111")

	// Test advancing and setting a bit.
	b.Reset()
	b.WriteByte(0xFF)
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "11111")
	checkDisplay(t, b, "1111 1")

	if err := b.SetBit(0, false); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 5)
	checkString(t, b, "01111")
	checkDisplay(t, b, "0111 1")

	b.Rewind(3)
	checkBits(t, b, 8)
	checkString(t, b, "11101111")
	checkDisplay(t, b, "1110 1111")
}

func TestSetBytes(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteByte(0x50)
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
	b.WriteByte(0x50)
	if err := b.SetBytes(0, []byte{0x4C, 0xFF}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 8)
	checkString(t, b, "00110010")
	checkDisplay(t, b, "0011 0010")

	// Test advancing and setting a byte.
	b.Reset()
	b.WriteByte(0x51)
	checkString(t, b, "10001010")
	checkDisplay(t, b, "1000 1010")
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	if err := b.SetBytes(0, []byte{0x00}); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 5)
	checkString(t, b, "00000")
	checkDisplay(t, b, "0000 0")

	b.Rewind(5)
	checkBits(t, b, 8)
	checkString(t, b, "10000000")
	checkDisplay(t, b, "1000 0000")
}

func TestRemoveBit(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteByte(0x50)
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
	b.WriteByte(0x51)
	if err := b.RemoveBit(0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 7)
	checkString(t, b, "0001010")
	checkDisplay(t, b, "0001 010")

	// Test removing the bit at the end of the buffer.
	b.Reset()
	b.WriteByte(0x51)
	if err := b.RemoveBit(b.Bits() - 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 7)
	checkString(t, b, "1000101")
	checkDisplay(t, b, "1000 101")

	// Test removing a bit when there's only 1 bit in the buffer.
	b.Reset()
	b.WriteBit(true)
	if err := b.RemoveBit(0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test removing the first bit when there are only 2 bits in the buffer.
	b.Reset()
	b.WriteBit(true)
	b.WriteBit(false)
	if err := b.RemoveBit(0); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "0")
	checkDisplay(t, b, "0")

	// Test removing the second bit when there are only 2 bits in the buffer.
	b.Reset()
	b.WriteBit(true)
	b.WriteBit(false)
	if err := b.RemoveBit(1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	// Test advancing and removing the first bit.
	b.Reset()
	b.WriteByte(0x51)
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
	b.WriteByte(0x51)
	b.Advance(3)
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	if err := b.RemoveBit(b.Bits() - 1); err != nil {
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteByte(0x50)
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
	b.WriteByte(0x51)
	if err := b.RemoveBits(0, 3); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 5)
	checkString(t, b, "01010")
	checkDisplay(t, b, "0101 0")

	// Test removing some bits at the end of the buffer.
	b.Reset()
	b.WriteByte(0x51)
	if err := b.RemoveBits(b.Bits()-3, 3); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 5)
	checkString(t, b, "10001")
	checkDisplay(t, b, "1000 1")

	// Test removing a bit when there's only 1 bit in the buffer.
	b.Reset()
	b.WriteBit(true)
	if err := b.RemoveBits(0, 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test removing the first bit when there are only 2 bits in the buffer.
	b.Reset()
	b.WriteBit(true)
	b.WriteBit(false)
	if err := b.RemoveBits(0, 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "0")
	checkDisplay(t, b, "0")

	// Test removing the second bit when there are only 2 bits in the buffer.
	b.Reset()
	b.WriteBit(true)
	b.WriteBit(false)
	if err := b.RemoveBits(1, 1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 1)
	checkString(t, b, "1")
	checkDisplay(t, b, "1")

	// Test advancing and removing some bits.
	b.Reset()
	b.WriteByte(0x51)
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
	b.WriteByte(0x51)
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
	b.WriteByte(0x51)
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
	checkBits(t, b, 3)
	checkString(t, b, "100")
	checkDisplay(t, b, "100")
	if n := b.Offset(); n != 0 {
		t.Error("Incorrect result from Offset() test")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", n)
	}

	// Test removing more bits than currently exist in the buffer.
	b.Reset()
	b.WriteByte(0x51)
	if err := b.RemoveBits(2, 10); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 2)
	checkString(t, b, "10")
	checkDisplay(t, b, "10")
}

func TestAdvance(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test a normal advance.
	b.WriteByte(0xFF)
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
	if n, err := b.Advance(10); n != 3 || err != nil {
		t.Error(err)
		t.Error("Incorrect result from Advance() test")
		t.Log("\tExpected: 3, <nil>")
		t.Log("\tReceived:", n, err)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Rewind the buffer to make sure that we didn't overrun the end.
	if n, err := b.Rewind(2); n != 2 || err != nil {
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test advancing and reversing back to the starting point.
	b.WriteByte(0xFF)
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

func TestJoin(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test merging an empty buffer with a full buffer.
	b.WriteBytes(0x11, 0x22, 0x33)
	checkBits(t, b, 24)
	checkString(t, b, "100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100")

	nb := hbit.New()
	checkBits(t, nb, 0)
	checkString(t, nb, "<empty>")
	checkDisplay(t, nb, "<empty>")

	if err := b.Join(nb); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 24)
	checkString(t, b, "100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100")

	checkBits(t, nb, 0)
	checkString(t, nb, "<empty>")
	checkDisplay(t, nb, "<empty>")

	// Test merging a full buffer with an empty one.
	b.Reset()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	nb.Reset()
	nb.WriteBytes(0x11, 0x22, 0x33)
	checkBits(t, nb, 24)
	checkString(t, nb, "100010000100010011001100")
	checkDisplay(t, nb, "1000 1000  0100 0100  1100 1100")

	if err := b.Join(nb); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 24)
	checkString(t, b, "100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100")

	checkBits(t, nb, 0)
	checkString(t, nb, "<empty>")
	checkDisplay(t, nb, "<empty>")

	// Test merging two empty buffers.
	b.Reset()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	nb.Reset()
	checkBits(t, nb, 0)
	checkString(t, nb, "<empty>")
	checkDisplay(t, nb, "<empty>")

	if err := b.Join(nb); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	checkBits(t, nb, 0)
	checkString(t, nb, "<empty>")
	checkDisplay(t, nb, "<empty>")

	// Test merging two full buffers.
	b.Reset()
	b.WriteBytes(0x11, 0x22, 0x33)
	checkBits(t, b, 24)
	checkString(t, b, "100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100")

	nb.Reset()
	nb.WriteBytes(0x44, 0x55, 0x66)
	checkBits(t, nb, 24)
	checkString(t, nb, "001000101010101001100110")
	checkDisplay(t, nb, "0010 0010  1010 1010  0110 0110")

	if err := b.Join(nb); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 48)
	checkString(t, b, "100010000100010011001100001000101010101001100110")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100  0010 0010  1010 1010  0110 0110")

	checkBits(t, nb, 0)
	checkString(t, nb, "<empty>")
	checkDisplay(t, nb, "<empty>")

	// Test merging the same buffer onto itself (shouldn't work).
	b.Reset()
	b.WriteBytes(0x11, 0x22, 0x33)
	checkBits(t, b, 24)
	checkString(t, b, "100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100")

	if err := b.Join(b); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	// Test copying a buffer and merging it back onto the original buffer.
	b.Reset()
	b.WriteBytes(0x11, 0x22, 0x33)
	checkBits(t, b, 24)
	checkString(t, b, "100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100")

	nb = b.Copy(b.Bits())
	if err := b.Join(nb); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 48)
	checkString(t, b, "100010000100010011001100100010000100010011001100")
	checkDisplay(t, b, "1000 1000  0100 0100  1100 1100  1000 1000  0100 0100  1100 1100")
}

func TestANDBit(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteByte(0x55)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteByte(0x55)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteByte(0x55)
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

func TestANDBytes(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
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

func TestANDBuffer(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	ref := hbit.New()
	checkBits(t, ref, 0)
	checkString(t, ref, "<empty>")
	checkDisplay(t, ref, "<empty>")

	ref.WriteBytes(0xF0, 0xF0, 0x0F, 0x0F)
	checkBits(t, ref, 32)
	checkString(t, ref, "00001111000011111111000011110000")
	checkDisplay(t, ref, "0000 1111  0000 1111  1111 0000  1111 0000")

	if err := b.ANDBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "00001111000000001111000000000000")
	checkDisplay(t, b, "0000 1111  0000 0000  1111 0000  0000 0000")
	checkBits(t, ref, 32)
	checkString(t, ref, "00001111000011111111000011110000")
	checkDisplay(t, ref, "0000 1111  0000 1111  1111 0000  1111 0000")

	// Test advancing, ANDing, and reversing.
	b.Advance(5)
	if err := b.ANDBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 27)
	checkString(t, b, "000000000000111000000000000")
	checkDisplay(t, b, "0000 0000  0000 1110  0000 0000  000")

	b.Rewind(5)
	checkBits(t, b, 32)
	checkString(t, b, "00001000000000000111000000000000")
	checkDisplay(t, b, "0000 1000  0000 0000  0111 0000  0000 0000")

	// Test buffers of different sizes.
	b.Reset()
	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
	ref.Reset()
	ref.WriteBytes(0xF0, 0xF0)

	if err := b.ANDBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "00001111000000001111111100000000")
	checkDisplay(t, b, "0000 1111  0000 0000  1111 1111  0000 0000")
}

func TestORBuffer(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	ref := hbit.New()
	checkBits(t, ref, 0)
	checkString(t, ref, "<empty>")
	checkDisplay(t, ref, "<empty>")

	ref.WriteBytes(0xF0, 0xF0, 0x0F, 0x0F)
	checkBits(t, ref, 32)
	checkString(t, ref, "00001111000011111111000011110000")
	checkDisplay(t, ref, "0000 1111  0000 1111  1111 0000  1111 0000")

	if err := b.ORBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "11111111000011111111111111110000")
	checkDisplay(t, b, "1111 1111  0000 1111  1111 1111  1111 0000")
	checkBits(t, ref, 32)
	checkString(t, ref, "00001111000011111111000011110000")
	checkDisplay(t, ref, "0000 1111  0000 1111  1111 0000  1111 0000")

	// Test advancing, ANDing, and reversing.
	b.Advance(5)
	if err := b.ORBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 27)
	checkString(t, b, "111011111111111111111110111")
	checkDisplay(t, b, "1110 1111  1111 1111  1111 1110  111")

	b.Rewind(5)
	checkBits(t, b, 32)
	checkString(t, b, "11111111011111111111111111110111")
	checkDisplay(t, b, "1111 1111  0111 1111  1111 1111  1111 0111")

	// Test buffers of different sizes.
	b.Reset()
	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
	ref.Reset()
	ref.WriteBytes(0xF0, 0xF0)

	if err := b.ORBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "11111111000011111111111100000000")
	checkDisplay(t, b, "1111 1111  0000 1111  1111 1111  0000 0000")
}

func TestXORBuffer(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	ref := hbit.New()
	checkBits(t, ref, 0)
	checkString(t, ref, "<empty>")
	checkDisplay(t, ref, "<empty>")

	ref.WriteBytes(0xF0, 0xF0, 0x0F, 0x0F)
	checkBits(t, ref, 32)
	checkString(t, ref, "00001111000011111111000011110000")
	checkDisplay(t, ref, "0000 1111  0000 1111  1111 0000  1111 0000")

	if err := b.XORBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "11110000000011110000111111110000")
	checkDisplay(t, b, "1111 0000  0000 1111  0000 1111  1111 0000")
	checkBits(t, ref, 32)
	checkString(t, ref, "00001111000011111111000011110000")
	checkDisplay(t, ref, "0000 1111  0000 1111  1111 0000  1111 0000")

	// Test advancing, ANDing, and reversing.
	b.Advance(5)
	if err := b.XORBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 27)
	checkString(t, b, "000011101110111000001110111")
	checkDisplay(t, b, "0000 1110  1110 1110  0000 1110  111")

	b.Rewind(5)
	checkBits(t, b, 32)
	checkString(t, b, "11110000011101110111000001110111")
	checkDisplay(t, b, "1111 0000  0111 0111  0111 0000  0111 0111")

	// Test buffers of different sizes.
	b.Reset()
	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
	ref.Reset()
	ref.WriteBytes(0xF0, 0xF0)

	if err := b.XORBuffer(ref); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "11110000000011111111111100000000")
	checkDisplay(t, b, "1111 0000  0000 1111  1111 1111  0000 0000")
}

func TestShiftLeft(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
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
	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
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
	b.WriteBit(true)
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
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
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
	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
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
	b.WriteBit(true)
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

func TestNOTBit(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteByte(0x55)
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

func TestNOTBits(t *testing.T) {
	b := hbit.New()
	checkBits(t, b, 0)
	checkString(t, b, "<empty>")
	checkDisplay(t, b, "<empty>")

	b.WriteBytes(0xFF, 0x00, 0xFF, 0x00)
	checkBits(t, b, 32)
	checkString(t, b, "11111111000000001111111100000000")
	checkDisplay(t, b, "1111 1111  0000 0000  1111 1111  0000 0000")

	if err := b.NOTBits(2); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 32)
	checkString(t, b, "00000000111111111111111100000000")
	checkDisplay(t, b, "0000 0000  1111 1111  1111 1111  0000 0000")

	// Test advancing, NOTing, and reversing.
	b.Advance(5)
	if err := b.NOTBits(1); err != nil {
		t.Error(err)
	}
	checkBits(t, b, 27)
	checkString(t, b, "111000001111111111100000000")
	checkDisplay(t, b, "1110 0000  1111 1111  1110 0000  000")

	b.Rewind(5)
	checkBits(t, b, 32)
	checkString(t, b, "00000111000001111111111100000000")
	checkDisplay(t, b, "0000 0111  0000 0111  1111 1111  0000 0000")

	if err := b.NOTBits(10); err != nil {
		t.Error(err)
	}
	// Test overrunning the buffer.
	checkBits(t, b, 32)
	checkString(t, b, "11111000111110000000000011111111")
	checkDisplay(t, b, "1111 1000  1111 1000  0000 0000  1111 1111")
}

func checkBits(t *testing.T, b *hbit.Buffer, want int) {
	if n := b.Bits(); n != want {
		t.Error("Incorrect number of bits")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", n)
	}
}

func checkString(t *testing.T, b *hbit.Buffer, want string) {
	if s := b.String(); s != want {
		t.Error("Incorrect string")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkDisplay(t *testing.T, b *hbit.Buffer, want string) {
	if s := b.Display(); s != want {
		t.Error("Incorrect display")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}
