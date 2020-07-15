package main

import "fmt"

// MergeSort is the main function for merge sort algorithm
func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2

	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merge performs merging on the recursively divided merge-sort arrays
func Merge(left, right []int) []int {
	length, i, j := len(left)+len(right), 0, 0

	slice := make([]int, length)

	for k := 0; k < length; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if i <= len(left)-1 && j > len(right)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}

	return slice
}

func main() {
	slice1 := []int{5, 4, 1, 8, 7, 2, 6, 3}
	fmt.Printf("%v\n", MergeSort(slice1))
	slice2 := []int{9, 5, 4, 1, 8, 7, 2, 6, 3}
	fmt.Printf("%v\n", MergeSort(slice2))
}
