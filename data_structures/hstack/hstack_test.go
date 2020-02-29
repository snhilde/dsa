package hstack

import (
	"testing"
)


// TESTS
func TestNew(t *testing.T) {
}

func TestAdd(t *testing.T) {
}

func TestPop(t *testing.T) {
}

func TestCount(t *testing.T) {
}

func TestClear(t *testing.T) {
}

func TestMerge(t *testing.T) {
}


// HELPERS
func checkString(t *testing.T, stack *Hstack, want string) {
	if stack.String() != want {
		t.Error("stack contents are incorrect")
		t.Log("Expected:", want)
		t.Log("Received:", stack)
	}
}

func checkCount(t *testing.T, stack *Hstack, want int) {
	if stack.Count() != want {
		t.Error("Incorrect length")
		t.Log("Expected:", want)
		t.Log("Received:", stack.Count())
	}
}
