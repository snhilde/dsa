// Package hsearch provides a proof-of-concept for multiple searching algorithms.
package hsearch

import (
	"errors"
)


// Find needle's position in haystack.
// Search progresses linearly.
func LinearInt(haystack []int, needle int) (int, error) {
	if len(haystack) < 1 {
		return -1, errors.New("Invalid haystack size")
	}

	for i, v := range haystack {
		if needle == v {
			return i, nil
		}
	}

	return -1, errors.New("Not found")
}

// Find needle's position in haystack (which must be sorted).
// Search progresses binarily.
func BinaryInt(haystack []int, needle int) (int, error) {
	// We're going to follow this sequence:
	// 1. Calculate start and end indices for the block where we want to search.
	// 2. Determine where the middle of this block is.
	//     2a. If the middle is the node we're looking for, return that index.
	//     2b. If it's greater than the value we're looking for, shift the block down so the end is right before it.
	//     2c. If it's less than the value we're looking for, shift the block up so the beginning is right after it.
	if len(haystack) < 1 {
		return -1, errors.New("Invalid haystack size")
	}

	start := 0
	end := len(haystack) - 1
	mid := (start + end) / 2
	value := haystack[mid]
	for start < end {
		if value > needle {
			end = mid - 1
		} else if value < needle {
			start = mid + 1
		}

		mid = (start + end) / 2
		value = haystack[mid]
		if value == needle {
			return mid, nil
		}
	}

	return -1, errors.New("Not found")
}
