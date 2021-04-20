package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_permute
func Test_permute(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{
			name:   "x1",
			input:  []int{1, 2, 3},
			expect: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expect, permute(c.input), c.name)
	}
}
