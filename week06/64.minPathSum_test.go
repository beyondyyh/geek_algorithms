package week06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_minPathSum
func Test_minPathSum(t *testing.T) {
	cases := []struct {
		name   string
		grid   [][]int
		expect int
	}{
		{
			name: "x1",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expect: 7,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		// assert.Equal(c.expect, minPathSum1(c.grid), c.name)
		assert.Equal(c.expect, minPathSum2(c.grid), c.name)
	}
}
