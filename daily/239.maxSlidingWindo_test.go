package daily

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_maxSlidingWindow
func Test_maxSlidingWindow(t *testing.T) {
	type input struct {
		nums []int
		k    int
	}
	cases := []struct {
		name   string
		input  input
		expect []int
	}{
		{
			name: "x1随机",
			input: input{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			expect: []int{3, 3, 5, 5, 6, 7},
		},
		// {
		// 	name: "x2单调递增",
		// 	input: input{
		// 		nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
		// 		k:    3,
		// 	},
		// 	expect: []int{3, 4, 5, 6, 7, 8},
		// },
		// {
		// 	name: "x3递增+递减",
		// 	input: input{
		// 		nums: []int{9, 10, 9, -7, -4, -8, 2, -6},
		// 		k:    5,
		// 	},
		// 	expect: []int{10, 10, 9, 2},
		// },
	}
	assert := assert.New(t)
	for _, c := range cases {
		// assert.Equal(c.expect, maxSlidingWindow1(c.input.nums, c.input.k), c.name)
		assert.Equal(c.expect, maxSlidingWindow2(c.input.nums, c.input.k), c.name)
	}
}
