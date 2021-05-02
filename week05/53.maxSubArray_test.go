package week05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_maxSubArray
func Test_maxSubArray(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "x1",
			input:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect: 6,
		},
		{
			name:   "x2",
			input:  []int{1},
			expect: 1,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, maxSubArray(c.input), c.name)
	}
}
