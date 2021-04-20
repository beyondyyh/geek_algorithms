package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_combine
func Test_combine(t *testing.T) {
	cases := []struct {
		name   string
		input  [2]int
		expect [][]int
	}{
		{
			name:   "x1",
			input:  [2]int{4, 2},
			expect: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, combine(c.input[0], c.input[1]), c.name)
	}
}
