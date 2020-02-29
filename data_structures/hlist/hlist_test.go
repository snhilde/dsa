package hlist

import (
	"testing"
	"fmt"
)


// TESTS
func TestNew(t *testing.T) {
	// Test the correct way to create a new linked list.
	good_list := New()
	if good_list == nil {
		t.Error("Failed to create new list")
	}

	// Test the incorrect way to create a new linked list.
	var bad_list *Hlist
	if bad_list != nil {
		t.Error("Unexpectedly created invalid list")
	}
}

func TestAppend(t *testing.T) {
	list := New()
	if list == nil {
		t.Error("Failed to create new list")
	}

	// Add one item. Test that item was added successfully and that list is displayed correctly.
	err := list.Append(5)
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "5")
	testLength(t, list, 1)

	// Add another item. Test that item was added successfully and that list is displayed correctly.
	err = list.Append("bananas")
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "5, bananas")
	testLength(t, list, 2)

	// Add another item. Test that item was added successfully and that list is displayed correctly.
	err = list.Append([]float64{3.14, 1.23})
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "5, bananas, [3.14 1.23]")
	testLength(t, list, 3)
}

func TestInsert(t *testing.T) {
	// First, test correct usage.
	list := New()
	if list == nil {
		t.Error("Failed to create new list")
	}

	// Add one item. Test that item was added successfully and that list is displayed correctly and length is correct.
	err := list.Insert(5, 0)
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "5")
	testLength(t, list, 1)

	// Insert another item at the beginning. Test that item was added successfully and that list is displayed correctly
	// and length is correct.
	err = list.Insert("bananas", 0)
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "bananas, 5")
	testLength(t, list, 2)

	// Add an item in the middle. Test that item was added successfully and that list is displayed correctly and length
	// is correct.
	err = list.Insert([]float64{3.14, 1.23}, 1)
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "bananas, [3.14 1.23], 5")
	testLength(t, list, 3)

	// Now, test incorrect usage.
	err = list.Insert("no way", 10)
	if err == nil {
		t.Error("Unexpectedly passed out-of-bounds test")
	}
	testString(t, list, "bananas, [3.14 1.23], 5")
	testLength(t, list, 3)

	var bad_list *Hlist
	err = bad_list.Insert("shouldn't work", 0)
	if err == nil {
		t.Error("Unexpectedly passed invalid list test")
	}
	testString(t, bad_list, "<nil>")
	testLength(t, bad_list, -1)
}

func TestPop(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Make sure the list was set up correctly.
	testString(t, list, "1, 2, 3, 4, 5")
	testLength(t, list, 5)

	// Pop the 3rd item.
	value := list.Pop(2)
	if value != 3 {
		t.Error("Error popping 3rd item")
		t.Log("Expected: 3")
		t.Log("Received:", value)
	}
	testString(t, list, "1, 2, 4, 5")
	testLength(t, list, 4)

	// Pop the new 3rd item.
	value = list.Pop(2)
	if value != 4 {
		t.Error("Error popping new 3rd item")
		t.Log("Expected: 4")
		t.Log("Received:", value)
	}
	testString(t, list, "1, 2, 5")
	testLength(t, list, 3)

	// Try to pop a non-existant item.
	value = list.Pop(10)
	if value != nil {
		t.Error("Unexpectedly passed out-of-bounds test")
		t.Log("Expected: nil")
		t.Log("Received:", value)
	}
	testString(t, list, "1, 2, 5")
	testLength(t, list, 3)

	// Try the pop operation on an invalid list.
	var bad_list *Hlist
	value = bad_list.Pop(0)
	if value != nil {
		t.Error("Unexpectedly passed invalid list test")
		t.Log("Expected: nil")
		t.Log("Received:", value)
	}
}

