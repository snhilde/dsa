// Package hsort provides a proof-of-concept for multiple sorting algorithms.
package hsort

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"

	"github.com/snhilde/dsa/data_structures/htree"
)

var (
	// errBadLength is the error message for an invalid list size.
	errBadLength = fmt.Errorf("invalid list size")

	// errBadCmp is the error message for an invalid comparison function.
	errBadCmp = fmt.Errorf("invalid comparison function")
)

// Insertion sorts the list using an insertion algorithm. The list must be a slice. The comparison
// function less should return true only if the item at index i is less than the item at index j.
func Insertion(list interface{}, less func(i, j int) bool) error {
	// We're going to follow this sequence for each item in the list:
	// 1. Get the value at the current index.
	// 2. While the value is less than the value to the left of it, swap the two values.
	// 3. When the value is greater than the value to the left, it will also be greater than the
	//    value to the right and therefore in sorted order for this portion of the list.
	length, swapper, err := initSort(list, less)
	if err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			// Scan down the section of the list that is now sorted until we find the insertion point.
			if less(j, j-1) {
				swapper(j, j-1)
			} else {
				break
			}
		}
	}

	return nil
}

// Selection sorts the list using a selection algorithm. The list must be a slice. The comparison
// function less should return true only if the item at index i is less than the item at index j.
func Selection(list interface{}, less func(int, int) bool) error {
	// We're going to follow this sequence for each item in the list:
	// 1. Scan the entire list from the current position forward for the lowest value.
	// 2. Swap the current value and the lowest value.
	length, swapper, err := initSort(list, less)
	if err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		pos := i
		for j := i + 1; j < length; j++ {
			// Check each value to see if it's lower than our current lowest.
			if less(j, pos) {
				// We found a value lower than we currently have. Select it.
				pos = j
			}
		}
		// Swap the selected value with the current value.
		swapper(i, pos)
	}

	return nil
}

// Bubble sorts the list using a bubble algorithm. The list must be a slice. The comparison function
// less should return true only if the item at index i is less than the item at index j.
func Bubble(list interface{}, less func(int, int) bool) error {
	// For this function, we're going to iterate through every item in the list. If an item has a
	// greater value than its neighbor to the right, then we'll swap them. When we get to the end,
	// we'll start again at the beginning and keep doing this until we have one pass with no swaps.
	length, swapper, err := initSort(list, less)
	if err != nil {
		return err
	}

	// At the beginning of every pass, we'll set clean to true. If we perform any operation during
	// the pass, we'll toggle it false. When clean stays true the entire pass, everything is sorted.
	clean := false
	for !clean {
		clean = true
		for i := 0; i < length-1; i++ {
			if less(i+1, i) {
				swapper(i, i+1)
				clean = false
			}
		}
	}

	return nil
}

