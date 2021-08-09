package week07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_solveNQueens
func Test_solveNQueens(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect [][]string
	}{
		{
			name:  "n=4",
			input: 4,
			expect: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		// solveNQueens(c.input)
		assert.Equal(c.expect, solveNQueens(c.input), c.name)
	}
}
