package week03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_letterCombinations
func Test_letterCombinations(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  string
		expect []string
	}{
		{
			name:   "x1",
			input:  "",
			expect: []string{},
		},
		{
			name:   "x2",
			input:  "2",
			expect: []string{"a", "b", "c"},
		},
		{
			name:   "x1",
			input:  "23",
			expect: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expect, letterCombinations1(c.input), c.name)
	}
}
