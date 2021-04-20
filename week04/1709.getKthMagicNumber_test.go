package week04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getKthMagicNumber(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "x1",
			input:  3,
			expect: 5,
		},
		{
			name:   "x2",
			input:  5,
			expect: 9,
		},
		{
			name:   "x3",
			input:  6,
			expect: 15,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, getKthMagicNumber1(c.input), c.name)
	}
}
