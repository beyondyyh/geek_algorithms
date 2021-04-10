package week02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	entry347input struct {
		nums []int
		k    int
	}
	entry347 struct {
		name     string
		input    entry347input
		expected []int
	}
)

// run: go test -run Test_topKFrequent
func Test_topKFrequent(t *testing.T) {
	assert := assert.New(t)
	cases := []entry347{
		{
			name: "x1",
			input: entry347input{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			expected: []int{1, 2},
		},
		{
			name: "x2",
			input: entry347input{
				nums: []int{1},
				k:    1,
			},
			expected: []int{1},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expected, topKFrequent(c.input.nums, c.input.k), c.name)
	}
}
