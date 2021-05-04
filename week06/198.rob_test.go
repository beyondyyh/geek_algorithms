package week06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_rob
func Test_rob(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "x1",
			input:  []int{1, 2, 3, 1},
			expect: 4,
		},
		{
			name:   "x2",
			input:  []int{2, 7, 9, 3, 1},
			expect: 12,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, rob1(c.input), "rob1->"+c.name)
		assert.Equal(c.expect, rob2(c.input), "rob2->"+c.name)
		assert.Equal(c.expect, rob3(c.input), "rob3->"+c.name)
	}
}
