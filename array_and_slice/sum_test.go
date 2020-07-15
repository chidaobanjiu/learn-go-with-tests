package array_and_slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array sum", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := ArraySum(numbers)
		want := 15

		if want != got {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	})

	t.Run("slice sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := SliceSum(numbers)
		want := 15

		if want != got {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 3, 4})
	want := []int{3, 9}

	if !reflect.DeepEqual(want, got) { // relect.DeepEqual 不是类型安全的，编译器无法发现其类型错误
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(want, got) { // relect.DeepEqual 不是类型安全的，编译器无法发现其类型错误
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	t.Run("normal test", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 3, 4})
		want := []int{2, 7}
		checkSum(t, got, want)
	})

	t.Run("zero slice test", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 4})
		want := []int{0, 7}
		checkSum(t, got, want)
	})
}
