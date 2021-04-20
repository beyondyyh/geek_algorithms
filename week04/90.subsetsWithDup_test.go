package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_subsetsWithDup
func Test_subsetsWithDup(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{
			name:   "x1",
			input:  []int{1, 2, 2},
			expect: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expect, subsetsWithDup(c.input), c.name)
	}
}
