package sorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ikbarfp/go-dsa/sorting"
)

func TestMergeSort_GivenValidData_ShouldBeSorted(t *testing.T) {
	array := []int{5, 4, 3, 2, 1}
	expArray := []int{1, 2, 3, 4, 5}
	result := sorting.MergeSort(array)
	assert.Equal(t, expArray, result)
}

func TestMergeSort_GivenNil_ShouldReturnNil(t *testing.T) {
	result := sorting.MergeSort(nil)
	assert.Nil(t, result)
}

func BenchmarkMergeSort(b *testing.B) {
	array := []int{5, 4, 3, 2, 1, 1, 23, 12, 69, 345, 345, 456, 96, 7, 1, 2, 321, 35, 5, 98, 6, 23, 42, 36, 7, 1}
	for i := 0; i < b.N; i++ {
		sorting.MergeSort(array)
	}
}
