// Package hsort provides a proof-of-concept for multiple sorting algorithms.
package hsort

import (
	"errors"
	"reflect"
	"fmt"
	"math"
)


// Sort the list using an insertion algorithm. The list must be a slice of a uniform data type.
func Insertion(list interface{}) error {
	// We're going to follow this sequence for each item in the list:
	// 1. Get the value at the current index.
	// 2. While the value is less than the value to the left of it, swap the two values.
	// 3. When the value is greater than the value to the left, it will also be greater than the value to the right and
	//    therefore in sorted order for this portion of the list.
	length, at, greater, swap, err := initSort(list)
	if err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			// Scan down the section of the list that is now sorted until we find the insertion point.
			curr := at(j)
			prev := at(j-1)
			if greater(prev, curr) {
				swap(j, j-1)
			} else {
				break
			}
		}
	}

	return nil
}

// Sort the list of ints using an insertion algorithm.
func InsertionInt(list []int) error {
	// We're going to follow this sequence for each item in the list:
	// 1. Get the value at the current index.
	// 2. For all previous items--starting from the current position and going down to the beginning--
	//    if the item at the index has a greater value, then shift it one to the right.
	// 3. Insert the value at the now-open index.
	if len(list) < 1 {
		return errors.New("Invalid list size")
	}

	for i, v := range list {
		for i > 0 {
			// Scan down the section of the list that is now sorted.
			previous := list[i-1]
			if previous > v {
				// Shift one to the right.
				list[i] = previous
				i--
			} else {
				break
			}
		}
		list[i] = v
	}

	return nil
}

// Sort the list using a selection algorithm. The list must be a slice of a uniform data type.
func Selection(list interface{}) error {
	// We're going to follow this sequence for each item in the list:
	// 1. Scan the entire list from the current position forward for the lowest value.
	// 2. Swap the current value and the lowest value.
	length, at, greater, swap, err := initSort(list)
	if err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		pos := i
		for j := i+1; j < length; j++ {
			// Check each value to see if it's lower than our current lowest.
			low := at(pos)
			try := at(j)
			if greater(low, try) {
				// We found a value lower than we currently have. Select it.
				pos = j
			}
		}
		// Swap the selected value with the current value.
		swap(i, pos)
	}

	return nil
}

// Sort the list of ints using a selection algorithm.
func SelectionInt(list []int) error {
	// We're going to follow this sequence for each item in the list:
	// 1. Scan the entire list from the current position forward for the lowest value.
	// 2. Swap the current value and the lowest value.
	length := len(list)
	if length < 1 {
		return errors.New("Invalid list size")
	}

	for i := range list {
		pos := i
		for j := i+1; j < length; j++ {
			// Check each value to see if it's lower than our current lowest.
			if list[j] < list[pos] {
				// We found a value lower than we currently have. Select it.
				pos = j
			}
		}
		// Swap the selected value with the current value.
		list[i], list[pos] = list[pos], list[i]
	}

	return nil
}

// Sort the list using a bubble algorithm. The list must be a slice of a uniform data type.
func Bubble(list interface{}) error {
	// For this function, we're going to iterate through every item in the list. If an item has a greater value than its
	// neighbor to the right, then we'll swap them. When we get to the end, we'll start again at the beginning and keep
	// doing this until we have one pass with no swaps.
	length, at, greater, swap, err := initSort(list)
	if err != nil {
		return err
	}

	// At the beginning of every pass, we'll set clean to true. If we perform any operation during the pass, we'll
	// toggle it false. When clean stays true the entire pass, everything is sorted.
	clean := false
	for !clean {
		clean = true
		for i := 0; i < length; i++ {
			if i + 1 == length {
				break
			}
			curr := at(i)
			next := at(i+1)
			if greater(curr, next) {
				swap(i, i+1)
				clean = false
			}
		}
	}

	return nil
}

