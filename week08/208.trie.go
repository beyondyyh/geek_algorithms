package week08

// 208. 实现 Trie (前缀树)
// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
// 请你实现 Trie 类：
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

// Trie是便于word插入和查找的数据结构
type Trie struct {
	children [26]*Trie // 26个字母映射表
	isLeaf   bool      // 该节点是否是一个单词的结尾
}

func Constructor() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, char := range word {
		if node.children[char-'a'] == nil {
			node.children[char-'a'] = &Trie{}
		}
		node = node.children[char-'a'] // 一层层进入子节点
	}
	node.isLeaf = true // 最终的子节点
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for _, char := range prefix {
		node = node.children[char-'a']
		if node == nil {
			return nil
		}
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isLeaf
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.searchPrefix(prefix)
	return node != nil
}
