package week02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_nthUglyNumber
func Test_nthUglyNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "x1",
			input:  3,
			expect: 3,
		},
		{
			name:   "x2",
			input:  5,
			expect: 5,
		},
		{
			name:   "x3",
			input:  10,
			expect: 12,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, nthUglyNumber(c.input), c.name)
	}
}
