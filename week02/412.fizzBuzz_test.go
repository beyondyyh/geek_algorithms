package week02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entry412 struct {
	name     string
	input    int
	expected []string
}

// run: go test -run Test_fizzBuzz
func Test_fizzBuzz(t *testing.T) {
	assert := assert.New(t)
	cases := []entry412{
		{
			name:     "x1",
			input:    3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			name:     "x2",
			input:    6,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz"},
		},
		{
			name:     "x3",
			input:    15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, c := range cases {
		assert.Equal(c.expected, fizzBuzz1(c.input), c.name)
		assert.Equal(c.expected, fizzBuzz2(c.input), c.name)
	}
}
