package week03

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_buildTree
func Test_buildTree(t *testing.T) {
	type order struct {
		preorder []int
		inorder  []int
	}
	cases := []struct {
		name   string
		input  order
		expect []int
	}{
		{
			name: "x1",
			input: order{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			expect: []int{3, 9, 20, 15, 7},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		output1 := kit.Tree2BFS(buildTree(c.input.preorder, c.input.inorder))
		assert.Equal(c.expect, output1, c.name)

		// output2 := kit.Tree2BFS(buildTree1(c.input.preorder, c.input.inorder))
		// assert.Equal(c.expect, output2, c.name)
	}
}
