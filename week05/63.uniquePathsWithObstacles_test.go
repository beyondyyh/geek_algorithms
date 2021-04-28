package week05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_uniquePathsWithObstacles
func Test_uniquePathsWithObstacles(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name:   "x0",
			input:  [][]int{{0}, {1}},
			expect: 0,
		},
		{
			name:   "x1",
			input:  [][]int{{1, 0}},
			expect: 0,
		},
		{
			name:   "x2",
			input:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expect: 2,
		},
		{
			name:   "x3",
			input:  [][]int{{0, 1}, {0, 0}},
			expect: 1,
		},
		{
			name:   "x4",
			input:  [][]int{{0, 0}, {1, 0}},
			expect: 1,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, uniquePathsWithObstacles(c.input), c.name)
		assert.Equal(c.expect, uniquePathsWithObstacles1(c.input), c.name)
		assert.Equal(c.expect, uniquePathsWithObstacles2(c.input), c.name)
	}
}