// Sort the list of ints using a bubble algorithm.
func BubbleInt(list []int) error {
	// For this function, we're going to iterate through every item in the list. If an item has a greater value than its
	// neighbor to the right, then we'll swap them. When we get to the end, we'll start again at the beginning and keep
	// doing this until we have one pass with no swaps.
	length := len(list)
	if length < 1 {
		return errors.New("Invalid list size")
	}

	// At the beginning of every pass, we'll set clean to true. If we perform any operation during the pass, we'll
	// toggle it false. When clean stays true the entire pass, everything is sorted.
	clean := false
	for !clean {
		clean = true
		for i, v := range list {
			if i + 1 == length {
				break
			}
			if v > list[i+1] {
				list[i], list[i+1] = list[i+1], v
				clean = false
			}
		}
	}

	return nil
}

// Sort the list using a merging algorithm. The list must be a slice of a uniform data type.
func Merge(list interface{}) error {
	// For this sorting function, we're going to focus on a stack of blocks. A block is a subsection of the total list.
	// First, we're going to create a block for the entire list. Then we're going to follow this sequence for each sub-block:
	// - Look at the top block on the stack.
	//     - If it hasn't been split yet, then make two blocks out of each half and add them to the stack.
	//     - If it has already been split, then merge its two halves together and throw away the block.
	type block struct {
		index  int
		length int
		merge  bool
	}

	length, at, cmp, swap, err := initSort(list)
	if err != nil {
		return err
	}

	// Use these two lists to track where the item's will be sorted after we are done calculating.
	indexOf := make([]int, length) // item's order -> item's index
	orderOf := make([]int, length) // item's index -> item's order

	b := block{0, length, false}
	s := []block{b}
	for len(s) > 0 {
		// Pop the top block.
		b = s[len(s)-1]
		s = s[:(len(s)-1)]

		leftIndex := b.index
		leftLen := b.length / 2

		rightIndex := b.index + leftLen
		rightLen := b.length - leftLen
		if b.merge {
			// Calculate the sorted order of each item.
			for i := 0; i < b.length; i++ {
				if leftLen == 0 {
					// We only have values on the right side still.
					indexOf[i] = rightIndex
					orderOf[rightIndex] = i
					rightIndex++
				} else if rightLen == 0 {
					// We only have values on the left side still.
					indexOf[i] = leftIndex
					orderOf[leftIndex] = i
					leftIndex++
				} else if cmp(at(rightIndex), at(leftIndex)) {
					indexOf[i] = leftIndex
					orderOf[leftIndex] = i
					leftIndex++
					leftLen--
				} else {
					indexOf[i] = rightIndex
					orderOf[rightIndex] = i
					rightIndex++
					rightLen--
				}
			}
			// Now that everything is calculated, put the items into sorted order.
			for i := 0; i < b.length; i++ {
				// The item with order i is currently in this index:
				curr := indexOf[i]
				// However, we want it to be in this index:
				want := b.index + i
				// But the item with this order is currently occupying the desired index:
				order := orderOf[want]
				// First, let's swap the two items.
				swap(curr, want)
				// Now, let's update the second item's index, so we can find it later when it's turn has come to be
				// swapped into the correct order.
				indexOf[order] = curr
				orderOf[curr] = order
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

// Sort the list of ints using a merging algorithm.
func MergeInt(list []int) error {
	// For this sorting function, we're going to focus on a stack of blocks. A block is a subsection of the total list.
	// First, we're going to create a block for the entire list. Then we're going to follow this sequence for each sub-block:
	// - Look at the top block on the stack.
	//     - If it hasn't been split yet, then make two blocks out of each half and add them to the stack.
	//     - If it has already been split, then merge its two halves together and throw away the block.
	type block struct {
		index  int
		length int
		merge  bool
	}

	length := len(list)
	if length < 1 {
		return errors.New("Invalid list size")
	}

	// Create a space to hold our new list while we are merging stacks.
	tmp := make([]int, length)

	b := block{0, length, false}
	s := []block{b}
	for len(s) > 0 {
		// Pop the top block.
		b = s[len(s)-1]
		s = s[:(len(s)-1)]

		leftIndex := b.index
		leftLen := b.length / 2

		rightIndex := b.index + leftLen
		rightLen := b.length - leftLen
		if b.merge {
			// Merge the two halves.
			for i := 0; i < b.length; i++ {
				if leftLen == 0 {
					// We only have values on the right side still.
					tmp[i] = list[rightIndex]
					rightIndex++
				} else if rightLen == 0 {
					// We only have values on the left side still.
					tmp[i] = list[leftIndex]
					leftIndex++
				} else if list[leftIndex] < list[rightIndex] {
					tmp[i] = list[leftIndex]
					leftIndex++
					leftLen--
				} else {
					tmp[i] = list[rightIndex]
					rightIndex++
					rightLen--
				}
			}
			copy(list[b.index:], tmp[:b.length])
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

// Sort the list using a merging algorithm that is optimized for low memory use. The list must be a slice of a uniform data type.
func MergeOptimized(list interface{}) error {
	// While the standard merging algorithm first divides the list to be sorted into iteratively smaller blocks and then
	// merges back up the tree, this implementation starts at the bottom and merges upward immediately. This reduces
	// the memory overhead, as there is no tree allocation/construction.
	// We're going to focus on stacks and blocks here. Stacks are already-sorted sublists, and blocks are two stacks
	// that are being merged. The algorithm starts with a stack size of 1, meaning at the bottom level of individual
	// items. It will form blocks by merging two stacks together, working through the entire list. It will then make
	// stacks out of those blocks and continuing operating in this manner until the stack size consumes the entire list
	// and everything is sorted.
	length, at, cmp, swap, err := initSort(list)
	if err != nil {
		return err
	}

	// Use these two lists to track where the item's will be sorted after we are done calculating.
	indexOf := make([]int, length) // item's order -> item's index
	orderOf := make([]int, length) // item's index -> item's order

	// Progressively work from smallest stack size up.
	for stackSize := 1; stackSize < length; stackSize *= 2 {
		// A block represents both stacks put together.
		blockSize := stackSize * 2
		numBlocks := (length / blockSize) + 1

		// Operate on each individual block.
		for i := 0; i < numBlocks; i++ {
			index := blockSize * i
			// If this is the last block in the row, we have to compensate for potentially not having a full block.
			if i == numBlocks - 1 {
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
			for j := 0; j < blockSize; j++ {
				if leftLen == 0 {
					// We only have values on the right side still.
					indexOf[j] = rightIndex
					orderOf[rightIndex] = j
					rightIndex++
				} else if rightLen == 0 {
					// We only have values on the left side still.
					indexOf[j] = leftIndex
					orderOf[leftIndex] = j
					leftIndex++
				} else if cmp(at(rightIndex), at(leftIndex)) {
					indexOf[j] = leftIndex
					orderOf[leftIndex] = j
					leftIndex++
					leftLen--
				} else {
					indexOf[j] = rightIndex
					orderOf[rightIndex] = j
					rightIndex++
					rightLen--
				}
			}
			// Now that everything is calculated, put the items into sorted order.
			for j := 0; j < blockSize; j++ {
				// The item with order i is currently in this index:
				curr := indexOf[j]
				// However, we want it to be in this index:
				want := index + j
				// But the item with this order is currently occupying the desired index:
				order := orderOf[want]
				// First, let's swap the two items.
				swap(curr, want)
				// Now, let's update the second item's index, so we can find it later when it's turn has come to be
				// swapped into the correct order.
				indexOf[order] = curr
				orderOf[curr] = order
			}
		}
	}

	return nil
}

// Sort the list of ints using a merging algorithm that is optimized for low memory use.
func MergeIntOptimized(list []int) error {
	// While the standard merging algorithm first divides the list to be sorted into iteratively smaller blocks and then
	// merges back up the tree, this implementation starts at the bottom and merges upward immediately. This reduces
	// the memory overhead, as there is no tree allocation/construction.
	// We're going to focus on stacks and blocks here. Stacks are already-sorted sublists, and blocks are two stacks
	// that are being merged. The algorithm starts with a stack size of 1, meaning at the bottom level of individual
	// items. It will form blocks by merging two stacks together, working through the entire list. It will then make
	// stacks out of those blocks and continuing operating in this manner until the stack size consumes the entire list
	// and everything is sorted.
	length := len(list)
	if length < 1 {
		return errors.New("Invalid list size")
	}

	// Create a space to hold our new list while we are merging stacks.
	tmp := make([]int, length)

	// Progressively work from smallest stack size up.
	for stackSize := 1; stackSize < length; stackSize *= 2 {
		// A block represents both stacks put together.
		blockSize := stackSize * 2
		numBlocks := (length / blockSize) + 1

		// Operate on each individual block.
		for i := 0; i < numBlocks; i++ {
			index := blockSize * i
			// If this is the last block in the row, we have to compensate for potentially not having a full block.
			if i == numBlocks - 1 {
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
			for j := 0; j < blockSize; j++ {
				if leftLen == 0 {
					// We only have values on the right side still.
					copy(tmp[j:], list[rightIndex:rightIndex+rightLen])
					break
				} else if rightLen == 0 {
					// We only have values on the left side still.
					copy(tmp[j:], list[leftIndex:leftIndex+leftLen])
					break
				} else if list[leftIndex] < list[rightIndex] {
					tmp[j] = list[leftIndex]
					leftIndex++
					leftLen--
				} else {
					tmp[j] = list[rightIndex]
					rightIndex++
					rightLen--
				}
			}
			copy(list[index:], tmp[:blockSize])
		}
	}

	return nil
}

// Sort the list of ints using a hashing algorithm.
// Note: The efficiency of this algorithm decreases as the range of possible values in the list increases.
func HashInt(list []int) error {
	// We're going to follow this sequence:
	// 1. Build a hash table and populate it with every item in the list. Because we do not have any prior knowledge of
	//    value range, our hash function is a simple value mod length. This gives distribution in the array equal to the
	//    value distribution in the list. We're going to handle collisions with chaining.
	// 2. As we are populating the table, we are also going to find the lowest and highest values.
	// 3. Iterate through every value from the lowest to the highest. If the value exists in the table, put it in the
	//    list at the current index and increment the index.
	// Note: Due to the low-to-high value iteration and table lookup, this algorithm is only efficient for low value
	// ranges. The time complexity is linear for input size AND linear for value range.
	length := len(list)
	if length < 1 {
		return errors.New("Invalid list size")
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

	// Iterate through our value range. If a value exists in the table, then we'll add it back to the list in now-sorted order.
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


// Helper function that will set up all the variables and functions necessary for determining the list's underlying type
// and acting on that type appropriately. It will return these values:
// 1. The length of the list
// 2. A function that will get the Value at the given index
// 3. A function that will compare the two Values, return M_TRUE if the first is greater and M_FALSE if the second is greater.
// 4. A function that will swap the two Values at the given indices.
// 5. Any error that occurred along the way, or nil if no error occurred
func initSort(list interface{}) (int, func(int) reflect.Value, func(i, j reflect.Value) bool, func(i, j int), error) {
	// Pull out the underlying Value, and make sure it's a slice.
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Slice {
		return -1, nil, nil, nil, errors.New("List must be slice")
	}

	// Find out how long our list is.
	l := v.Len()
	if l < 1 {
		return -1, nil, nil, nil, errors.New("Invalid list size")
	}

	// Make sure we have a type that we can work with.
	var k reflect.Kind
	switch v := v.Index(0).Kind(); v {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		k = reflect.Int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		k = reflect.Uint64
	case reflect.Float32, reflect.Float64:
		k = reflect.Float64
	case reflect.Bool:
		k = reflect.Bool
	case reflect.String:
		k = reflect.String
	default:
		return -1, nil, nil, nil, errors.New(fmt.Sprintf("Invalid underlying type (%s)", v))
	}

	// Construct the function that will return the Value at the given index.
	at := func(i int) reflect.Value {
		return v.Index(i)
	}

	// Construct the function that will compare the two Values.
	// Returns true if i is greater than j and items need to be swapped.
	greater := func(i, j reflect.Value) bool {
		switch k {
		case reflect.Int64:
			if i.Int() > j.Int() {
				return true
			}
		case reflect.Uint64:
			if i.Uint() > j.Uint() {
				return true
			}
		case reflect.Float64:
			if i.Float() > j.Float() {
				return true
			}
		case reflect.Bool:
			if i.Bool() && !j.Bool() {
				return true
			}
		case reflect.String:
			if i.String() > j.String() {
				return true
			}
		}

		return false
	}

	// Our swapping function is straight from the reflect library (thanks).
	swap := reflect.Swapper(list)

	return l, at, greater, swap, nil
}
