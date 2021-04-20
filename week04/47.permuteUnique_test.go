package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_permuteUnique
func Test_permuteUnique(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{
			name:   "x1",
			input:  []int{1, 1, 2},
			expect: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expect, permuteUnique(c.input), c.name)
	}
}
