package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_findMinWithDup
func Test_findMinWithDup(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "x1",
			input:  []int{3, 4, 5, 1, 2},
			expect: 1,
		},
		{
			name:   "x1",
			input:  []int{2, 2, 2, 0, 1},
			expect: 0,
		},
		{
			name:   "x3",
			input:  []int{7, 0, 1, 1, 1, 1, 1, 2, 3, 4},
			expect: 0,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, findMinWithDup(c.input), c.name)
	}
}
