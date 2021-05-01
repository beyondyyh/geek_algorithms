package week04

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_largestValues
func Test_largestValues(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "x1",
			input:  []int{3, 9, 20, null, null, 15, 7},
			expect: []int{3, 30, 15},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		root := kit.Ints2Tree(c.input)
		assert.Equal(c.expect, largestValues(root), c.name)
	}
}
