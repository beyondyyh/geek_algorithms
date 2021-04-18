package week03

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_reversePrint
func Test_reversePrint(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "x",
			input:  []int{1, 3, 2},
			expect: []int{2, 3, 1},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		head := kit.Ints2List(c.input)
		assert.Equal(c.expect, reversePrint1(head), c.name)
		assert.Equal(c.expect, reversePrint2(head), c.name)
	}
}
