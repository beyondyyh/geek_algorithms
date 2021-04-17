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
			name:   "x1",
			input:  []int{1, 2, 3},
			expect: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
			// expect: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:   "x2",
			input:  []int{9, 0, 3, 5, 7},
			expect: [][]int{{}, {7}, {5}, {5, 7}, {3}, {3, 7}, {3, 5}, {3, 5, 7}, {0}, {0, 7}, {0, 5}, {0, 5, 7}, {0, 3}, {0, 3, 7}, {0, 3, 5}, {0, 3, 5, 7}, {9}, {9, 7}, {9, 5}, {9, 5, 7}, {9, 3}, {9, 3, 7}, {9, 3, 5}, {9, 3, 5, 7}, {9, 0}, {9, 0, 7}, {9, 0, 5}, {9, 0, 5, 7}, {9, 0, 3}, {9, 0, 3, 5}, {9, 0, 3, 5}, {9, 0, 3, 5, 7}},
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
		expected := c.expect
		sortSliceAsc(c.expect)

		output1 := subsets1(c.input)
		sortSliceAsc(output1)
		assert.Equal(expected, output1, c.name)

		output2 := subsets2(c.input)
		sortSliceAsc(output2)
		assert.Equal(expected, output2, c.name)
	}
}
