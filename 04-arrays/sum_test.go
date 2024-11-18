package _4_arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	checkSum := func(t *testing.T, got, want int, numbers []int) {
		if got != want {
			t.Errorf("expected '%d', but got '%d', given %v", want, got, numbers)
		}
	}

	t.Run("calculate the sum of an array: 1, 2, 3, 4, 5", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		sum := Sum(array)
		expected := 15
		checkSum(t, sum, expected, array)
	})
	t.Run("calculate the sum of an array: 1, 2, 3, 4", func(t *testing.T) {
		array := []int{1, 2, 3, 4}
		sum := Sum(array)
		expected := 10
		checkSum(t, sum, expected, array)
	})
}

func TestSumAll(t *testing.T) {

	checkSum := func(t *testing.T, got, want, numbers1, numbers2 []int) {
		if !slices.Equal(got, want) {
			t.Errorf("expected %v, but got %v, given array1: %v and array2: %v", got, want, numbers1, numbers2)
		}
	}

	t.Run("calculate sum of arrays {1,2,3} and {5, 7}", func(t *testing.T) {
		array1 := []int{1, 2, 3}
		array2 := []int{5, 7}
		sum := SumAll(array1, array2)
		expected := []int{6, 12}
		checkSum(t, sum, expected, array1, array2)
	})
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want, numbers1, numbers2 []int) {
		if !slices.Equal(got, want) {
			t.Errorf("expected %v, but got %v, given array1: %v and array2: %v", got, want, numbers1, numbers2)
		}
	}

	t.Run("sum of tails of arrays {1, 2, 3} and {5, 7, 9}", func(t *testing.T) {
		array1 := []int{1, 2, 3}
		array2 := []int{5, 7, 9}
		sum := SumAllTails(array1, array2)
		expected := []int{5, 16}
		checkSum(t, sum, expected, array1, array2)
	})
	t.Run("sum of tails of arrays {1, 2, 3} and {}", func(t *testing.T) {
		array1 := []int{1, 2, 3}
		array2 := []int{}
		sum := SumAllTails(array1, array2)
		expected := []int{5, 0}
		checkSum(t, sum, expected, array1, array2)
	})
}

func BenchmarkSum(b *testing.B) {
	var numbers = make([]int, b.N+1)
	for i := 0; i < b.N; i++ {
		numbers[i] = i + 1
		Sum(numbers[:i])
	}
}
