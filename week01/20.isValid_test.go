package week01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entry20 struct {
	name     string
	input    string
	expected bool
}

// run: go test -run Test_isValid
func Test_isValid(t *testing.T) {
	assert := assert.New(t)
	cases := []entry20{
		{
			name:     "x1",
			input:    "()",
			expected: true,
		},
		{
			name:     "x2",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "x3",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "x4",
			input:    "{[]}",
			expected: true,
		},
	}
	for _, c := range cases {
		assert.Equal(c.expected, isValid1(c.input), c.name)
		assert.Equal(c.expected, isValid2(c.input), c.name)
	}
}
