package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_search
func Test_search(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	cases := []struct {
		name   string
		input  input
		expect int
	}{
		{
			name: "x1",
			input: input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			expect: 4,
		},
		{
			name: "x2",
			input: input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			expect: -1,
		},
		{
			name: "x2",
			input: input{
				nums:   []int{5, 1, 3},
				target: 3,
			},
			expect: 2,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, search(c.input.nums, c.input.target), c.name)
	}
}
