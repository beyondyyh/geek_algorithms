package week06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_robII
func Test_robII(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "x1",
			input:  []int{2, 3, 2},
			expect: 3,
		},
		{
			name:   "x2",
			input:  []int{1, 2, 3, 1},
			expect: 4,
		},
		{
			name:   "x3",
			input:  []int{0},
			expect: 0,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, robII1(c.input), "robII1->"+c.name)
		assert.Equal(c.expect, robII2(c.input), "robII1->"+c.name)
	}
}
