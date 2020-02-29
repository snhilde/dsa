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
func testSort(t *testing.T, sortFunc func([]int) error, iters int, length int) {
	for i := 0; i < iters; i++ {
		seed   := time.Now().UnixNano()
		source := rand.NewSource(seed)
		random := rand.New(source)

		// Populate the slice with random values.
		list := make([]int, length)
		for i := 0; i < length; i++ {
			list[i] = random.Int()
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
	testSort(t, InsertionInt, 100, 10000)
}

func TestSelectionInt(t *testing.T) {
	testSort(t, SelectionInt, 100, 10000)
}

func TestMergeInt(t *testing.T) {
	testSort(t, MergeInt, 100, 10000)
}

func TestMergeIntOptimized(t *testing.T) {
	testSort(t, MergeIntOptimized, 100, 10000)
}

func TestHashInt(t *testing.T) {
	testSort(t, HashInt, 100, 10000)
}

func TestBubbleInt(t *testing.T) {
	testSort(t, BubbleInt, 100, 10000)
}
