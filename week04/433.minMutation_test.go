package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_minMutation
func Test_minMutation(t *testing.T) {
	type input struct {
		start string
		end   string
		bank  []string
	}
	cases := []struct {
		name   string
		input  input
		expect int
	}{
		// {
		// 	name: "x1",
		// 	input: input{
		// 		start: "AACCGGTT",
		// 		end:   "AACCGGTA",
		// 		bank:  []string{"AACCGGTA"},
		// 	},
		// 	expect: 1,
		// },
		// {
		// 	name: "x2",
		// 	input: input{
		// 		start: "AACCGGTT",
		// 		end:   "AAACGGTA",
		// 		bank:  []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
		// 	},
		// 	expect: 2,
		// },
		// {
		// 	name: "x3",
		// 	input: input{
		// 		start: "AAAAACCC",
		// 		end:   "AACCCCCC",
		// 		bank:  []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
		// 	},
		// 	expect: 3,
		// },
		{
			name: "x4",
			input: input{
				start: "AACCTTGG",
				end:   "AATTCCGG",
				bank:  []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"},
			},
			expect: 2,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, minMutation1(c.input.start, c.input.end, c.input.bank), c.name)
		assert.Equal(c.expect, minMutation2(c.input.start, c.input.end, c.input.bank), c.name)
	}
}
