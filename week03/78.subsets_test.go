package week03

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_subsets
func Test_subsets(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{
			name:  "x1",
			input: []int{1, 2, 3},
			// expect: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
			expect: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	// 二维数组升序排序
	sortSliceAsc := func(slice [][]int) {
		sort.Slice(slice, func(i, j int) bool {
			n1, n2 := len(slice[i]), len(slice[j])
			if n1 == n2 {
				for k := 0; k < n1; k++ {
					if slice[i][k] > slice[j][k] {
						return false
					}
				}
				return true
			}
			return n1 < n2
		})
	}

	for _, c := range cases {
		output1 := subsets1(c.input)
		sortSliceAsc(output1)
		assert.Equal(c.expect, output1, c.name)

		output2 := subsets2(c.input)
		sortSliceAsc(output2)
		assert.Equal(c.expect, output2, c.name)
	}
}
