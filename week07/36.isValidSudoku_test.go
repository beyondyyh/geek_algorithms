package week07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_isValidSudoku
func Test_isValidSudoku(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]byte
		expect bool
	}{
		{
			name: "x1",
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			expect: true,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, isValidSudoku(c.input), c.name)
	}
}
