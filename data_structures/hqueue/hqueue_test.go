package hqueue_test

import (
	"reflect"
	"testing"

	"github.com/snhilde/dsa/data_structures/hqueue"
)

func TestBadPtr(t *testing.T) {
	var q *hqueue.Queue
	checkString(t, q, "<nil>")
	checkCount(t, q, -1)

	// Test Add().
	if err := q.Add("value"); err == nil {
		t.Error("unexpectedly passed Add() test with bad pointer")
	}

	// Test Pop().
	if v := q.Pop(); v != nil {
		t.Error("unexpectedly passed Pop() test with bad pointer")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Test Count().
	if n := q.Count(); n != -1 {
		t.Error("unexpectedly passed Count() test with bad pointer")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", n)
	}

	// Test Copy().
	if _, err := q.Copy(); err == nil {
		t.Error("unexpectedly passed Copy() test with bad pointer")
	}

	// Test Merge().
	if err := q.Merge(hqueue.New()); err == nil {
		t.Error("unexpectedly passed Merge() test with bad pointer")
	}

	// Test Clear().
	if err := q.Clear(); err == nil {
		t.Error("unexpectedly passed Clear() test with bad pointer")
	}

	// Test Same().
	if q.Same(hqueue.New()) {
		t.Error("unexpectedly passed Same() test with bad pointer")
	}

	// Test String().
	if s := q.String(); s != "<nil>" {
		t.Error("unexpectedly passed String() test with bad pointer")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}
}

func TestNew(t *testing.T) {
	q := hqueue.New()
	if q == nil {
		t.Error("new queue unexpectedly nil")
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)

}

func TestAdd(t *testing.T) {
	q := hqueue.New()

	// Testing adding an int.
	if err := q.Add(5); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5")
	checkCount(t, q, 1)

	// Test adding a string.
	if err := q.Add("kangaroo"); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo")
	checkCount(t, q, 2)

	// Testing adding a float.
	if err := q.Add(3.1415); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo, 3.1415")
	checkCount(t, q, 3)

	// Test out adding multiple items at once.
	if err := q.Add("a", "b", 3); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo, 3.1415, a, b, 3")
	checkCount(t, q, 6)

	// Test adding a slice.
	if err := q.Add([]int{1, 2, 3}); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo, 3.1415, a, b, 3, [1 2 3]")
	checkCount(t, q, 7)

	// Testing adding an empty queue.
	if err := q.Add(hqueue.New()); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo, 3.1415, a, b, 3, [1 2 3], <empty>")
	checkCount(t, q, 8)

	// Test adding a non-empty queue.
	a := hqueue.New()
	a.Add("orange, apple, banana")
	if err := q.Add(a); err != nil {
		t.Error(err)
	}
	checkString(t, q, "5, kangaroo, 3.1415, a, b, 3, [1 2 3], <empty>, orange, apple, banana")
	checkCount(t, q, 9)

	// Test adding queue to itself.
	if err := q.Add(q); err == nil {
		t.Error("unexpectedly passed adding queue to itself")
	}
	checkString(t, q, "5, kangaroo, 3.1415, a, b, 3, [1 2 3], <empty>, orange, apple, banana")
	checkCount(t, q, 9)
}

func TestPop(t *testing.T) {
	q := hqueue.New()

	// Add some items first.
	q.Add("sizzle")
	q.Add(1e5)
	q.Add(3.1415)
	q.Add(15)
	checkString(t, q, "sizzle, 100000, 3.1415, 15")
	checkCount(t, q, 4)

	// Test out popping the items.
	if val := q.Pop(); val != "sizzle" {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: sizzle")
		t.Log("\tExpected:", val)
	}
	checkString(t, q, "100000, 3.1415, 15")
	checkCount(t, q, 3)

	if val := q.Pop(); val != 1e5 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: 1e5")
		t.Log("\tExpected:", val)
	}
	checkString(t, q, "3.1415, 15")
	checkCount(t, q, 2)

	if val := q.Pop(); val != 3.1415 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: 3.1415")
		t.Log("\tExpected:", val)
	}
	checkString(t, q, "15")
	checkCount(t, q, 1)

	if val := q.Pop(); val != 15 {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: 15")
		t.Log("\tExpected:", val)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)

	// Test popping from an empty queue.
	if val := q.Pop(); val != nil {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected: nil")
		t.Log("\tExpected:", val)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)

	// Test popping a slice.
	s := []int{1, 2, 3}
	q.Add(s)
	checkString(t, q, "[1 2 3]")
	checkCount(t, q, 1)

	if p := q.Pop().([]int); !reflect.DeepEqual(s, p) {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected:", s)
		t.Log("\tExpected:", p)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)

	// Test popping a queue.
	a := hqueue.New()
	a.Add("orange, apple, banana")
	q.Add(a)
	checkString(t, q, "orange, apple, banana")
	checkCount(t, q, 1)

	if val := q.Pop(); val.(*hqueue.Queue) != a {
		t.Error("Incorrect value from pop")
		t.Log("\tExpected:", a)
		t.Log("\tExpected:", val)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)
}

