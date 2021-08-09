package practice

/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	res := []string{}
	if len(board) == 0 || len(board[0]) == 0 {
		return res
	}

	// 1.将words构建字典树
	trie := newMyTrie()
	for _, word := range words {
		trie.insert(word)
	}

	// 当前坐标(x,y)的上下左右四个方向偏移量
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// 判断当前坐标是否在 board 内
	rows, cols := len(board), len(board[0])
	inBoard := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	// 2.1 以当前字符 board[x][y] 开始进行dfs搜索
	var dfsMarking func(int, int, string, *myTrie)
	dfsMarking = func(x, y int, word string, trie *myTrie) {
		if !inBoard(x, y) || '#' == board[x][y] || trie == nil {
			return
		}

		char := board[x][y]
		word += string(char)
		trie = trie.children[char-'a'] // 进入子节点
		if trie != nil && trie.isEnd { // trie.Search(word) 是一个单词
			res = append(res, word)
			trie.isEnd = false // 防止一个单词被重复添加
		}

		// 将当前位置标记为 '#'
		board[x][y] = '#'
		for _, direct := range directions {
			newX, newY := x+direct[0], y+direct[1]
			dfsMarking(newX, newY, word, trie)
		}
		// revert
		board[x][y] = char
	}

	// 2.遍历board进行dfsMarking
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfsMarking(i, j, "", trie)
		}
	}

	return res
}

// 字典树、前缀树
type myTrie struct {
	children [26]*myTrie // 26个字母映射表
	isEnd    bool        // 该节点是否是单词
}

func newMyTrie() *myTrie {
	return &myTrie{}
}

func (t *myTrie) insert(word string) {
	node := t
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = newMyTrie()
		}
		node = node.children[c-'a']
	}
	node.isEnd = true
}

// @lc code=end
