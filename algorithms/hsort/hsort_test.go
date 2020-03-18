package hsort

import (
	"testing"
	"time"
	"math/rand"
	"sort"
)


// Build a slice of random numbers and sort it with the provided sorting function.
// t        testing object
// sortFunc callback sort function
// iters    num of iterations to run
// length   length of slice to sort
// isHash   if true, limit the range of random values to not overload a hashing algorithm
func testSort(t *testing.T, sortFunc func(interface{}) error, iters int, length int, isHash bool) {
	for i := 0; i < iters; i++ {
		seed   := time.Now().UnixNano()
		source := rand.NewSource(seed)
		random := rand.New(source)

		// Populate the slice with random values.
		list := make([]int, length)
		for i := 0; i < length; i++ {
			if isHash {
				list[i] = random.Intn(1e6)
			} else {
				list[i] = random.Int()
			}
		}

		// Sort the slice using the provided algorithm.
		listCopy := make([]int, length)
		copy(listCopy, list)
		err := sortFunc(list)
		if err != nil {
			t.Log("Sorting failed:")
			t.Error(err)
		}

		// Check that the sorting algorithm was correct.
		sort.Ints(listCopy)
		for i, v := range list {
			if v != listCopy[i] {
				t.Error("Values at index", i, "differ")
				t.Log("should be:", listCopy[i])
				t.Log("really is:", v)
			}
		}
	}
}

func testSortInt(t *testing.T, sortFunc func([]int) error, iters int, length int, isHash bool) {
	for i := 0; i < iters; i++ {
		seed   := time.Now().UnixNano()
		source := rand.NewSource(seed)
		random := rand.New(source)

		// Populate the slice with random values.
		list := make([]int, length)
		for i := 0; i < length; i++ {
			if isHash {
				list[i] = random.Intn(1e6)
			} else {
				list[i] = random.Int()
			}
		}

		// Sort the slice using the provided algorithm.
		listCopy := make([]int, length)
		copy(listCopy, list)
		err := sortFunc(list)
		if err != nil {
			t.Log("Sorting failed:")
			t.Error(err)
		}

		// Check that the sorting algorithm was correct.
		sort.Ints(listCopy)
		for i, v := range list {
			if v != listCopy[i] {
				t.Error("Values at index", i, "differ")
				t.Log("should be:", listCopy[i])
				t.Log("really is:", v)
			}
		}
	}
}

func TestInsertionInt(t *testing.T) {
	testSort(t, Insertion, 100, 1000, false)
	testSortInt(t, InsertionInt, 100, 1000, false)
}

func TestSelectionInt(t *testing.T) {
	testSort(t, Selection, 100, 1000, false)
	testSortInt(t, SelectionInt, 100, 1000, false)
}

func TestMergeInt(t *testing.T) {
	testSortInt(t, MergeInt, 100, 1000, false)
}

func TestMergeIntOptimized(t *testing.T) {
	testSortInt(t, MergeIntOptimized, 100, 1000, false)
}

func TestHashInt(t *testing.T) {
	testSortInt(t, HashInt, 100, 1000, true)
}

func TestBubbleInt(t *testing.T) {
	testSort(t, Bubble, 100, 1000, false)
	testSortInt(t, BubbleInt, 100, 1000, false)
}
