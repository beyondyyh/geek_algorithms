package week03

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_isValidBST
func Test_isValidBST(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect bool
	}{
		{
			name:   "x1",
			input:  []int{2, 1, 3},
			expect: true,
		},
		{
			name:   "x2",
			input:  []int{5, 1, 4, null, null, 3, 6},
			expect: false,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, isValidBST(kit.Ints2Tree(c.input)), c.name)
	}
}
