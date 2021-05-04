package week06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_maxProfit
func Test_maxProfit(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "x1",
			input:  []int{7, 1, 5, 3, 6, 4},
			expect: 5,
		},
		{
			name:   "x2",
			input:  []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, maxProfit1(c.input), "maxProfit1->"+c.name)
	}
}
