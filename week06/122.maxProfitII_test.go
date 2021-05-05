package week06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_maxProfitII
func Test_maxProfitII(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "x1",
			input:  []int{7, 1, 5, 3, 6, 4},
			expect: 7,
		},
		{
			name:   "x2",
			input:  []int{1, 2, 3, 4, 5},
			expect: 4,
		},
		{
			name:   "x3",
			input:  []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, maxProfitII1(c.input), "maxProfitII1->"+c.name)
		assert.Equal(c.expect, maxProfitII2(c.input), "maxProfitII2->"+c.name)
		assert.Equal(c.expect, maxProfitII3(c.input), "maxProfitII3->"+c.name)
		assert.Equal(c.expect, maxProfitII4(c.input), "maxProfitII4->"+c.name)
	}
}
