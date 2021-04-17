package week03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_minPathSum
func Test_minPathSum(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name:   "x1",
			input:  [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expect: 7,
		},
		{
			name:   "x2",
			input:  [][]int{{1, 2, 3}, {4, 5, 6}},
			expect: 12,
		},
	}

	for _, c := range cases {
		assert.Equal(c.expect, minPathSum(c.input), c.name)
	}
}
