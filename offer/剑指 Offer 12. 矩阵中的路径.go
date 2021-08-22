package offer

// 剑指 Offer 12. 矩阵中的路径
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。

// 示例 1：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true

// 示例 2：
// 输入：board = [["a","b"],["c","d"]], word = "abcd"
// 输出：false
// @lc: https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/

// dfs+回溯
// 将访问过得位置标记为'#'避免重复访问，会改变board
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// 当前坐标(x,y)是否在board界内
	inBoard := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	// 从当前坐标(x, y) 分别向上下左右扩散
	// - x, y: 当前坐标
	// - k: 扩散到的单词长度索引
	var dfsMarking func(int, int, int) bool
	dfsMarking = func(x, y, k int) bool {
		// terminator，当扩散的深度和word长度相等时递归结束。返回当前坐标字符与目标字符是否相同
		if k == len(word)-1 {
			return board[x][y] == word[k]
		}

		// 否则还可以继续扩散：如果当前坐标字符与目标字符相同，则继续扩散
		if board[x][y] == word[k] {
			// 当前坐标标记为已访问
			board[x][y] = '#'
			// 分别向上下左右4个方向扩散
			for _, direct := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nextX, nextY := x+direct[0], y+direct[1]
				// 没越界且没访问过，则扩散，深度加一
				if inBoard(nextX, nextY) && board[nextX][nextY] != '#' {
					if dfsMarking(nextX, nextY, k+1) {
						return true
					}
				}
			}
			// revert 回复当前状态
			board[x][y] = word[k]
		}
		// 棋盘都搞完还没找到，返回false
		return false
	}

	// main process
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfsMarking(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// dfs+回溯，用visited数组记录访问过得位置，不改变board
func exist2(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// 当前坐标(x,y)是否在board界内
	inBoard := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}
	// 标记访问过的位置
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// 从当前坐标(x, y) 分别向上下左右扩散
	// - x, y: 当前坐标
	// - k: 扩散到的单词长度索引
	var dfsMarking func(int, int, int) bool
	dfsMarking = func(x, y, k int) bool {
		// terminator，当扩散的深度和word长度相等时递归结束。返回当前坐标字符与目标字符是否相同
		if k == len(word)-1 {
			return board[x][y] == word[k]
		}

		// 否则还可以继续扩散：如果当前坐标字符与目标字符相同，则继续扩散
		if board[x][y] == word[k] {
			// 当前坐标标记为已访问
			visited[x][y] = true
			// 分别向上下左右4个方向扩散
			for _, direct := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nextX, nextY := x+direct[0], y+direct[1]
				// 没越界且没访问过，则扩散，深度加一
				if inBoard(nextX, nextY) && !visited[nextX][nextY] {
					if dfsMarking(nextX, nextY, k+1) {
						return true
					}
				}
			}
			// revert 回复当前状态
			visited[x][y] = false
		}
		// 棋盘都搞完还没找到，返回false
		return false
	}

	// main process
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfsMarking(i, j, 0) {
				return true
			}
		}
	}
	return false
}