func TestCopy(t *testing.T) {
	q := hqueue.New()

	// Copy an empty queue.
	nq, err := q.Copy()
	if err != nil {
		t.Error(err)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)
	checkString(t, nq, "<empty>")
	checkCount(t, nq, 0)

	// Copy a non-empty queue.
	q.Add("sizzle", 1e5, 3.1415, 15)
	checkString(t, q, "sizzle, 100000, 3.1415, 15")
	checkCount(t, q, 4)

	nq, err = q.Copy()
	if err != nil {
		t.Error(err)
	}
	checkString(t, q, "sizzle, 100000, 3.1415, 15")
	checkCount(t, q, 4)
	checkString(t, nq, "sizzle, 100000, 3.1415, 15")
	checkCount(t, nq, 4)

}

func TestMerge(t *testing.T) {
	// Create two queues and merge them.
	q := hqueue.New()
	q.Add("monkey", "gazelle", 131)
	checkString(t, q, "monkey, gazelle, 131")
	checkCount(t, q, 3)

	tmp := hqueue.New()
	tmp.Add(3.1415, 16, []uint{5, 6, 7})
	checkString(t, tmp, "3.1415, 16, [5 6 7]")
	checkCount(t, tmp, 3)

	// Merge and check that tmp was added behind q and that tmp was emptied out.
	if err := q.Merge(tmp); err != nil {
		t.Error(err)
	}
	checkString(t, q, "monkey, gazelle, 131, 3.1415, 16, [5 6 7]")
	checkCount(t, q, 6)
	checkString(t, tmp, "<empty>")
	checkCount(t, tmp, 0)

	// Test merging an invalid queue on top of a good one. Merge should succeed, but everything should remain untouched.
	var nq *hqueue.Queue
	checkString(t, nq, "<nil>")
	checkCount(t, nq, -1)

	if err := q.Merge(nq); err != nil {
		t.Error(err)
	}
	checkString(t, q, "monkey, gazelle, 131, 3.1415, 16, [5 6 7]")
	checkCount(t, q, 6)
	checkString(t, nq, "<nil>")
	checkCount(t, nq, -1)

	// Test merging a good queue on top of a bad one. Merge should fail, and everything should remain untouched.
	if err := nq.Merge(q); err == nil {
		t.Error("unexpectedly passed merge on top of bad queue test")
	}
	checkString(t, nq, "<nil>")
	checkCount(t, nq, -1)
	checkString(t, q, "monkey, gazelle, 131, 3.1415, 16, [5 6 7]")
	checkCount(t, q, 6)
}

