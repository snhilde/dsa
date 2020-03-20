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


func TestInsertion(t *testing.T) {
	i := intSort{sort: Insertion}
	testSort(t, &i, 100, 1000, false)

	u := uintSort{sort: Insertion}
	testSort(t, &u, 100, 1000, false)

	f := floatSort{sort: Insertion}
	testSort(t, &f, 100, 1000, false)

	b := boolSort{sort: Insertion}
	testSort(t, &b, 100, 1000, false)

	s := stringSort{sort: Insertion}
	testSort(t, &s, 100, 1000, false)
}

func TestSelection(t *testing.T) {
	i := intSort{sort: Selection}
	testSort(t, &i, 100, 1000, false)

	u := uintSort{sort: Selection}
	testSort(t, &u, 100, 1000, false)

	f := floatSort{sort: Selection}
	testSort(t, &f, 100, 1000, false)

	b := boolSort{sort: Selection}
	testSort(t, &b, 100, 1000, false)

	s := stringSort{sort: Selection}
	testSort(t, &s, 100, 1000, false)
}

func TestBubble(t *testing.T) {
	i := intSort{sort: Bubble}
	testSort(t, &i, 100, 1000, false)

	u := uintSort{sort: Bubble}
	testSort(t, &u, 100, 1000, false)

	f := floatSort{sort: Bubble}
	testSort(t, &f, 100, 1000, false)

	b := boolSort{sort: Bubble}
	testSort(t, &b, 100, 1000, false)

	s := stringSort{sort: Bubble}
	testSort(t, &s, 100, 1000, false)
}

func TestMergeInt(t *testing.T) {
}

func TestMergeIntOptimized(t *testing.T) {
}

func TestHashInt(t *testing.T) {
}


// Test out the various types/algorithms.
func testSort(t *testing.T, s sorter, n int, l int, isHash bool) {
	for i := 0; i < n; i++ {
		s.Build(l, isHash)

		if err := s.Sort(); err != nil {
			t.Error(err)
			return
		}

		s.SortStd()

		if !s.Cmp(t) {
			t.Error("-- Failed sort", i, "/", n, "--")
			return
		}
	}
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