// Merge sorts the list using a merging algorithm. The list must be a slice. The comparison function
// less should return true only if the item at index i is less than the item at index j.
func Merge(list interface{}, less func(int, int) bool) error {
	// For this sorting function, we're going to focus on a stack of blocks. A block is a subsection
	// of the total list. First, we're going to create a block for the entire list. Then we're going
	// to follow this sequence for each sub-block:
	// - Look at the top block on the stack.
	//     - If it hasn't been split yet, then make two blocks out of each half and add them to the stack.
	//     - If it has already been split, then merge its two halves together and throw away the block.
	// Typically, on the merging step, you would use a temporary array to handle sorting the stacks
	// together. Unfortunately, we will not know the necessary underlying type beforehand, so we
	// can't create this temporary array. Instead of moving items to another list and then moving
	// them back in sorted order, we're going to keep track of where each item should be. After we
	// calculate each item's required index for sorting, we'll start swapping each item into its
	// correct position.
	type block struct {
		index  int
		length int
		merge  bool
	}

	length, swapper, err := initSort(list, less)
	if err != nil {
		return err
	}

	// This list will track where the item at each index needs to be moved to in order to sort the
	// list properly.
	moveTo := make([]int, length)

	b := block{0, length, false}
	s := []block{b}
	for len(s) > 0 {
		// Pop the top block.
		b = s[len(s)-1]
		s = s[:(len(s) - 1)]

		leftIndex := b.index
		leftLen := b.length / 2

		rightIndex := b.index + leftLen
		rightLen := b.length - leftLen

		if b.merge {
			// Calculate the sorted order of each item.
			for i := b.index; i < b.index+b.length; i++ {
				if leftLen == 0 {
					// We only have values on the right side still.
					moveTo[rightIndex] = i
					rightIndex++
				} else if rightLen == 0 {
					// We only have values on the left side still.
					moveTo[leftIndex] = i
					leftIndex++
				} else if less(leftIndex, rightIndex) {
					moveTo[leftIndex] = i
					leftIndex++
					leftLen--
				} else {
					moveTo[rightIndex] = i
					rightIndex++
					rightLen--
				}
			}

			// Now that everything is calculated, put the items into sorted order.
			for i := b.index; i < b.index+b.length; i++ {
				// Keep swapping items until the one at this index should be in this index.
				for i != moveTo[i] {
					// The item at this index needs to be moved to this index:
					to := moveTo[i]
					// Swap the two items.
					swapper(i, to)
					// Update the sorted index of each item.
					moveTo[i] = moveTo[to]
					moveTo[to] = to
				}
			}
		} else {
			// We're still on the splitting phase.
			b.merge = true
			s = append(s, b)
			if leftLen > 1 {
				// Add left-side block to stack.
				s = append(s, block{leftIndex, leftLen, false})
			}
			if rightLen > 1 {
				// Add right-side block to stack.
				s = append(s, block{rightIndex, rightLen, false})
			}
		}
	}

	return nil
}

// MergeOptimized sorts the list using a merging algorithm that is optimized for low memory use. The
// list must be a slice. The comparison function less should return true only if the item at index i
// is less than the item at index j.
func MergeOptimized(list interface{}, less func(int, int) bool) error {
	// While the standard merging algorithm first divides the list to be sorted into iteratively
	// smaller blocks and then merges back up the tree, this implementation starts at the bottom and
	// merges upward immediately. This reduces the memory overhead, as there is no tree
	// allocation/construction.
	// We're going to focus on stacks and blocks here. Stacks are already-sorted sublists, and
	// blocks are two stacks that are being merged. The algorithm starts with a stack size of 1,
	// meaning at the bottom level of individual items. It will form blocks by merging two stacks
	// together, working through the entire list. It will then make stacks out of those blocks and
	// continuing operating in this manner until the stack size consumes the entire list and
	// everything is sorted.
	// Typically, you would use a temporary array to handle sorting the stacks together.
	// Unfortunately, we will not know the necessary underlying type beforehand, so we can't create
	// this temporary array. Instead of moving items to another list and then moving them back in
	// sorted order, we're going to keep track of where each item should be. After we calculate each
	// item's required index for sorting, we'll start swapping each item into its correct position.
	length, swapper, err := initSort(list, less)
	if err != nil {
		return err
	}

	// This list will track where the item at each index needs to be moved to in order to sort the
	// list properly.
	moveTo := make([]int, length)

	// Progressively work from smallest stack size up.
	for stackSize := 1; stackSize < length; stackSize *= 2 {
		// A block represents both stacks put together.
		blockSize := stackSize * 2
		numBlocks := (length / blockSize) + 1

		// Operate on each individual block.
		for i := 0; i < numBlocks; i++ {
			index := blockSize * i
			// If this is the last block in the row, we have to compensate for potentially not
			// having a full block.
			if i == numBlocks-1 {
				blockSize = length - index
				if blockSize <= stackSize {
					// Already sorted
					break
				}
			}

			leftIndex := index
			leftLen := stackSize

			rightIndex := index + stackSize
			rightLen := blockSize - stackSize

			// Merge both stacks together.
			for j := index; j < index+blockSize; j++ {
				if leftLen == 0 {
					// We only have values on the right side still.
					moveTo[rightIndex] = j
					rightIndex++
				} else if rightLen == 0 {
					// We only have values on the left side still.
					moveTo[leftIndex] = j
					leftIndex++
				} else if less(leftIndex, rightIndex) {
					moveTo[leftIndex] = j
					leftIndex++
					leftLen--
				} else {
					moveTo[rightIndex] = j
					rightIndex++
					rightLen--
				}
			}

			// Now that everything is calculated, put the items into sorted order.
			for j := index; j < index+blockSize; j++ {
				// Keep swapping items until the one at this index should be in this index.
				for j != moveTo[j] {
					// The item at this index needs to be moved to this index:
					to := moveTo[j]
					// Swap the two items.
					swapper(j, to)
					// Update the sorted index of each item.
					moveTo[j] = moveTo[to]
					moveTo[to] = to
				}
			}
		}
	}

	return nil
}

