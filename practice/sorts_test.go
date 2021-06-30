package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_mergeSort
func Test_mergeSort(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "x1",
			nums:   []int{4, 1, 3, 2, 5, 0},
			expect: []int{0, 1, 2, 3, 4, 5},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, mergeSort(c.nums), c.name)
	}
}

// run: go test -v -run Test_quickSort
func Test_quickSort(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "x1",
			nums:   []int{4, 1, 3, 2, 5, 0},
			expect: []int{0, 1, 2, 3, 4, 5},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, quickSort(c.nums, 0, len(c.nums)-1), c.name)
	}
}
