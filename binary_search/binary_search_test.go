package binarySearch

import "testing"

var testArray = []int{1, 2, 3, 4, 5, 6}

func TestBinarySearch(t *testing.T) {
	t.Run("Should return an element", func(t *testing.T) {
		got, _ := BinarySearch(testArray, 6)
		want := 5

		assertIndex(t, got, want)
	})

	t.Run("Should return ErrNotFound", func(t *testing.T) {
		_, got := BinarySearch(testArray, -1)
		want := ErrNotFound

		assertError(t, got, want)
	})
}

func assertIndex(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
