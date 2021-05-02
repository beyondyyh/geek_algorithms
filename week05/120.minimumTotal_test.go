package week05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_minimumTotal
func Test_minimumTotal(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name:   "x1",
			input:  [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			expect: 11,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, minimumTotal1(c.input), c.name)
	}
}
