/*
Given an array of integers where every integer occurs three times
except for one integer, which only occurs once, find and return
the non-duplicated integer.

For example,
given [6, 1, 3, 3, 3, 6, 6], return 1,
given [13, 19, 13, 13], return 19.

Do this in O(N) time and O(1) space.
NOTE: I choose to use O(log N) space instead of bitwise operators.
*/

package problem40

// UniqueChecker returns the only element that's not duplicated
func UniqueChecker(arr []int) int {
	results := make(map[int]int)
	for _, valKey := range arr {
		results[valKey]++
	}

	for _, valKey := range results {
		if results[valKey] == 1 {
			return valKey
		}
	}

	return 0
}