func TestPopMatch(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append("apples")
	list.Append(3)
	list.Append(4)
	list.Append(3.14)

	// Make sure the list was set up correctly.
	testString(t, list, "1, apples, 3, 4, 3.14")
	testLength(t, list, 5)

	// Pop the "3".
	ret := list.PopMatch(3)
	if !ret {
		t.Error("Error popping \"3\": no match")
	}
	testString(t, list, "1, apples, 4, 3.14")
	testLength(t, list, 4)

	// Pop the "1".
	ret = list.PopMatch(1)
	if !ret {
		t.Error("Error popping \"1\": no match")
	}
	testString(t, list, "apples, 4, 3.14")
	testLength(t, list, 3)

	// Pop "apples".
	ret = list.PopMatch("apples")
	if !ret {
		t.Error("Error popping \"apples\": no match")
	}
	testString(t, list, "4, 3.14")
	testLength(t, list, 2)

	// Pop the "3.14".
	ret = list.PopMatch(3.14)
	if !ret {
		t.Error("Error popping \"3.14\": no match")
	}
	testString(t, list, "4")
	testLength(t, list, 1)

	// Try to pop a non-existant item.
	ret = list.PopMatch(10)
	if ret {
		t.Error("Unexpectedly passed no-match test")
	}
	testString(t, list, "4")
	testLength(t, list, 1)

	// Try the pop operation on an invalid list.
	var bad_list *Hlist
	ret = bad_list.PopMatch(1)
	if ret {
		t.Error("Unexpectedly passed invalid list test")
	}
}

func TestIndex(t *testing.T) {
	list := New()
	list.Append("apples")
	list.Append(1)
	list.Append(3)
	list.Append(3.14)
	list.Append(4)

	// Make sure the list was set up correctly.
	testString(t, list, "apples, 1, 3, 3.14, 4")
	testLength(t, list, 5)

	// Check index of 1.
	index := list.Index(1)
	if index != 1 {
		t.Error("Incorrect index for 1")
		t.Log("Expected: 1")
		t.Log("Received:", index)
	}

	// Check index of "apples".
	index = list.Index("apples")
	if index != 0 {
		t.Error("Incorrect index for \"apples\"")
		t.Log("Expected: 0")
		t.Log("Received:", index)
	}

	// Check index of 1.
	index = list.Index(3.14)
	if index != 3 {
		t.Error("Incorrect index for 3.14")
		t.Log("Expected: 3")
		t.Log("Received:", index)
	}

	// Try to find a non-existant item.
	index = list.Index(10)
	if index != -1 {
		t.Error("Unexpectedly passed no-match test")
	}

	// Try the Index() operation on an invalid list.
	var bad_list *Hlist
	index = bad_list.Index(1)
	if index != -1 {
		t.Error("Unexpectedly passed invalid list test")
	}
}

func TestExists(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append(3)
	list.Append("apples")
	list.Append(3.14)
	list.Append(4)

	// Make sure the list was set up correctly.
	testString(t, list, "1, 3, apples, 3.14, 4")
	testLength(t, list, 5)

	// Check that 3 exists.
	ret := list.Exists(3)
	if !ret {
		t.Error("Expected match for 3, but didn't find it")
	}

	// Check that "apples" exists.
	ret = list.Exists("apples")
	if !ret {
		t.Error("Expected match for \"apples\", but didn't find it")
	}

	// Check that 3.14 exists.
	ret = list.Exists(3.14)
	if !ret {
		t.Error("Expected match for 3,14, but didn't find it")
	}

	// Try to find a non-existant item.
	ret = list.Exists(10)
	if ret {
		t.Error("Unexpectedly passed no-match test")
	}

	// Try the Exists() operation on an invalid list.
	var bad_list *Hlist
	ret = bad_list.PopMatch(1)
	if ret {
		t.Error("Unexpectedly passed invalid list test")
	}
}

func TestMerge(t *testing.T) {
	// Test merging two good lists togther.
	list := New()
	for i := 0; i < 5; i++ {
		list.Append(i)
	}
	testString(t, list, "0, 1, 2, 3, 4")
	testLength(t, list, 5)

	addition := New()
	for i := 5; i < 10; i++ {
		addition.Append(i)
	}
	testString(t, addition, "5, 6, 7, 8, 9")
	testLength(t, addition, 5)

	err := list.Merge(addition)
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "0, 1, 2, 3, 4, 5, 6, 7, 8, 9")
	testString(t, addition, "<empty>")
	testLength(t, addition, 0)

	// Test merging a good list and a bad list.
	list = New()
	for i := 0; i < 5; i++ {
		list.Append(i)
	}
	testString(t, list, "0, 1, 2, 3, 4")
	testLength(t, list, 5)

	var bad_list *Hlist
	testString(t, bad_list, "<nil>")
	testLength(t, bad_list, -1)

	err = list.Merge(bad_list)
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "0, 1, 2, 3, 4")
	testLength(t, list, 5)

	// Test merging a bad list with a good list.
	// var bad_list *Hlist
	testString(t, bad_list, "<nil>")
	testLength(t, bad_list, -1)

	list = New()
	for i := 0; i < 5; i++ {
		list.Append(i)
	}
	testString(t, list, "0, 1, 2, 3, 4")
	testLength(t, list, 5)

	err = bad_list.Merge(list)
	if err == nil {
		t.Error("Unexpectedly passed bad merge")
	}
	testString(t, bad_list, "<nil>")
	testLength(t, bad_list, -1)
	testString(t, list, "0, 1, 2, 3, 4")
	testLength(t, list, 5)
}

