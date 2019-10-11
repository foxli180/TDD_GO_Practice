package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T)  {

	assetSumEquals := func(t *testing.T, got, want int, numbers []int) {
		if want != got{
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}
		got := Sum(numbers);
		want := 15
		assetSumEquals(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assetSumEquals(t, got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	CompareArrayEquals(got, want, t)

}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		CompareArrayEquals(got, want, t)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		CompareArrayEquals(got, want, t)
	})
}

func CompareArrayEquals(got []int, want []int, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)

	}
}