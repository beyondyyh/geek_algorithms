package week08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_trie
func Test_trie(t *testing.T) {
	assert := assert.New(t)
	trie := Constructor()

	trie.Insert("apple")
	assert.True(trie.Search("apple"), "Search apple")
	assert.False(trie.Search("app"), "Search app")
	assert.True(trie.StartsWith("app"), "StartsWith app")

	trie.Insert("app")
	assert.True(trie.Search("app"), "Search app")
}
