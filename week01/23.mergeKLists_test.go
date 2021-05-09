package week01

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_mergeKLists
func Test_mergeKLists(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect []int
	}{
		{
			name: "x1",
			input: [][]int{
				{1, 4, 5}, {1, 3, 4}, {2, 6},
			},
			expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name: "x1",
			input: [][]int{
				{1, 4, 5}, {}, {2, 6},
			},
			expect: []int{1, 2, 4, 5, 6},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		lists := make([]*ListNode, 0, len(c.input))
		for _, ints := range c.input {
			lists = append(lists, kit.Ints2List(ints))
		}
		output1 := kit.List2Ints(mergeKLists(lists))
		assert.Equal(c.expect, output1, "mergeKLists1->"+c.name)
	}
}
