package week08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_permutation
func Test_permutation(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect []string
	}{
		{
			name:   "x1",
			input:  "abc",
			expect: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name:   "x2",
			input:  "aba",
			expect: []string{"aab", "aba", "baa"},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, permutation(c.input), c.name)
	}
}
