package problem2

import (
	"reflect"
	"testing"
)

func TestMultiplier(t *testing.T) {

	t.Run("Standard 5 element array", func(t *testing.T) {
		t.Helper()
		arr := []int{1, 2, 3, 4, 5}
		got := Multiplier(arr)
		want := []int{120, 60, 40, 30, 24}
		assertError(t, got, want)
	})

	t.Run("Standard 3 element array", func(t *testing.T) {
		t.Helper()
		arr := []int{3, 2, 1}
		got := Multiplier(arr)
		want := []int{2, 3, 6}
		assertError(t, got, want)
	})
}

func assertError(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
