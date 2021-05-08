package week01

import (
	"testing"

	"beyondyyh/geek_algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_mergeTwoSortedLists
func Test_mergeTwoSortedLists(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "x1",
			input:    [][]int{{1, 2, 4}, {1, 3, 4}},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "x2",
			input:    [][]int{{1, 2, 4}, {}},
			expected: []int{1, 2, 4},
		},
		{
			name:     "x3",
			input:    [][]int{{1, 3, 5}, {2, 4, 6, 7, 8}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			l1 := kit.Ints2List(c.input[0])
			l2 := kit.Ints2List(c.input[1])
			// assert.Equal(c.expected, kit.List2Ints(mergeTwoLists1(l1, l2)), "mergeTwoLists1->"+c.name)
			assert.Equal(c.expected, kit.List2Ints(mergeTwoLists2(l1, l2)), "mergeTwoLists2->"+c.name)
		})
	}
}