// Hash sorts the list of ints using a hashing algorithm.
// Note: The efficiency of this algorithm decreases as the range of possible values in the list increases.
func Hash(list []int) error {
	// We're going to follow this sequence:
	// 1. Build a hash table and populate it with every item in the list. Because we do not have any
	//    prior knowledge of value range, our hash function is a simple value mod length. This gives
	//    distribution in the array equal to the value distribution in the list. We're going to
	//    handle collisions with chaining.
	// 2. As we are populating the table, we are also going to find the lowest and highest values.
	// 3. Iterate through every value from the lowest to the highest. If the value exists in the
	//    table, put it in the list at the current index and increment the index.
	// Note: Due to the low-to-high value iteration and table lookup, this algorithm is only
	// efficient for low value ranges. The time complexity is linear for input size AND linear for
	// value range.
	length := len(list)
	if length < 1 {
		return errBadLength
	}

	// Give the table a 75% fill to decrease the number of collisions and subsequent append operations.
	length = int(float64(length) * 1.33)

	// Build out our hash table.
	low := list[0]
	high := list[0]
	table := make([][]int, length)
	for _, v := range list {
		if v < low {
			low = v
		} else if v > high {
			high = v
		}

		hash := int(math.Abs(float64(v % length)))
		table[hash] = append(table[hash], v)
	}

	// Iterate through our value range. If a value exists in the table, then we'll add it back to
	// the list in now-sorted order.
	index := 0
	for i := low; i <= high; i++ {
		hash := int(math.Abs(float64(i % length)))
		for _, v := range table[hash] {
			if v == i {
				list[index] = v
				index++
			}
		}
	}

	return nil
}

// Bogo sorts the list using a bogo algorithm. The list must be a slice. The comparison function
// less should return true only if the item at index i is less than the item at index j.
// Note: This is a bogus algorithm intended to be highly inefficient.
func Bogo(list interface{}, less func(int, int) bool) error {
	// The loop for this process is simple:
	// 1. Make sure the list is not currently sorted.
	// 2. Randomize the items in the list.
	// We will continue doing this until the loop is sorted.
	length, swapper, err := initSort(list, less)
	if err != nil {
		return err
	}

	sorted := false
	for !sorted {
		// Check to see if the list is sorted.
		sorted = true
		for i := 0; i < length-1; i++ {
			if less(i+1, i) {
				sorted = false
				break
			}
		}

		if !sorted {
			// The list is not sorted yet. Randomly shuffle it.
			rand.Shuffle(length, swapper)
		}
	}

	return nil
}

// Binary sorts the list of ints using a binary search tree.
func Binary(list []int) error {
	length := len(list)
	if length < 1 {
		return errBadLength
	}

	tree := htree.New()
	for _, v := range list {
		if err := tree.Add(v, v); err != nil {
			return err
		}
	}

	sorted := tree.DFS()
	for i, v := range sorted {
		list[i] = v.(int)
	}

	return nil
}

// Helper function that will set up all the variables and functions necessary for determining the
// list's underlying type and acting on that type appropriately. It will return these values:
// 1. The length of the list
// 2. A function that will swap the two Values at the given indices.
// 3. Any error that occurred along the way, or nil if no error occurred.
func initSort(list interface{}, cmp func(int, int) bool) (int, func(i, j int), error) {
	// Pull out the underlying Value, and make sure it's a slice.
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Slice {
		return -1, nil, fmt.Errorf("list must be slice")
	}

	// Make sure that we have a comparison function.
	if cmp == nil {
		return -1, nil, errBadCmp
	}

	// Find out how long our list is.
	length := v.Len()
	if length < 1 {
		return 0, nil, errBadLength
	}

	// Our swapping function is straight from the reflect library (thank you to the authors).
	swapper := reflect.Swapper(list)

	return length, swapper, nil
}
