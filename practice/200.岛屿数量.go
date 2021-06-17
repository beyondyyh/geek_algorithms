package practice

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 行、列数
	rows, cols := len(grid), len(grid[0])
	// 上下左右 offset
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	inGrid := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}
	var dfsMarking func(int, int)
	dfsMarking = func(x, y int) {
		if !inGrid(x, y) || '0' == grid[x][y] {
			return
		}
		grid[x][y] = '0'
		for _, direct := range directions {
			dfsMarking(x+direct[0], y+direct[1])
		}
	}

	var count int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if '1' == grid[i][j] {
				count++
				dfsMarking(i, j)
			}
		}
	}
	return count
}

// @lc code=end
