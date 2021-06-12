package daily

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_isIsomorphic
func Test_isIsomorphic(t *testing.T) {
	cases := []struct {
		name   string
		input  [2]string
		expect bool
	}{
		{
			name:   "x1",
			input:  [2]string{"paper", "title"},
			expect: true,
		},
		{
			name:   "x2",
			input:  [2]string{"add", "egg"},
			expect: true,
		},
		{
			name:   "x3",
			input:  [2]string{"foo", "bar"},
			expect: false,
		},
	}

	assert := assert.New(t)
	for _, c := range cases {
		assert.Equal(c.expect, isIsomorphic(c.input[0], c.input[1]), c.name)
	}
}
