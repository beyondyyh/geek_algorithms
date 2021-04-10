package week02

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	entry1 struct {
		name     string
		input    entry1input
		expected []int
	}
	entry1input struct {
		nums   []int
		target int
	}
)

// run: go test -run Test_twoSum
func Test_twoSum(t *testing.T) {
	assert := assert.New(t)
	cases := []entry1{
		{
			name: "x1",
			input: entry1input{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			name: "x2",
			input: entry1input{
				nums:   []int{3, 2, 4},
				target: 8,
			},
			expected: nil,
		},
		{
			name: "x3",
			input: entry1input{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := twoSum(c.input.nums, c.input.target)
			sort.Ints(out)
			assert.Equal(c.expected, out, c.name)
		})
	}
}
