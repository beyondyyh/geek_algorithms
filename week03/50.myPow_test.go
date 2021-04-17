package week03

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_myPow
func Test_myPow(t *testing.T) {
	type input struct {
		x float64
		n int
	}
	assert := assert.New(t)
	cases := []struct {
		name   string
		input  input
		expect float64
	}{
		{
			name:   "x1",
			input:  input{x: 2, n: 10},
			expect: 1024,
		},
		{
			name:   "x2",
			input:  input{x: 2.1, n: 3},
			expect: 9.261,
		},
		{
			name:   "x3",
			input:  input{x: 2, n: -2},
			expect: 0.25,
		},
	}
	for _, c := range cases {
		output1, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", myPow1(c.input.x, c.input.n)), 64)
		assert.Equal(c.expect, output1, c.name)

		output2, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", myPow2(c.input.x, c.input.n)), 64)
		assert.Equal(c.expect, output2, c.name)
	}
}
