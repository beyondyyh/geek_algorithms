package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_combinationSum
func Test_combinationSum(t *testing.T) {
	assert := assert.New(t)
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
				nums:   []int{2, 3, 6, 7},
				target: 7,
			},
			expect: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "x2",
			input: input{
				nums:   []int{3, 5, 8},
				target: 11,
			},
			expect: [][]int{{3, 3, 5}, {3, 8}},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expect, combinationSum(c.input.nums, c.input.target), c.name)
	}
}
