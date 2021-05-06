package week06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_fourSum
func Test_fourSum(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	cases := []struct {
		name   string
		input  input
		expect [][]int
	}{
		{
			name: "x1",
			input: input{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			expect: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name: "x2",
			input: input{
				nums:   []int{-1, 0, 1, 2, -1, -4},
				target: -1,
			},
			expect: [][]int{{-4, 0, 1, 2}, {-1, -1, 0, 1}},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, fourSum1(c.input.nums, c.input.target), "fourSum1->"+c.name)
		assert.Equal(c.expect, fourSum2(c.input.nums, c.input.target), "fourSum2->"+c.name)
	}
}
