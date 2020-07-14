package problem40

import (
	"testing"
)

func TestUniqueChecker(t *testing.T) {

	t.Run("Standard testing", func(t *testing.T) {
		t.Helper()
		arr := []int{6, 1, 1, 1, 3, 6, 6}
		got := UniqueChecker(arr)
		want := 3
		assertError(t, got, want)
	})

	t.Run("With 0 and negative numbers", func(t *testing.T) {
		t.Helper()
		arr := []int{0, -1, -1, -1}
		got := UniqueChecker(arr)
		want := 0
		assertError(t, got, want)
	})
}

func assertError(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
