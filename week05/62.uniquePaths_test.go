package week05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_uniquePaths
func Test_uniquePaths(t *testing.T) {
	type input struct {
		m int
		n int
	}
	cases := []struct {
		name   string
		input  input
		expect int
	}{
		{
			name: "x1",
			input: input{
				m: 3,
				n: 7,
			},
			expect: 28,
		},
		{
			name: "x2",
			input: input{
				m: 3,
				n: 2,
			},
			expect: 3,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, uniquePaths1(c.input.m, c.input.n), "uniquePaths1->"+c.name)
		assert.Equal(c.expect, uniquePaths2(c.input.m, c.input.n), "uniquePaths2->"+c.name)
	}
}
