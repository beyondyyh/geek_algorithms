package week02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	entry242input [2]string
	entry242      struct {
		name     string
		input    entry242input
		expected bool
	}
)

// run: go test -run Test_isAnagram
func Test_isAnagram(t *testing.T) {
	assert := assert.New(t)
	cases := []entry242{
		{
			name:     "x1",
			input:    [2]string{"anagram", "nagaram"},
			expected: true,
		},
		{
			name:     "x2",
			input:    [2]string{"rat", "car"},
			expected: false,
		},
	}
	for _, c := range cases {
		assert.Equal(c.expected, isAnagram1(c.input[0], c.input[1]), "isAnagram1->"+c.name)
		assert.Equal(c.expected, isAnagram2(c.input[0], c.input[1]), "isAnagram2->"+c.name)
		assert.Equal(c.expected, isAnagram3(c.input[0], c.input[1]), "isAnagram3->"+c.name)
	}
}
