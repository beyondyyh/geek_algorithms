package week01

import (
	"reflect"
	"testing"

	"beyondyyh/geek_algorithms/kit"
)

// run: go test -run Test_mergeTwoSortedLists
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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			l1 := kit.Ints2List(c.input[0])
			l2 := kit.Ints2List(c.input[1])
			output := mergeTwoLists(l1, l2)
			out2ints := kit.List2Ints(output)
			if !reflect.DeepEqual(out2ints, c.expected) {
				t.Errorf("mergeTwoLists(%v, %v)=%v, expected=%v", c.input[0], c.input[1], out2ints, c.expected)
			}
		})
	}
}
