package week03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_replaceSpace
func Test_replaceSpace(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "x1",
			input:  "We are happy.",
			expect: "We%20are%20happy.",
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, replaceSpace1(c.input), c.name)
		assert.Equal(c.expect, replaceSpace2(c.input), c.name)
	}
}
