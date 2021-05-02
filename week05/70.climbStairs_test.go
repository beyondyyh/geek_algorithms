package week05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_climbStairs
func Test_climbStairs(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "x1",
			input:  2,
			expect: 2,
		},
		{
			name:   "x2",
			input:  5,
			expect: 8,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, climbStairs1(c.input), "climbStairs1->"+c.name)
		assert.Equal(c.expect, climbStairs2(c.input), "climbStairs2->"+c.name)
		assert.Equal(c.expect, climbStairs3(c.input), "climbStairs3->"+c.name)
	}
}
