package week05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_maxProduct
func Test_maxProduct(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "x1",
			input:  []int{2, 3, -2, 4},
			expect: 6,
		},
		{
			name:   "x2",
			input:  []int{-2, 0, -1},
			expect: 0,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, maxProduct(c.input), c.name)
	}
}
