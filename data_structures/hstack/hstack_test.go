package hstack

import (
	"testing"
)


// TESTS
func TestNew(t *testing.T) {
	// Test out making a proper stack.
	stack := New()
	if stack == nil {
		t.Error("new stack unexpectedly nil")
	}
	checkString(t, stack, "<empty>")
	checkCount(t, stack, 0)

	// Test out making a stack pointer to nothing.
	var stack_ptr *Stack
	checkString(t, stack_ptr, "<nil>")
	checkCount(t, stack_ptr, -1)

	// Test out the backdoor method.
	var backdoor Stack
	checkString(t, &backdoor, "<empty>")
	checkCount(t, &backdoor, 0)
}

func TestAdd(t *testing.T) {
	stack := New()
	checkString(t, stack, "<empty>")
	checkCount(t, stack, 0)

	// Test out adding some items to the stack.
	err := stack.Add(5)
	if err != nil {
		t.Error(err)
	}
	checkString(t, stack, "5")
	checkCount(t, stack, 1)

	err = stack.Add("kangaroo")
	if err != nil {
		t.Error(err)
	}
	checkString(t, stack, "kangaroo, 5")
	checkCount(t, stack, 2)

	err = stack.Add(3.1415)
	if err != nil {
		t.Error(err)
	}
	checkString(t, stack, "3.1415, kangaroo, 5")
	checkCount(t, stack, 3)

	// Test out adding to a bad stack.
	var bad_stack *Stack
	checkString(t, bad_stack, "<nil>")
	checkCount(t, bad_stack, -1)

	err = bad_stack.Add(10)
	if err == nil {
		t.Error("unexpectedly passed add to bad stack test")
	}
	checkString(t, bad_stack, "<nil>")
	checkCount(t, bad_stack, -1)

	// Test out adding to a backdoor stack.
	var backdoor Stack
	checkString(t, &backdoor, "<empty>")
	checkCount(t, &backdoor, 0)

	err = backdoor.Add(20)
	if err != nil {
		t.Error(err)
	}
	checkString(t, &backdoor, "20")
	checkCount(t, &backdoor, 1)
}

func TestPop(t *testing.T) {
	stack := New()
	checkString(t, stack, "<empty>")
	checkCount(t, stack, 0)

	// Add some items first.
	stack.Add("sizzle")
	stack.Add(1e5)
	stack.Add(3.1415)
	stack.Add(15)
	checkString(t, stack, "15, 3.1415, 100000, sizzle")
	checkCount(t, stack, 4)

	// Test out popping the items.
	val := stack.Pop()
	if val != 15 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected 15")
		t.Log("\tReceived", val)
	}
	checkString(t, stack, "3.1415, 100000, sizzle")
	checkCount(t, stack, 3)

	val = stack.Pop()
	if val != 3.1415 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected 3.1415")
		t.Log("\tReceived", val)
	}
	checkString(t, stack, "100000, sizzle")
	checkCount(t, stack, 2)

	val = stack.Pop()
	if val != 1e5 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected 1e5")
		t.Log("\tReceived", val)
	}
	checkString(t, stack, "sizzle")
	checkCount(t, stack, 1)

	val = stack.Pop()
	if val != "sizzle" {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected sizzle")
		t.Log("\tReceived", val)
	}
	checkString(t, stack, "<empty>")
	checkCount(t, stack, 0)

	val = stack.Pop()
	if val != nil {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected nil")
		t.Log("\tReceived", val)
	}
	checkString(t, stack, "<empty>")
	checkCount(t, stack, 0)

	// Test out popping from a bad stack.
	var bad_stack *Stack
	checkString(t, bad_stack, "<nil>")
	checkCount(t, bad_stack, -1)

	val = bad_stack.Pop()
	if val != nil {
		t.Error("unexpectedly passed pop from bad stack test")
	}
	checkString(t, bad_stack, "<nil>")
	checkCount(t, bad_stack, -1)

	// Test out popping from a backdoor stack.
	var backdoor Stack
	checkString(t, &backdoor, "<empty>")
	checkCount(t, &backdoor, 0)

	backdoor.Add(20)
	backdoor.Add("zebra")
	checkString(t, &backdoor, "zebra, 20")
	checkCount(t, &backdoor, 2)

	val = backdoor.Pop()
	if val != "zebra" {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected zebra")
		t.Log("\tReceived", val)
	}
	checkString(t, &backdoor, "20")
	checkCount(t, &backdoor, 1)

	val = backdoor.Pop()
	if val != 20 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected 20")
		t.Log("\tReceived", val)
	}
	checkString(t, &backdoor, "<empty>")
	checkCount(t, &backdoor, 0)

	val = backdoor.Pop()
	if val != nil {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected nil")
		t.Log("\tReceived", val)
	}
	checkString(t, &backdoor, "<empty>")
	checkCount(t, &backdoor, 0)
}

