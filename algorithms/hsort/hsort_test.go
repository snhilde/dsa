package hsort

import (
	"testing"
	"time"
	"math/rand"
	"sort"
)


// sorter is the main interface for this test package. It allows for different types to be used with the same sorting
// test harness.
type sorter interface {
	// Build will build a list to sort. If isHash is true, the range of values should be kept low to allow for hash sorting.
	Build(length int, isHash bool)

	// Sort will sort the list using the sorting function in development.
	Sort() error

	// SortStd will sort the list using a standard and accepted sorting function.
	SortStd()

	// Cmp will compare the two lists and determine if they are the same or not.
	Cmp(*testing.T) bool
}

type intSort struct {
	dev  []int
	std  []int
	sort   func(interface{}) error
}

type uintSort struct {
	dev  []uint
	std  []int
	sort   func(interface{}) error
}

type floatSort struct {
	dev  []float64
	std  []float64
	sort   func(interface{}) error
}

type boolSort struct {
	dev  []bool
	std  []int
	sort   func(interface{}) error
}

type stringSort struct {
	dev  []string
	std  []string
	sort   func(interface{}) error
}


// Build a slice of random numbers, and sort it with the provided sorting function.
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


func (s *intSort) Build(length int, isHash bool) {
	r := newRand()

	s.dev = make([]int, length)
	s.std = make([]int, length)

	for i := 0; i < length; i++ {
		if isHash {
			s.dev[i] = r.Intn(1e6)
		} else {
			s.dev[i] = r.Int()
		}

		if (r.Int() % 2) == 0 {
			s.dev[i] *= -1
		}
		s.std[i] = s.dev[i]
	}
}

func (s *intSort) Sort() error {
	return s.sort(s.dev)
}

func (s *intSort) SortStd() {
	sort.Ints(s.std)
}

func (s *intSort) Cmp(t *testing.T) bool {
	good := true
	for i, v := range s.dev {
		if v != s.std[i] {
			good = false
			t.Error("Values at index", i, "differ")
			t.Log("should be:", s.std[i])
			t.Log("really is:", v)
		}
	}

	return good
}


func (s *uintSort) Build(length int, isHash bool) {
	r := newRand()

	// std will be an int slice so we can use package sort's Ints() function.
	s.dev = make([]uint, length)
	s.std = make([]int, length)

	for i := 0; i < length; i++ {
		if isHash {
			s.dev[i] = uint(r.Intn(1e6))
		} else {
			s.dev[i] = uint(r.Uint32())
		}
		s.std[i] = int(s.dev[i])
	}
}

func (s *uintSort) Sort() error {
	return s.sort(s.dev)
}

func (s *uintSort) SortStd() {
	sort.Ints(s.std)
}

func (s *uintSort) Cmp(t *testing.T) bool {
	good := true
	for i, v := range s.dev {
		if v != uint(s.std[i]) {
			good = false
			t.Error("Values at index", i, "differ")
			t.Log("should be:", s.std[i])
			t.Log("really is:", v)
		}
	}

	return good
}


func (s *floatSort) Build(length int, isHash bool) {
	r := newRand()

	s.dev = make([]float64, length)
	s.std = make([]float64, length)

	for i := 0; i < length; i++ {
		if isHash {
			s.dev[i] = float64(r.Intn(1e6)) * r.Float64()
		} else {
			s.dev[i] = float64(r.Int()) * r.Float64()
		}

		if (r.Int() % 2) == 0 {
			s.dev[i] *= -1
		}
		s.std[i] = s.dev[i]
	}
}

func (s *floatSort) Sort() error {
	return s.sort(s.dev)
}

func (s *floatSort) SortStd() {
	sort.Float64s(s.std)
}

func (s *floatSort) Cmp(t *testing.T) bool {
	good := true
	for i, v := range s.dev {
		if v != s.std[i] {
			good = false
			t.Error("Values at index", i, "differ")
			t.Log("should be:", s.std[i])
			t.Log("really is:", v)
		}
	}

	return good
}


func (s *boolSort) Build(length int, isHash bool) {
	r := newRand()

	s.dev = make([]bool, length)
	s.std = make([]int, length)

	for i := 0; i < length; i++ {
		r := r.Int()
		if r % 2 == 1 {
			s.dev[i] = true
			s.std[i] = 1
		} else {
			s.dev[i] = false
			s.std[i] = 0
		}
	}
}

func (s *boolSort) Sort() error {
	return s.sort(s.dev)
}

func (s *boolSort) SortStd() {
	sort.Ints(s.std)
}

func (s *boolSort) Cmp(t *testing.T) bool {
	good := true
	for i, v := range s.dev {
		if (v && s.std[i] == 0) || (!v && s.std[i] == 1) {
			good = false
			t.Error("Values at index", i, "differ")
			if v {
				t.Log("should be: true")
				t.Log("really is: false")
			} else {
				t.Log("should be: false")
				t.Log("really is: true")
			}
		}
	}

	return good
}


func (s *stringSort) Build(length int, isHash bool) {
	r := newRand()

	s.dev = make([]string, length)
	s.std = make([]string, length)

	for i := 0; i < length; i++ {
		l := 1
		for l > 0 {
			l = r.Intn(32)
		}
		b := make([]byte, l)
		for j := 0; j < l; j++ {
			// Fill each byte with a random printable character (0x21 - 0x7E).
			n := byte(r.Intn(93))
			b[j] = n + 33
		}
		s.dev[i] = string(b)
		s.std[i] = string(b)
	}
}

func (s *stringSort) Sort() error {
	return s.sort(s.dev)
}

func (s *stringSort) SortStd() {
	sort.Strings(s.std)
}

func (s *stringSort) Cmp(t *testing.T) bool {
	good := true
	for i, v := range s.dev {
		if v != s.std[i] {
			good = false
			t.Error("Values at index", i, "differ")
			t.Log("should be:", s.std[i])
			t.Log("really is:", v)
		}
	}

	return good
}



func newRand() *rand.Rand {
	seed   := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	return random
}
