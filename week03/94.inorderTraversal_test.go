package week03

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_inorderTraversal
func Test_inorderTraversal(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "x1",
			input:  []int{1, null, 2, 3},
			expect: []int{1, 3, 2},
		},
	}
	for _, c := range cases {
		root := kit.Ints2Tree(c.input)
		assert.Equal(c.expect, inorderTraversal1(root), c.name)
		assert.Equal(c.expect, inorderTraversal2(root), c.name)
	}
}