func TestClear(t *testing.T) {
	stack := New()
	checkString(t, stack, "<empty>")
	checkCount(t, stack, 0)

	// Add some items first.
	stack.Add("kangaroo")
	stack.Add(5)
	stack.Add(3.1415)
	checkString(t, stack, "3.1415, 5, kangaroo")
	checkCount(t, stack, 3)

	// Test out clearing the stack.
	err := stack.Clear()
	if err != nil {
		t.Error(err)
	}
	checkString(t, stack, "<empty>")
	checkCount(t, stack, 0)

	// Test out clearing a backdoor stack.
	var backdoor Stack
	checkString(t, &backdoor, "<empty>")
	checkCount(t, &backdoor, 0)

	backdoor.Add(20)
	backdoor.Add(1e2)
	backdoor.Add("lion")
	checkString(t, &backdoor, "lion, 100, 20")
	checkCount(t, &backdoor, 3)

	err = backdoor.Clear()
	if err != nil {
		t.Error(err)
	}
	checkString(t, &backdoor, "<empty>")
	checkCount(t, &backdoor, 0)
}

func TestMerge(t *testing.T) {
	// Create two stacks and merge them.
	base := New()
	base.Add("monkey")
	base.Add("gazelle")
	base.Add(131)
	checkString(t, base, "131, gazelle, monkey")
	checkCount(t, base, 3)

	tmp := New()
	tmp.Add(3.14)
	tmp.Add(16)
	tmp.Add("elephant")
	checkString(t, tmp, "elephant, 16, 3.14")
	checkCount(t, tmp, 3)

	// Merge and check that tmp was added on top of base and that tmp was emptied out.
	err := base.Merge(tmp)
	if err != nil {
		t.Error(err)
	}
	checkString(t, base, "elephant, 16, 3.14, 131, gazelle, monkey")
	checkCount(t, base, 6)
	checkString(t, tmp, "<empty>")
	checkCount(t, tmp, 0)

	// Test merging a bad stack on top of a good one.
	var bad_stack *Stack
	checkString(t, bad_stack, "<nil>")
	checkCount(t, bad_stack, -1)

	// Test that merge succeeds but everything remains untouched.
	err = base.Merge(bad_stack)
	if err != nil {
		t.Error(err)
	}
	checkString(t, base, "elephant, 16, 3.14, 131, gazelle, monkey")
	checkCount(t, base, 6)
	checkString(t, bad_stack, "<nil>")
	checkCount(t, bad_stack, -1)

	// Test merging a good stack on top of a bad one.
	var nil_stack *Stack
	checkString(t, nil_stack, "<nil>")
	checkCount(t, nil_stack, -1)

	// Test that merge fails and everything remains untouched.
	err = nil_stack.Merge(base)
	if err == nil {
		t.Error("unexpectedly passed merge on top of bad stack test")
	}
	checkString(t, nil_stack, "<nil>")
	checkCount(t, nil_stack, -1)
	checkString(t, base, "elephant, 16, 3.14, 131, gazelle, monkey")
	checkCount(t, base, 6)
}


// HELPERS
func checkString(t *testing.T, stack *Stack, want string) {
	if stack.String() != want {
		t.Error("stack contents are incorrect")
		t.Log("Expected:", want)
		t.Log("Received:", stack)
	}
}

func checkCount(t *testing.T, stack *Stack, want int) {
	if stack.Count() != want {
		t.Error("Incorrect length")
		t.Log("Expected:", want)
		t.Log("Received:", stack.Count())
	}
}
