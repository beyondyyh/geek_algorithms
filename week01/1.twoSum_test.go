package week01

import (
	"reflect"
	"sort"
	"testing"
)

type (
	entry1 struct {
		name     string
		input    entry1input
		expected []int
	}
	entry1input struct {
		nums   []int
		target int
	}
)

// run: go test -run Test_twoSum
func Test_twoSum(t *testing.T) {
	cases := []entry1{
		{
			name: "x1",
			input: entry1input{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			name: "x2",
			input: entry1input{
				nums:   []int{3, 2, 4},
				target: 8,
			},
			expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output1 := twoSum1(c.input.nums, c.input.target)
			sort.Ints(output1)
			sort.Ints(c.expected)
			if !reflect.DeepEqual(output1, c.expected) {
				t.Errorf("twoSum1(%v, %d)=%v, expected=%v", c.input.nums, c.input.target, output1, c.expected)
			}

			output2 := twoSum2(c.input.nums, c.input.target)
			sort.Ints(output2)
			sort.Ints(c.expected)
			if !reflect.DeepEqual(output2, c.expected) {
				t.Errorf("twoSum2(%v, %d)=%v, expected=%v", c.input.nums, c.input.target, output2, c.expected)
			}
		})
	}
}
