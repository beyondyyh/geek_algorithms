package week08

// 212. 单词搜索 II
// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
// @lc: https://leetcode-cn.com/problems/word-search-ii/

// 方法一：遍历words，每个单词都去board中执行一次 dfs+回溯 (79. 单词搜索)。。。
// 时间复杂度：O(N* m*n * 4^k), N:words长度，m*n:表示board为m*n的矩阵，4:表示上下左右4个方向，k:表示单词的平级长度，因为每个字符都会走一次dfsMarking

// 方法二：trie + dfs+回溯
// 1. 所有单词先添加到 trie树 中，构建前缀树
// 2. 遍历board，4联通生成一个单词，去前缀树中搜索看是否存在

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return res
	}

	rows, cols := len(board), len(board[0])                  // 网格行、列数
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右 与当前坐标的offset

	// 当前坐标是否在网格内
	inBoard := func(i, j int) bool {
		return i >= 0 && i < rows && j >= 0 && j < cols
	}
	// 标记已经访问过的坐标
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	// 1. 把所有单词构造成前缀树
	trie := &Trie{}
	for _, word := range words {
		trie.Insert(word)
	}

	// 2. dfsMarking
	var dfs func(int, int, []byte)
	dfs = func(i, j int, path []byte) {
		// 1. terminator, 当前坐标越界 或 已经访问过，返回
		if !inBoard(i, j) || visited[i][j] {
			return
		}

		// // 剪枝：board中不存在该字符，返回
		// c := board[i][j]
		// if !trie.StartsWith(string(c)) {
		// 	return
		// }

		// 2. trie 存在就可以设置访问标志，添加路径
		visited[i][j] = true
		path = append(path, board[i][j])
		if trie.Search(string(path)) {
			res = append(res, string(path))
			return
		}
		// 3. drill down
		for _, direct := range directions {
			dfs(i+direct[0], j+direct[1], path)
		}
		// 3. revert 回撤
		visited[i][j] = false
		path = path[:len(path)-1]
	}

	// 3. 对矩阵每一个点进行回溯，找出所有在前缀树中的单词，注意找到了之后要修改trie对应节点的 isLeaf
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, []byte{})
		}
	}
	return res
}
