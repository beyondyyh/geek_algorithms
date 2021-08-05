package practice

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children [26]*Trie // 26字母映射表，ASCII码减a，[0,25]
	isEnd    bool
}

/** Initialize your data structure here. */
// func Constructor() Trie {
func NewTrie() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Trie{}
		}
		node = node.children[c-'a'] // 一层层进入子节点
	}
	node.isEnd = true // 遍历完单词的字符构建树后，把该结点标记为一个串的结束
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t.searchPrefix(prefix)
	return node != nil
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for _, c := range prefix {
		node = node.children[c-'a']
		if node == nil {
			return nil
		}
	}
	return node
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
