package hqueue

import (
	"testing"
)

// TESTS
func TestBadPtr(t *testing.T) {
	var q *Queue

	// Test Add().
	if err := q.Add("value"); err == nil {
		t.Error("unexpectedly passed Add() test with bad pointer")
	}

	// Test Pop().
	if v := q.Pop(); v != nil {
		t.Error("unexpectedly passed Pop() test with bad pointer")
		t.Log("Expected: nil")
		t.Log("Received:", v)
	}

	// Test Count().
	if n := q.Count(); n != -1 {
		t.Error("unexpectedly passed Count() test with bad pointer")
		t.Log("Expected: -1")
		t.Log("Received:", n)
	}

	// Test Clear().
	if err := q.Clear(); err == nil {
		t.Error("unexpectedly passed Clear() test with bad pointer")
	}

	// Test Merge().
	if err := q.Merge(New()); err == nil {
		t.Error("unexpectedly passed Merge() test with bad pointer")
	}

	// Test String().
	if s := q.String(); s != "<nil>" {
		t.Error("unexpectedly passed String() test with bad pointer")
		t.Log("Expected: <nil>")
		t.Log("Received:", s)
	}
}

func TestNew(t *testing.T) {
	q := New()
	if q == nil {
		t.Error("new queue unexpectedly nil")
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)

}

func TestAdd(t *testing.T) {
	q := New()

	// Test out adding some items to the queue.
	if err := q.Add(5); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5")
	checkCount(t, q, 1)

	if err := q.Add("kangaroo"); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo")
	checkCount(t, q, 2)

	if err := q.Add(3.1415); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo, 3.1415")
	checkCount(t, q, 3)

	if err := q.Add(New()); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo, 3.1415, <empty>")
	checkCount(t, q, 4)

	// TODO: test adding a slice and a non-empty queue
}

func TestPop(t *testing.T) {
	q := New()

	// Add some items first.
	q.Add("sizzle", 1e5, 3.1415, 15)
	checkString(t, q, "sizzle, 100000, 3.1415, 15")
	checkCount(t, q, 4)

	// Test out popping the items.
	if val := q.Pop(); val != "sizzle" {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected sizzle")
		t.Log("\tReceived", val)
	}
	checkString(t, q, "100000, 3.1415, 15")
	checkCount(t, q, 3)

	if val := q.Pop(); val != 1e5 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected 1e5")
		t.Log("\tReceived", val)
	}
	checkString(t, q, "3.1415, 15")
	checkCount(t, q, 2)

	if val := q.Pop(); val != 3.1415 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected 3.1415")
		t.Log("\tReceived", val)
	}
	checkString(t, q, "15")
	checkCount(t, q, 1)

	if val := q.Pop(); val != 15 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected 15")
		t.Log("\tReceived", val)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)

	if val := q.Pop(); val != nil {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected nil")
		t.Log("\tReceived", val)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)

	// TODO: add tests popping slices and queues
}

func TestClear(t *testing.T) {
	q := New()

	// Add some items first.
	q.Add("kangaroo", 5, 3.1415)
	checkString(t, q, "kangaroo, 5, 3.1415")
	checkCount(t, q, 3)

	// Test out clearing the queue.
	err := q.Clear()
	if err != nil {
		t.Error(err)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)
}

func TestMerge(t *testing.T) {
	// Create two queues and merge them.
	q := New()
	q.Add("monkey", "gazelle", 131)
	checkString(t, q, "monkey, gazelle, 131")
	checkCount(t, q, 3)

	tmpQ := New()
	tmpQ.Add(3.1415, 16, "elephant")
	checkString(t, tmpQ, "3.1415, 16, elephant")
	checkCount(t, tmpQ, 3)

	// Merge and check that tmpQ was added behind q and that tmpQ was emptied out.
	err := q.Merge(tmpQ)
	if err != nil {
		t.Error(err)
	}
	checkString(t, q, "monkey, gazelle, 131, 3.1415, 16, elephant")
	checkCount(t, q, 6)
	checkString(t, tmpQ, "<empty>")
	checkCount(t, tmpQ, 0)

	// Test merging a bad queue on top of a good one. Merge should succeed, but everything should remain untouched.
	var bq *Queue
	if err = q.Merge(bq); err != nil {
		t.Error(err)
	}
	checkString(t, q, "monkey, gazelle, 131, 3.1415, 16, elephant")
	checkCount(t, q, 6)
	checkString(t, bq, "<nil>")
	checkCount(t, bq, -1)

	// Test merging a good queue on top of a bad one. Merge should fail, and everything should remain untouched.
	if err = bq.Merge(q); err == nil {
		t.Error("unexpectedly passed merge on top of bad queue test")
	}
	checkString(t, bq, "<nil>")
	checkCount(t, bq, -1)
	checkString(t, q, "monkey, gazelle, 131, 3.1415, 16, elephant")
	checkCount(t, q, 6)
}

// HELPERS
func checkString(t *testing.T, q *Queue, want string) {
	if q.String() != want {
		t.Error("queue contents are incorrect")
		t.Log("Expected:", want)
		t.Log("Received:", q)
	}
}

func checkCount(t *testing.T, q *Queue, want int) {
	if q.Count() != want {
		t.Error("Incorrect length")
		t.Log("Expected:", want)
		t.Log("Received:", q.Count())
	}
}
