package main

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMergeSort(t *testing.T) {
	slice1 := []int{5, 4, 1, 8, 7, 2, 6, 3}
	answer1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !equal(MergeSort(slice1), answer1) {
		t.Errorf("Merge Sort incorrect, got: %v, want: %v\n", slice1, answer1)
	}

	slice2 := []int{9, 5, 4, 1, 8, 7, 2, 6, 3}
	answer2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !equal(MergeSort(slice2), answer2) {
		t.Errorf("Merge Sort incorrect, got: %v, want: %v\n", slice2, answer2)
	}
}
