package binarySearch

import "errors"

// ErrNotFound for no element in the sorted array
var ErrNotFound = errors.New("could not find the element in the array")

// BinarySearch searches for an element in a sorted array and returns its index
func BinarySearch(arr []int, item int) (int, error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == item {
			return mid, nil
		} else if arr[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1, ErrNotFound
}
