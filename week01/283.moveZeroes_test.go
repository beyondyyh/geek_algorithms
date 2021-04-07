package week01

import (
	"reflect"
	"testing"
)

type entry283 struct {
	name     string
	input    []int
	expected []int
}

// run: go test -run Test_moveZeros
func Test_moveZeros(t *testing.T) {
	cases := []entry283{
		{
			name:     "x1",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "x2",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "x3",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// moveZeroes1(c.input)
			// if !reflect.DeepEqual(c.input, c.expected) {
			// 	t.Errorf("moveZeroes1(%v)=%v, expected=%v", c.input, c.input, c.expected)
			// }

			moveZeroes2(c.input)
			if !reflect.DeepEqual(c.input, c.expected) {
				t.Errorf("moveZeroes2(%v)=%v, expected=%v", c.input, c.input, c.expected)
			}
		})
	}
}
