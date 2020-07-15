package main

import "fmt"

// CountInversion counts all the inverted pairs of an array
func CountInversion(slice []int) ([]int, int) {
	if len(slice) < 2 {
		return slice, 0
	}

	mid := len(slice) / 2
	left := slice[:mid]
	right := slice[mid:]

	sortedLeft, leftInvCount := CountInversion(left)
	sortedRight, rightInvCount := CountInversion(right)

	return MergeAndCount(sortedLeft, sortedRight, leftInvCount+rightInvCount)
}

// MergeAndCount merges the divided arrays and counts split inversions
func MergeAndCount(left, right []int, prevCount int) ([]int, int) {
	slice := make([]int, len(left)+len(right))
	i, j, splitInvCount := 0, 0, 0

	for k := range slice {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			slice[k] = left[i]
			i++
		} else if left[i] > right[j] {
			slice[k] = right[j]
			splitInvCount += len(left) - i
			j++
		}
	}
	return slice, splitInvCount + prevCount
}

func main() {
	slice1 := []int{1, 3, 5, 2, 4, 6}
	sorted1, invCount1 := CountInversion(slice1)
	fmt.Printf("%v, count: %d\n", sorted1, invCount1)

	slice2 := []int{5, 3, 1, 6, 4, 2}
	sorted2, invCount2 := CountInversion(slice2)
	fmt.Printf("%v, count: %d\n", sorted2, invCount2)
}
