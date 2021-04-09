package week02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	entry350input struct {
		nums1 []int
		nums2 []int
	}
	entry350 struct {
		name     string
		input    entry350input
		expected []int
	}
)

// run: go test -run Test_intersect
func Test_intersect(t *testing.T) {
	assert := assert.New(t)
	cases := []entry350{
		{
			name: "x1",
			input: entry350input{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			expected: []int{2, 2},
		},
		{
			name: "x2",
			input: entry350input{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			expected: []int{4, 9},
		},
		{
			name: "x3",
			input: entry350input{
				nums1: []int{9, 4, 9, 8, 4, 4},
				nums2: []int{1, 4, 4, 4, 1, 4},
			},
			expected: []int{4, 4, 4},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expected, intersect1(c.input.nums1, c.input.nums2), c.name)
		assert.Equal(c.expected, intersect2(c.input.nums1, c.input.nums2), c.name)
	}
}