func TestClear(t *testing.T) {
	q := hqueue.New()

	// Add some items first.
	q.Add("kangaroo", 5, 3.1415)
	checkString(t, q, "kangaroo, 5, 3.1415")
	checkCount(t, q, 3)

	// Test out clearing the queue.
	if err := q.Clear(); err != nil {
		t.Error(err)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)

	// Test out clearing an empty queue.
	if err := q.Clear(); err != nil {
		t.Error(err)
	}
	checkString(t, q, "<empty>")
	checkCount(t, q, 0)
}

func TestSame(t *testing.T) {
	q1 := hqueue.New()
	q2 := q1

	// Test checking against nothing.
	if q1.Same(nil) {
		t.Error("unexpectedly passed Same() test for missing arg")
	}

	// Test an empty queue.
	if !q1.Same(q2) {
		t.Error("Queues differ (test #1)")
	}
	if !q2.Same(q1) {
		t.Error("Queues differ (test #2)")
	}
	checkString(t, q1, "<empty>")
	checkCount(t, q1, 0)
	checkString(t, q2, "<empty>")
	checkCount(t, q2, 0)

	// And an item and test again.
	q1.Add("item")
	if !q1.Same(q2) {
		t.Error("Queues differ (test #3)")
	}
	if !q2.Same(q1) {
		t.Error("Queues differ (test #4)")
	}
	checkString(t, q1, "item")
	checkCount(t, q1, 1)
	checkString(t, q2, "item")
	checkCount(t, q2, 1)

	// And multiple items to the 2nd reference.
	q2.Add(3.14, []int{1, 2, 3}, 5)
	if !q1.Same(q2) {
		t.Error("Queues differ (test #5)")
	}
	if !q2.Same(q1) {
		t.Error("Queues differ (test #6)")
	}
	checkString(t, q1, "item, 3.14, [1 2 3], 5")
	checkCount(t, q1, 4)
	checkString(t, q2, "item, 3.14, [1 2 3], 5")
	checkCount(t, q2, 4)

	// Pop an item and test again.
	q1.Pop()
	if !q1.Same(q2) {
		t.Error("Queues differ (test #7)")
	}
	if !q2.Same(q1) {
		t.Error("Queues differ (test #8)")
	}
	checkString(t, q1, "3.14, [1 2 3], 5")
	checkCount(t, q1, 3)
	checkString(t, q2, "3.14, [1 2 3], 5")
	checkCount(t, q2, 3)

	// Test copying, and make sure that the old queues and the new queue are no longer the same.
	q3, err := q1.Copy()
	if err != nil {
		t.Error(err)
	}
	if q1.Same(q3) {
		t.Error("Queues are unexpectedly the same (test #9)")
	}
	if q2.Same(q3) {
		t.Error("Queues are unexpectedly the same (test #10)")
	}

	// Test clearing a list.
	q2.Clear()
	if !q1.Same(q2) {
		t.Error("Queues differ (test #11)")
	}
	if !q2.Same(q1) {
		t.Error("Queues differ (test #12)")
	}
	checkString(t, q1, "<empty>")
	checkCount(t, q1, 0)
	checkString(t, q2, "<empty>")
	checkCount(t, q2, 0)

	// Test reassigning to make a same list.
	q3 = q2
	if !q1.Same(q3) {
		t.Error("Queues differ (test #13)")
	}
	if !q2.Same(q3) {
		t.Error("Queues differ (test #14)")
	}

	// Make two Queues that have the same contents but are not the same underlying Queues.
	q1 = hqueue.New()
	q1.Add("apple", "banana", "carrot")
	q2 = hqueue.New()
	q2.Add("apple", "banana", "carrot")
	if q1.Same(q2) {
		t.Error("Queues are unexpectedly the same (test #15)")
	}
	if q2.Same(q1) {
		t.Error("Queues are unexpectedly the same (test #16)")
	}
}

func checkString(t *testing.T, q *hqueue.Queue, want string) {
	if q.String() != want {
		t.Error("queue contents are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", q)
	}
}

func checkCount(t *testing.T, q *hqueue.Queue, want int) {
	if q.Count() != want {
		t.Error("Incorrect length")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", q.Count())
	}
}
