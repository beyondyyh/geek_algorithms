package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_findMin
func Test_findMin(t *testing.T) {
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
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			expect: 0,
		},
		{
			name:   "x3",
			input:  []int{11, 13, 15, 17},
			expect: 11,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, findMin(c.input), c.name)
	}
}
