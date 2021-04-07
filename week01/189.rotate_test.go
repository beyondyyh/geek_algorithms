package week01

import (
	"reflect"
	"testing"
)

type (
	entry189input struct {
		nums []int
		k    int
	}
	entry189 struct {
		name     string
		input    entry189input
		expected []int
	}
)

// run: go test -run Test_rotate
func Test_rotate(t *testing.T) {
	cases := []entry189{
		{
			name: "x1",
			input: entry189input{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "x2",
			input: entry189input{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expected: []int{3, 99, -1, -100},
		},
		{
			name: "x3",
			input: entry189input{
				nums: []int{1, -1},
				k:    1,
			},
			expected: []int{-1, 1},
		},
	}

	for _, c := range cases {
		// 注意slice底层是指针数组，rotate1之后会修改c.input.nums的顺序，这样的话rotate2的input就不对了，所以一次run一个
		t.Run(c.name, func(t *testing.T) {
			// rotate1(c.input.nums, c.input.k)
			// if !reflect.DeepEqual(c.input.nums, c.expected) {
			// 	t.Errorf("rotate1(%v,%d)=%v, expected=%v", c.input.nums, c.input.k, c.input.nums, c.expected)
			// }

			rotate2(c.input.nums, c.input.k)
			if !reflect.DeepEqual(c.input.nums, c.expected) {
				t.Errorf("rotate2(%v,%d)=%v, expected=%v", c.input.nums, c.input.k, c.input.nums, c.expected)
			}
		})
	}
}
