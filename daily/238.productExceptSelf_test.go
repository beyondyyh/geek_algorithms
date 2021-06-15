package daily

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_productExceptSelf
func Test_productExceptSelf(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "x1",
			input:  []int{1, 2, 3, 4},
			expect: []int{24, 12, 8, 6},
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, productExceptSelf(c.input), c.name)
	}
}
