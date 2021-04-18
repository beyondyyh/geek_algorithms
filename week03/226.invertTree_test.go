package week03

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_invertTree
func Test_invertTree(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "x1",
			input:  []int{4, 2, 7, 1, 3, 6, 9},
			expect: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:   "x2",
			input:  []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
			expect: []int{3, 1, 5, 8, 0, 2, 6, 4, 7},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		root := kit.Ints2Tree(c.input)
		output := kit.Tree2BFS(invertTree(root))
		assert.Equal(c.expect, output, c.name)
	}
}