func TestSortInt(t *testing.T) {
	// Test a power of two.
	list := New()
	for i := 1; i <= 8; i++ {
		list.Insert(i, 0)
	}
	testString(t, list, "8, 7, 6, 5, 4, 3, 2, 1")
	testLength(t, list, 8)
	err := list.SortInt()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "1, 2, 3, 4, 5, 6, 7, 8")
	testLength(t, list, 8)

	// Test a block with less than one full stack.
	list = New()
	for i := 1; i <= 9; i++ {
		list.Insert(i, 0)
	}
	testString(t, list, "9, 8, 7, 6, 5, 4, 3, 2, 1")
	testLength(t, list, 9)
	err = list.SortInt()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "1, 2, 3, 4, 5, 6, 7, 8, 9")
	testLength(t, list, 9)

	// Test a block with more than one full stack but less than a full block
	list = New()
	for i := 1; i <= 11; i++ {
		list.Insert(i, 0)
	}
	testString(t, list, "11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1")
	testLength(t, list, 11)
	err = list.SortInt()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11")
	testLength(t, list, 11)

	// Test small list sizes.
	list = New()
	list.Append(1)
	list.Append(0)
	testString(t, list, "1, 0")
	testLength(t, list, 2)
	err = list.SortInt()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "0, 1")
	testLength(t, list, 2)

	list = New()
	list.Append(1)
	testString(t, list, "1")
	testLength(t, list, 1)
	err = list.SortInt()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "1")
	testLength(t, list, 1)

	// Test invalid lists.
	list = New()
	testString(t, list, "<empty>")
	testLength(t, list, 0)
	err = list.SortInt()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "<empty>")
	testLength(t, list, 0)

	var bad_list *Hlist
	testString(t, bad_list, "<nil>")
	testLength(t, bad_list, -1)
	err = bad_list.SortInt()
	if err == nil {
		t.Error("Unexpectedly passed bad list sort")
	}
	testString(t, bad_list, "<nil>")
	testLength(t, bad_list, -1)
}

func TestSortStr(t *testing.T) {
	// Test uniform length and no same characters.
	list := New()
	list.Append("ccc")
	list.Append("bbb")
	list.Append("aaa")
	testString(t, list, "ccc, bbb, aaa")
	testLength(t, list, 3)
	err := list.SortStr()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "aaa, bbb, ccc")
	testLength(t, list, 3)

	// Test variable lengths and no same characters.
	list = New()
	list.Append("ccccc")
	list.Append("bbb")
	list.Append("aaaaaaaaaaa")
	testString(t, list, "ccccc, bbb, aaaaaaaaaaa")
	testLength(t, list, 3)
	err = list.SortStr()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "aaaaaaaaaaa, bbb, ccccc")
	testLength(t, list, 3)

	// Test uniform length and similar characters.
	list = New()
	list.Append("cabc")
	list.Append("caab")
	list.Append("abab")
	testString(t, list, "cabc, caab, abab")
	testLength(t, list, 3)
	err = list.SortStr()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "abab, caab, cabc")
	testLength(t, list, 3)

	// Test variable length and similar characters.
	list = New()
	list.Append("cababcabbcbababca")
	list.Append("cabacbcabacb")
	list.Append("cabababacba")
	testString(t, list, "cababcabbcbababca, cabacbcabacb, cabababacba")
	testLength(t, list, 3)
	err = list.SortStr()
	if err != nil {
		t.Error(err)
	}
	testString(t, list, "cabababacba, cababcabbcbababca, cabacbcabacb")
	testLength(t, list, 3)
}



// HELPERS
func testString(t *testing.T, list *Hlist, want string) {
	if fmt.Sprintf("%v", list) != want {
		t.Error("List contents are incorrect")
		t.Log("Expected:", want)
		t.Log("Received:", list)
	}
}

func testLength(t *testing.T, list *Hlist, want int) {
	if list.Length() != want {
		t.Error("Incorrect length")
		t.Log("Expected:", want)
		t.Log("Received:", list.Length())
	}
}
