package week02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	entry59input struct {
		nums []int
		k    int
	}
	entry59 struct {
		name     string
		input    entry59input
		expected []int
	}
)

// run: go test -run Test_maxSlidingWindow
func Test_maxSlidingWindow(t *testing.T) {
	assert := assert.New(t)
	cases := []entry59{
		{
			name: "x1",
			input: entry59input{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			expected: []int{3, 3, 5, 5, 6, 7},
		},
	}
	for _, c := range cases {
		// assert.Equal(c.expected, maxSlidingWindow1(c.input.nums, c.input.k), "maxSlidingWindow1->"+c.name)
		assert.Equal(c.expected, maxSlidingWindow2(c.input.nums, c.input.k), "maxSlidingWindow2->"+c.name)
	}
}
