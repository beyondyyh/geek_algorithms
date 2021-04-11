package week02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entry84 struct {
	name     string
	input    []int
	expected int
}

// run: go test -run Test_largestRectangleArea
func Test_largestRectangleArea(t *testing.T) {
	assert := assert.New(t)
	cases := []entry84{
		{
			name:     "x1",
			input:    []int{6, 7, 5, 2, 4, 5, 9, 3},
			expected: 16,
		},
		// {
		// 	name:     "x2",
		// 	input:    []int{2, 4},
		// 	expected: 4,
		// },
		{
			name:     "x3",
			input:    []int{2},
			expected: 2,
		},
	}
	for _, c := range cases {
		assert.Equal(c.expected, largestRectangleArea1(c.input), "largestRectangleArea1->"+c.name)
		assert.Equal(c.expected, largestRectangleArea2(c.input), "largestRectangleArea2->"+c.name)
		assert.Equal(c.expected, largestRectangleArea3(c.input), "largestRectangleArea3->"+c.name)
	}
}
