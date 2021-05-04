package week05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_coinChange
func Test_coinChange(t *testing.T) {
	type input struct {
		coins  []int
		amount int
	}
	cases := []struct {
		name   string
		input  input
		expect int
	}{
		{
			name: "x1",
			input: input{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			expect: 3,
		},
		{
			name: "x2",
			input: input{
				coins:  []int{2},
				amount: 3,
			},
			expect: -1,
		},
		{
			name: "x3",
			input: input{
				coins:  []int{1},
				amount: 0,
			},
			expect: 0,
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, coinChange1(c.input.coins, c.input.amount), "coinChange1->"+c.name)
		assert.Equal(c.expect, coinChange2(c.input.coins, c.input.amount), "coinChange2->"+c.name)
	}
}
