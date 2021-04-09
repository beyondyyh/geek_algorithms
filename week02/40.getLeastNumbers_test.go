package week02

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	entry40input struct {
		nums []int
		k    int
	}
	entry40 struct {
		name     string
		input    entry40input
		expected []int
	}
)

// run: go test -run Test_getLeastNumbers
func Test_getLeastNumbers(t *testing.T) {
	assert := assert.New(t)
	cases := []entry40{
		{
			name: "x1",
			input: entry40input{
				nums: []int{1, 3, 2},
				k:    2,
			},
			expected: []int{1, 2},
		},
		{
			name: "x2",
			input: entry40input{
				nums: []int{3, 2, 1},
				k:    0,
			},
			expected: nil,
		},
		{
			name: "x3",
			input: entry40input{
				nums: []int{0, 1, 2, 1},
				k:    1,
			},
			expected: []int{0},
		},
	}
	for _, c := range cases {
		output1 := getLeastNumbers1(c.input.nums, c.input.k)
		sort.Ints(output1)
		assert.Equal(c.expected, output1, "getLeastNumbers1->"+c.name)

		output2 := getLeastNumbers2(c.input.nums, c.input.k)
		sort.Ints(output2)
		assert.Equal(c.expected, output2, "getLeastNumbers2->"+c.name)

		output3 := getLeastNumbers2(c.input.nums, c.input.k)
		sort.Ints(output3)
		assert.Equal(c.expected, output3, "getLeastNumbers3->"+c.name)
	}
}
