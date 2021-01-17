package hstack_test

import (
	"reflect"
	"testing"

	"github.com/snhilde/dsa/data_structures/hstack"
)

func TestBadPtr(t *testing.T) {
	var s *hstack.Stack
	checkString(t, s, "<nil>")
	checkCount(t, s, -1)

	// Test Add().
	if err := s.Add("value"); err == nil {
		t.Error("unexpectedly passed Add() test with bad pointer")
	}

	// Test Pop().
	if v := s.Pop(); v != nil {
		t.Error("unexpectedly passed Pop() test with bad pointer")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Test Count().
	if n := s.Count(); n != -1 {
		t.Error("unexpectedly passed Count() test with bad pointer")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", n)
	}

	// Test Copy().
	if _, err := s.Copy(); err == nil {
		t.Error("unexpectedly passed Copy() test with bad pointer")
	}

	// Test Merge().
	if err := s.Merge(hstack.New()); err == nil {
		t.Error("unexpectedly passed Merge() test with bad pointer")
	}

	// Test Clear().
	if err := s.Clear(); err == nil {
		t.Error("unexpectedly passed Clear() test with bad pointer")
	}

	// Test Same().
	if ok := s.Same(hstack.New()); ok {
		t.Error("unexpectedly passed Same() test with bad pointer")
	}

	// Test String().
	if v := s.String(); v != "<nil>" {
		t.Error("unexpectedly passed String() test with bad pointer")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", v)
	}
}

func TestNew(t *testing.T) {
	s := hstack.New()
	if s == nil {
		t.Error("new stack unexpectedly nil")
	}
	checkString(t, s, "<empty>")
	checkCount(t, s, 0)
}

func TestAdd(t *testing.T) {
	s := hstack.New()

	// Testing adding an int.
	if err := s.Add(5); err != nil {
		t.Error(err)
	}
	checkString(t, s, "5")
	checkCount(t, s, 1)

	// Test adding a string.
	if err := s.Add("kangaroo"); err != nil {
		t.Error(err)
	}
	checkString(t, s, "kangaroo, 5")
	checkCount(t, s, 2)

	// Testing adding a float.
	if err := s.Add(3.1415); err != nil {
		t.Error(err)
	}
	checkString(t, s, "3.1415, kangaroo, 5")
	checkCount(t, s, 3)

	// Test out adding multiple items at once.
	if err := s.Add("a", "b", 3); err != nil {
		t.Error(err)
	}
	checkString(t, s, "3, b, a, 3.1415, kangaroo, 5")
	checkCount(t, s, 6)

	// Test adding a slice.
	if err := s.Add([]int{1, 2, 3}); err != nil {
		t.Error(err)
	}
	checkString(t, s, "[1 2 3], 3, b, a, 3.1415, kangaroo, 5")
	checkCount(t, s, 7)

	// Testing adding an empty stack.
	if err := s.Add(hstack.New()); err != nil {
		t.Error(err)
	}
	checkString(t, s, "<empty>, [1 2 3], 3, b, a, 3.1415, kangaroo, 5")
	checkCount(t, s, 8)

	// Test adding a non-empty stack.
	a := hstack.New()
	if err := a.Add("orange, apple, banana"); err != nil {
		t.Error(err)
	}
	if err := s.Add(a); err != nil {
		t.Error(err)
	}
	checkString(t, s, "orange, apple, banana, <empty>, [1 2 3], 3, b, a, 3.1415, kangaroo, 5")
	checkCount(t, s, 9)

	// Test adding stack to itself. This should fail.
	if err := s.Add(s); err == nil {
		t.Error("should not be able to add a stack to itself")
	}
	checkString(t, s, "orange, apple, banana, <empty>, [1 2 3], 3, b, a, 3.1415, kangaroo, 5")
	checkCount(t, s, 9)
}

func TestPop(t *testing.T) {
	s := hstack.New()

	// Add some items first.
	if err := s.Add("sizzle"); err != nil {
		t.Error(err)
	}
	if err := s.Add(1e5); err != nil {
		t.Error(err)
	}
	if err := s.Add(3.1415); err != nil {
		t.Error(err)
	}
	if err := s.Add(15); err != nil {
		t.Error(err)
	}
	checkString(t, s, "15, 3.1415, 100000, sizzle")
	checkCount(t, s, 4)

	// Test out popping the items.
	if val := s.Pop(); val != 15 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: 15")
		t.Log("\tReceived:", val)
	}
	checkString(t, s, "3.1415, 100000, sizzle")
	checkCount(t, s, 3)

	if val := s.Pop(); val != 3.1415 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: 3.1415")
		t.Log("\tReceived:", val)
	}
	checkString(t, s, "100000, sizzle")
	checkCount(t, s, 2)

	if val := s.Pop(); val != 1e5 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: 1e5")
		t.Log("\tReceived:", val)
	}
	checkString(t, s, "sizzle")
	checkCount(t, s, 1)

	if val := s.Pop(); val != "sizzle" {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: sizzle")
		t.Log("\tReceived:", val)
	}
	checkString(t, s, "<empty>")
	checkCount(t, s, 0)

	// Test popping from an empty stack.
	if val := s.Pop(); val != nil {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", val)
	}
	checkString(t, s, "<empty>")
	checkCount(t, s, 0)

	// Test popping a slice.
	slice := []int{1, 2, 3}
	if err := s.Add(slice); err != nil {
		t.Error(err)
	}
	checkString(t, s, "[1 2 3]")
	checkCount(t, s, 1)

	if p := s.Pop().([]int); !reflect.DeepEqual(slice, p) {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected:", slice)
		t.Log("\tReceived:", p)
	}
	checkString(t, s, "<empty>")
	checkCount(t, s, 0)

	// Test popping a stack.
	a := hstack.New()
	if err := a.Add("orange, apple, banana"); err != nil {
		t.Error(err)
	}
	checkString(t, a, "orange, apple, banana")
	checkCount(t, a, 1)

	if err := s.Add(a); err != nil {
		t.Error(err)
	}
	checkString(t, s, "orange, apple, banana")
	checkCount(t, s, 1)

	if val := s.Pop(); val.(*hstack.Stack) != a {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected:", a)
		t.Log("\tReceived:", val)
	}
	checkString(t, s, "<empty>")
	checkCount(t, s, 0)
}

