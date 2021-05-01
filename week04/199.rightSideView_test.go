package week04

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_rightSideView
func Test_rightSideView(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "x1",
			input:  []int{1, 2, 3, null, 5, null, 4},
			expect: []int{1, 3, 4},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		root := kit.Ints2Tree(c.input)
		assert.Equal(c.expect, rightSideView(root), c.name)
	}
}
