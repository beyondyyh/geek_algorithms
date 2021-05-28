package week08

// 79. 单词搜索
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 示例 1：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
// @lc: https://leetcode-cn.com/problems/word-search/

// DFS+回溯
func exist(board [][]byte, word string) bool {
	// 分别为 当前位置与 上、下、左、右 的offset
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	// 初始化，rows,cols,visited
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	// 坐标是否在 board 内，防止越界
	inBoard := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	// dfs marking func
	n := len(word)
	var dfs func(int, int, int) bool
	dfs = func(i, j, begin int) bool {
		if begin == n-1 { // terminator
			return board[i][j] == word[begin]
		}
		if board[i][j] == word[begin] {
			visited[i][j] = true                   // marking
			for _, direction := range directions { // process
				x := i + direction[0]
				y := j + direction[1]
				if inBoard(x, y) && !visited[x][y] {
					if dfs(x, y, begin+1) {
						return true
					}
				}
			}
			visited[i][j] = false // revert
		}
		return false
	}

	// dfs marking
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// AC过了，懒得写单测了
// @lc https://leetcode-cn.com/problems/word-search/solution/79-dan-ci-sou-suo-dfshui-su-golangban-be-ekcn/
