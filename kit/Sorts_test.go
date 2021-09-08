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
		// fmt.Println(c.nums)
		nums := c.nums
		start, end := 0, len(nums)-1
		assert.Equal(c.expect, quickSort(nums, start, end), c.name)
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
			name:   "mergeSort2-1",
			nums:   []int{2, 1, 5, 3, 4, 6, 7, 0},
			expect: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, mergeSort(c.nums), c.name)
	}
}

// run: go test -v -run Test_mergeSort2
func Test_mergeSort2(t *testing.T) {
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
		// fmt.Println(c.nums)
		nums := c.nums
		left, right := 0, len(nums)-1
		mergeSort2(nums, left, right)
		assert.Equal(c.expect, nums, c.name)
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
