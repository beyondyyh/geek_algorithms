package week05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_longestCommonSubsequence
func Test_longestCommonSubsequence(t *testing.T) {
	type input struct {
		text1 string
		text2 string
	}
	cases := []struct {
		name   string
		input  input
		expect int
	}{
		{
			name: "x1",
			input: input{
				text1: "abcde",
				text2: "ace",
			},
			expect: 3,
		},
		{
			name: "x2",
			input: input{
				text1: "abc",
				text2: "abc",
			},
			expect: 3,
		},
		{
			name: "x3",
			input: input{
				text1: "abc",
				text2: "def",
			},
			expect: 0,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, longestCommonSubsequence1(c.input.text1, c.input.text2), c.name)
		assert.Equal(c.expect, longestCommonSubsequence2(c.input.text1, c.input.text2), c.name)
	}
}
