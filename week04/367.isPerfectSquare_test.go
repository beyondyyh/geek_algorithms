package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_isPerfectSquare
func Test_isPerfectSquare(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect bool
	}{
		{
			name:   "x1",
			input:  2,
			expect: false,
		},
		{
			name:   "x2",
			input:  16,
			expect: true,
		},
		{
			name:   "x1",
			input:  14,
			expect: false,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, isPerfectSquare(c.input), c.name)
	}
}
