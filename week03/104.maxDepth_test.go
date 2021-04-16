package week03

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_maxDepth
func Test_maxDepth(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "x1",
			input:  []int{null},
			expect: 1,
		},
		{
			name:   "x2",
			input:  []int{3, 9, 20, null, null, 15, 7},
			expect: 3,
		},
	}
	for _, c := range cases {
		assert.Equal(c.expect, maxDepth(kit.Ints2Tree(c.input)), c.name)
	}
}
