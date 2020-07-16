package selectionSort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Run("Already sorted", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6}
		got := SelectionSort(arr)
		want := []int{1, 2, 3, 4, 5, 6}
		assertEqual(t, got, want)
	})

	t.Run("Completely reversed", func(t *testing.T) {
		arr := []int{6, 5, 4, 3, 2, 1}
		got := SelectionSort(arr)
		want := []int{1, 2, 3, 4, 5, 6}
		assertEqual(t, got, want)
	})

	t.Run("With negatives", func(t *testing.T) {
		arr := []int{-3, 3, -2, 2, -1, 1, 0}
		got := SelectionSort(arr)
		want := []int{-3, -2, -1, 0, 1, 2, 3}
		assertEqual(t, got, want)
	})
}

func assertEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q want %q", got, want)
	}
}
