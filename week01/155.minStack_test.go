package week01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -run Test_minStack
func Test_minStack(t *testing.T) {
	assert := assert.New(t)
	s := NewMinStack()

	s.push(-2)
	s.push(0)
	s.push(-3)
	assert.Equal(-3, s.getMin())
	s.pop()
	assert.Equal(0, s.top())
	assert.Equal(-2, s.getMin())
}
