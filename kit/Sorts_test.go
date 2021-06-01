package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type arraySortFunc func([]int, int, int) []int

// run: go test -v -run Test_quickSort
func Test_quickSort(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "quickSort-1",
			nums:   []int{2, 1, 5, 3, 4, 6, 7, 0},
			expect: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		nums := c.nums
		left, right := 0, len(nums)-1
		assert.Equal(c.expect, quickSort(nums, left, right), c.name)
	}
}

// run: go test -v -run Test_mergeSort
func Test_mergeSort(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "mergeSort-1",
			nums:   []int{2, 1, 5, 3, 4, 6, 7, 0},
			expect: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, mergeSort(c.nums), c.name)
	}
}

// run: go test -v -run Test_heapSort
func Test_heapSort(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "heapSort-1",
			nums:   []int{2, 1, 5, 3, 4, 6, 7, 0},
			expect: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, heapSort(c.nums), c.name)
	}
}