func TestCopy(t *testing.T) {
	s := hstack.New()

	// Copy an empty stack.
	ns, err := s.Copy()
	if err != nil {
		t.Error(err)
	}
	checkString(t, s, "<empty>")
	checkCount(t, s, 0)
	checkString(t, ns, "<empty>")
	checkCount(t, ns, 0)

	// Copy a non-empty stack.
	if err := s.Add("sizzle", 1e5, 3.1415, 15); err != nil {
		t.Error(err)
	}
	checkString(t, s, "15, 3.1415, 100000, sizzle")
	checkCount(t, s, 4)

	ns, err = s.Copy()
	if err != nil {
		t.Error(err)
	}
	checkString(t, s, "15, 3.1415, 100000, sizzle")
	checkCount(t, s, 4)
	checkString(t, ns, "15, 3.1415, 100000, sizzle")
	checkCount(t, ns, 4)

}

func TestMerge(t *testing.T) {
	// Create two stacks and merge them.
	s := hstack.New()
	if err := s.Add("monkey"); err != nil {
		t.Error(err)
	}
	if err := s.Add("gazelle"); err != nil {
		t.Error(err)
	}
	if err := s.Add(131); err != nil {
		t.Error(err)
	}
	checkString(t, s, "131, gazelle, monkey")
	checkCount(t, s, 3)

	tmp := hstack.New()
	if err := tmp.Add(3.14); err != nil {
		t.Error(err)
	}
	if err := tmp.Add(16); err != nil {
		t.Error(err)
	}
	if err := tmp.Add([]uint{5, 6, 7}); err != nil {
		t.Error(err)
	}
	checkString(t, tmp, "[5 6 7], 16, 3.14")
	checkCount(t, tmp, 3)

	// Merge and check that tmp was added below s and that tmp was emptied out.
	if err := s.Merge(tmp); err != nil {
		t.Error(err)
	}
	checkString(t, s, "[5 6 7], 16, 3.14, 131, gazelle, monkey")
	checkCount(t, s, 6)
	checkString(t, tmp, "<empty>")
	checkCount(t, tmp, 0)

	// Test merging an invalid stack on top of a good one. Merge should succeed, but everything should remain untouched.
	var ns *hstack.Stack
	checkString(t, ns, "<nil>")
	checkCount(t, ns, -1)

	if err := s.Merge(ns); err != nil {
		t.Error(err)
	}
	checkString(t, s, "[5 6 7], 16, 3.14, 131, gazelle, monkey")
	checkCount(t, s, 6)
	checkString(t, ns, "<nil>")
	checkCount(t, ns, -1)

	// Test merging a good stack on top of a bad one. Merge should fail, and everything should remain untouched.
	if err := ns.Merge(s); err == nil {
		t.Error("unexpectedly passed merge on top of bad stack test")
	}
	checkString(t, ns, "<nil>")
	checkCount(t, ns, -1)
	checkString(t, s, "[5 6 7], 16, 3.14, 131, gazelle, monkey")
	checkCount(t, s, 6)

	// Test merging a stack with itself.
	if err := s.Merge(s); err != nil {
		t.Error(err)
	}
	checkString(t, s, "[5 6 7], 16, 3.14, 131, gazelle, monkey, [5 6 7], 16, 3.14, 131, gazelle, monkey")
	checkCount(t, s, 12)
}

