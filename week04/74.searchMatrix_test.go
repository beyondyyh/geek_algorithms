package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entry04input struct {
	name     string
	input    int
	expected bool
}

// run: go test -run Test_searchMatrix
func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	cases := []entry04input{
		{
			name:     "x1",
			input:    5,
			expected: true,
		},
		{
			name:     "x2",
			input:    18,
			expected: true,
		},
		{
			name:     "x3",
			input:    20,
			expected: false,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expected, searchMatrix(matrix, c.input), c.name)
	}
}
