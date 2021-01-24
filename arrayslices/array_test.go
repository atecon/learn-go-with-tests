package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %d but want %d.", got, want)
		}
	})

	t.Run("collection of 4 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("Got %d but want %d.", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("array of two slices", func(t *testing.T) {
		// arr := ([]int{1,2}, []int{9,0}) -- list of slices?

		got := SumAll([]int{1, 2}, []int{9, 0})
		want := []int{3, 9}

		// reflect.DeepEqual() does a kind of element-wise comparison
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but want %v", got, want)
		}
	})
}

func TestSumAllWithAppend(t *testing.T) {
	t.Run("array of two slices", func(t *testing.T) {
		// arr := ([]int{1,2}, []int{9,0}) -- list of slices?

		got := SumAllWithAppend([]int{1, 2}, []int{9, 0})
		want := []int{3, 9}

		// reflect.DeepEqual() does a kind of element-wise comparison
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but want %v", got, want)
		}
	})
}

func TestSumTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but want %v.", got, want)
		}
	}

	t.Run("sum up some slices", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumTails([]int{}, []int{3, 4})
		want := []int{0, 4}

		checkSums(t, got, want)
	})

}
