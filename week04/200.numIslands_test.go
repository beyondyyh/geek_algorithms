package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_numIslands
func Test_numIslands(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]byte
		expect int
	}{
		{
			name: "x1",
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expect: 1,
		},
		{
			name: "x2",
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expect: 3,
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, numIslands(c.input), c.name)
	}
}
