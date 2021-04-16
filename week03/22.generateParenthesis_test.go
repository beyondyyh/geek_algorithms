package week03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_generateParenthesis
func Test_generateParenthesis(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  int
		expect []string
	}{
		// {
		// 	name:   "x1",
		// 	input:  1,
		// 	expect: []string{"()"},
		// },
		// {
		// 	name:   "x2",
		// 	input:  2,
		// 	expect: []string{"()()", "(())"},
		// },
		{
			name:   "x3",
			input:  3,
			expect: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expect, generateParenthesis(c.input), c.name)
	}
}
