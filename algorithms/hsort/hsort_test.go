package hsort

import (
	"testing"
	"time"
	"math/rand"
	"sort"
)


// sorter is the main interface for this test package. It defines four functions:
// 1. Build, which will build a list to sort
// 2. Sort, which will sort the list using the sorting function in development
// 3. SortStd, which will sort the list using a standard and accepted sorting function
// 4. Cmp, which will compare the two lists and determine if they are the same or not.
type sorter interface {
	Build(length int, isHash bool)
	Sort()
	SortStd()
	Cmp() bool
}

type intSort struct {
	dev []int
	std []int
}

type uintSort struct {
	dev []uint
	std []uint
}

type floatSort struct {
	dev []float32
	std []float32
}

type boolSort struct {
	dev []bool
	std []bool
}

type stringSort struct {
	dev []string
	std []string
}


// Build a slice of random numbers and sort it with the provided sorting function.
// t        testing object
// sortFunc callback sort function
// iters    num of iterations to run
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


// Build slices of various types.
// length: length of slice to sort
// isHash: if true, limit the range of random values to not overload a hashing algorithm

func (s *intSort) Build(length int, isHash bool) {
	seed   := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	s.dev := make([]int, length)
	s.std := make([]int, length)

	for i := 0; i < length; i++ {
		if isHash {
			s.dev[i] = random.Intn(1e6)
		} else {
			s.dev[i] = random.Int()
		}
	}
}

func (s *uintSort) Build(length int, isHash bool) {
	seed   := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	s.dev := make([]int, length)
	s.std := make([]int, length)

	for i := 0; i < length; i++ {
		if isHash {
			s.dev[i] = uint(random.Intn(1e6))
		} else {
			s.dev[i] = uint(random.Uint32())
		}
	}
}

func (s *floatSort) Build(length int, isHash bool) {
	seed   := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	s.dev := make([]int, length)
	s.std := make([]int, length)

	for i := 0; i < length; i++ {
		if isHash {
			s.dev[i] = float32(random.Intn(1e6)) * random.Float32()
		} else {
			s.dev[i] = float32(random.Int()) * random.Float32()
		}
	}
}

func (s *boolSort) Build(length int, isHash bool) {
	seed   := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	s.dev := make([]int, length)
	s.std := make([]int, length)

	for i := 0; i < length; i++ {
		r := random.Int()
		if r % 2 == 1 {
			s.dev[i] = true
		} else {
			s.dev[i] = false
		}
	}
}

func (s *stringSort) Build(length int, isHash bool) {
	seed   := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	s.dev := make([]int, length)
	s.std := make([]int, length)

	for i := 0; i < length; i++ {
		l := 1
		for l > 0 {
			l := random.Intn(32)
		}
		s := make([]byte, l)
		for j := 0; j < l; j++ {
			n := random.Intn(93)
			s[j] = n + 33
		}
		s.dev[i] = string(s)
	}
}