func TestClear(t *testing.T) {
	s := hstack.New()

	// Add some items first.
	if err := s.Add("kangaroo", 5, 3.1415); err != nil {
		t.Error(err)
	}
	checkString(t, s, "3.1415, 5, kangaroo")
	checkCount(t, s, 3)

	// Test out clearing the stack.
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
	checkString(t, s, "<empty>")
	checkCount(t, s, 0)

	// Test out clearing an empty stack.
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
	checkString(t, s, "<empty>")
	checkCount(t, s, 0)
}

func TestSame(t *testing.T) {
	s := hstack.New()

	// Add some items first.
	if err := s.Add("kangaroo", 5, 3.1415); err != nil {
		t.Error(err)
	}
	checkString(t, s, "3.1415, 5, kangaroo")
	checkCount(t, s, 3)

	// Test out with the same stack.
	ns := s
	checkString(t, s, "3.1415, 5, kangaroo")
	checkCount(t, s, 3)
	checkString(t, ns, "3.1415, 5, kangaroo")
	checkCount(t, ns, 3)
	if ok := s.Same(ns); !ok {
		t.Error("identical stacks should pass Same")
	}

	// Test out with two different stacks that have the same contents.
	ns = hstack.New()
	if err := ns.Add("kangaroo", 5, 3.1415); err != nil {
		t.Error(err)
	}
	checkString(t, s, "3.1415, 5, kangaroo")
	checkCount(t, s, 3)
	checkString(t, ns, "3.1415, 5, kangaroo")
	checkCount(t, ns, 3)
	if ok := s.Same(ns); ok {
		t.Error("similar stacks should not pass Same")
	}

	// Test out with two different stacks with different contents.
	ns = hstack.New()
	if err := ns.Add(struct{}{}, 10, rune('b')); err != nil {
		t.Error(err)
	}
	checkString(t, s, "3.1415, 5, kangaroo")
	checkCount(t, s, 3)
	checkString(t, ns, "98, 10, {}")
	checkCount(t, ns, 3)
	if ok := s.Same(ns); ok {
		t.Error("different stacks should not pass Same")
	}
}

func checkString(t *testing.T, s *hstack.Stack, want string) {
	if s.String() != want {
		t.Error("stack contents are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkCount(t *testing.T, s *hstack.Stack, want int) {
	if s.Count() != want {
		t.Error("Incorrect length")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s.Count())
	}
}
