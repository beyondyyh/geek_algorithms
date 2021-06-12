package daily

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_longestValidParentheses
func Test_longestValidParentheses(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "x1",
			input:  "())(())",
			expect: 4,
		},
		{
			name:   "x1",
			input:  "()()(())",
			expect: 8,
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, longestValidParentheses1(c.input), c.name)
	}
}
