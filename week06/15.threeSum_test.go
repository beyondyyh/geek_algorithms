package week06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_threeSum
func Test_threeSum(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{
			name:   "x1",
			input:  []int{-1, 0, 1, 2, -1, -4},
			expect: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:   "x2",
			input:  []int{0},
			expect: [][]int{},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, threeSum1(c.input), "threeSum1->"+c.name)
		assert.Equal(c.expect, threeSum2(c.input), "threeSum2->"+c.name)
	}
}
