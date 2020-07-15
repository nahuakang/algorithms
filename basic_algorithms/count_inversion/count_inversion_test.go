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
	slice1 := []int{0, 0, 0, 0, 0, 0, 0}
	_, count1 := CountInversion(slice1)
	answer1 := 0
	if count1 != answer1 {
		t.Errorf("CountInversion incorrect, got: %d, want: %d\n", count1, answer1)
	}

	slice2 := []int{1, 3, 5, 2, 4, 6}
	_, count2 := CountInversion(slice2)
	answer2 := 3
	if count2 != answer2 {
		t.Errorf("Merge Sort incorrect, got: %d, want: %d\n", count2, answer2)
	}

	slice3 := []int{1, 3, 5, 7, 2, 4, 6}
	_, count3 := CountInversion(slice3)
	answer3 := 6
	if count2 != answer2 {
		t.Errorf("Merge Sort incorrect, got: %d, want: %d\n", count3, answer3)
	}

	slice4 := []int{5, 3, 1, 6, 4, 2}
	_, count4 := CountInversion(slice4)
	answer4 := 9
	if count4 != answer4 {
		t.Errorf("Merge Sort incorrect, got: %d, want: %d\n", count4, answer4)
	}

	slice5 := []int{6, 5, 4, 3, 2, 1}
	_, count5 := CountInversion(slice5)
	answer5 := 15
	if count5 != answer5 {
		t.Errorf("Merge Sort incorrect, got: %d, want: %d\n", count5, answer5)
	}
}
