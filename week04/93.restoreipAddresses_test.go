package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_restoreipAddresses
func Test_restoreipAddresses(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  string
		expect []string
	}{
		{
			name:   "x1",
			input:  "0000",
			expect: []string{"0.0.0.0"},
		},
		{
			name:   "x2",
			input:  "25525511135",
			expect: []string{"255.255.11.135", "255.255.111.35"},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expect, restoreIpAddresses(c.input), c.name)
		assert.Equal(c.expect, restoreIpAddresses1(c.input), c.name)
	}
}
