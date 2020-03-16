// Package hsort provides a proof-of-concept for multiple sorting algorithms.
package hsort

import (
	"errors"
)


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

		hash := v % length
		table[hash] = append(table[hash], v)
	}

	// Iterate through our value range. If a value exists in the table, then we'll add it back to the list in now-sorted order.
	index := 0
	for i := low; i <= high; i++ {
		hash := i % length
		for _, v := range table[hash] {
			if v == i {
				list[index] = v
				index++
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
