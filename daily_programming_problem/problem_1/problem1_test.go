package problem1

import "testing"

func TestListChecker(t *testing.T) {
	t.Run("Standard test", func(t *testing.T) {
		list := []int{10, 15, 3, 7}
		got := ListChecker(list, 17)
		want := true
		assertError(t, got, want)
	})

	t.Run("With negative value", func(t *testing.T) {
		list := []int{2, 0, -1, 3}
		got := ListChecker(list, 1)
		want := true
		assertError(t, got, want)
	})

	t.Run("Single element", func(t *testing.T) {
		list := []int{1}
		got := ListChecker(list, 1)
		want := true
		assertError(t, got, want)
	})

	t.Run("Should not return true", func(t *testing.T) {
		list := []int{3, 10, 4, 6}
		got := ListChecker(list, 1)
		want := false
		assertError(t, got, want)
	})
}

func assertError(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
