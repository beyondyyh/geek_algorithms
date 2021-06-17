package daily

// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
// 示例 1：
// 输入：grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// 输出：1
// @lc: https://leetcode-cn.com/problems/number-of-islands/

// 方法一：dfs
// 时间复杂度：O(m*n)，m,n分别为行数和列数。
// 空间复杂度：O(m*n)，在最坏情况下，整个网格均为陆地，深度优先搜索的深度达到 m*n
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 行、列数
	rows, cols := len(grid), len(grid[0])
	// 当前坐标的 上下左右 offset，不然需要写4行雷同代码
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// inGrid 判断当前坐标是否在网格内
	inGrid := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}
	// dfsMarking 将当前坐标的上下左右都标记为'0'
	var dfsMarking func(int, int)
	dfsMarking = func(x, y int) {
		// treminator
		if !inGrid(x, y) || '0' == grid[x][y] {
			return
		}
		// mark将当前坐标
		grid[x][y] = '0'
		// 递归将当前坐标的 4个方向 都标记为'0'
		for _, direct := range directions {
			dfsMarking(x+direct[0], y+direct[1])
		}
	}

	// 遍历网格，找到'1'
	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if '1' == grid[i][j] {
				res++
				dfsMarking(i, j)
			}
		}
	}
	return res
}
