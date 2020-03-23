package hlist

import (
	"testing"
)


// TESTS
func TestBadPtr(t *testing.T) {
	// Test that using a non-initialized list is handled correctly.
	var l *List

    // Test String().
	if s := l.String(); s != "<nil>" {
		t.Error("unexpectedly passed String() test with bad pointer")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}

    // Test Length().
	if n := l.Length(); n != -1 {
		t.Error("unexpectedly passed Length() test with bad pointer")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", n)
	}

    // Test Insert().
	if err := l.Insert(0, "item"); err == nil {
		t.Error("unexpectedly passed Insert() test with bad pointer")
	}

    // Test Append().
	if err := l.Append("item"); err == nil {
		t.Error("unexpectedly passed Append() test with bad pointer")
	}

    // Test Remove().
	if v := l.Remove(0); v != nil {
		t.Error("unexpectedly passed Remove() test with bad pointer")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

    // Test RemoveMatch().
	if err := l.RemoveMatch("item"); err == nil {
		t.Error("unexpectedly passed RemoveMatch() test with bad pointer")
	}

    // Test Index().
	if i := l.Index("item"); i != -1 {
		t.Error("unexpectedly passed Index() test with bad pointer")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", i)
	}

    // Test Exists().
	if b := l.Exists("item"); b != false {
		t.Error("unexpectedly passed Exists() test with bad pointer")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", b)
	}

    // Test Copy().
	if _, err := l.Copy(); err == nil {
		t.Error("unexpectedly passed Copy() test with bad pointer")
	}

    // Test Merge().
	if err := l.Merge(New()); err == nil {
		t.Error("unexpectedly passed Merge() test with bad pointer")
	}

    // Test Clear().
	if err := l.Clear(); err == nil {
		t.Error("unexpectedly passed Clear() test with bad pointer")
	}

    // Test Sort().
	if err := l.Sort(func(l, r interface{}) bool {return false}); err == nil {
		t.Error("unexpectedly passed Sort() test with bad pointer")
	}

    // Test SortInt().
	if err := l.SortInt(); err == nil {
		t.Error("unexpectedly passed SortInt() test with bad pointer")
	}

    // Test SortStr().
	if err := l.SortStr(); err == nil {
		t.Error("unexpectedly passed SortStr() test with bad pointer")
	}
}

func TestBadArgs(t *testing.T) {
	// Test that passing bad values to methods is handled correctly.
	l := New()
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Insert().
	if err := l.Insert(-1, "item"); err == nil {
		t.Error("unexpectedly passed Insert() test for negative index")
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Remove().
	if v := l.Remove(-1); v != nil {
		t.Error("unexpectedly passed Remove() test for negative index")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Remove().
	if v := l.Remove(0); v != nil {
		t.Error("unexpectedly passed Remove() test for empty list")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Remove().
	if v := l.Remove(100); v != nil {
		t.Error("unexpectedly passed Remove() test for out-of-range index")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Sort().
	if err := l.Sort(nil); err == nil {
		t.Error("unexpectedly passed Sort() test for missing sort cb")
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
}

func TestNew(t *testing.T) {
	if l := New(); l == nil {
		t.Error("Failed to create new list")
	}
}


// HELPERS
func checkString(t *testing.T, l *List, want string) {
	if l.String() != want {
		t.Error("List contents are incorrect")
		t.Log("Expected:", want)
		t.Log("Received:", l)
	}
}

func checkLength(t *testing.T, l *List, want int) {
	if l.Length() != want {
		t.Error("Incorrect length")
		t.Log("Expected:", want)
		t.Log("Received:", l.Length())
	}
}
