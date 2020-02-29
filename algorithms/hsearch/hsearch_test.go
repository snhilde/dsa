package hsearch

import (
	"testing"
	"time"
	"math/rand"
	"sort"
)


// Build two slices of random numbers: one with our target value and one without.
// Send each slice and target value through the provided search function.
// t          testing object
// searchFunc callback search function
// iters      num of iterations to run
// length     length of slice to sort
// sorted     whether or not to sort the arrays before searching
func testSearch(t *testing.T, searchFunc func([]int, int)(int, error), iters int, length int, sorted bool) {
	for i := 0; i < iters; i++ {
		seed   := time.Now().UnixNano()
		source := rand.NewSource(seed)
		random := rand.New(source)

		target := random.Int()

		// Create our first slice, which will include our target value.
		list := make([]int, length)
		for i := 0; i < length; i++ {
			list[i] = random.Int()
		}
		index := int(random.Int31n(int32(length)))
		list[index] = target
		if sorted {
			// Sort the list and find our new target index.
			sort.Ints(list)
			index = sort.SearchInts(list, target)
		}

		// Test the first slice.
		i, err := searchFunc(list, target)
		if err != nil {
			t.Log("Expected to find target at index", index)
			t.Error(err)
		} else if i != index {
			t.Error("Incorrect index")
			t.Log("Expected index =", index)
			t.Log("Returned index =", i)
		}

		// Create our second slice, which will not include our target value.
		list = make([]int, length)
		for i := 0; i < length; i++ {
			v := target
			for v == target {
				v = random.Int()
			}
			list[i] = v
		}
		if sorted {
			sort.Ints(list)
		}

		// Test the second slice.
		i, err = searchFunc(list, target)
		if i != -1 {
			t.Error("Found unexpected target in list B")
		} else if err == nil {
			t.Error("Unexpectedly passed Test 2")
		} else if err.Error() != "Not found" {
			t.Error("Received unexpected error:")
			t.Error(err)
		}
	}
}

func TestLinearInt(t *testing.T) {
	testSearch(t, LinearInt, 100, 10000, false)
}

func TestBinaryInt(t *testing.T) {
	testSearch(t, BinaryInt, 100, 10000, true)
}
