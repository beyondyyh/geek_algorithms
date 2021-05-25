package week08

// 695. 岛屿的最大面积
// 给定一个包含了一些 0 和 1 的非空二维数组 grid 。
// 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
// @lc: https://leetcode-cn.com/problems/max-area-of-island/

// dfs递归
func maxAreaOfIsland(grid [][]int) int {
	// 合法性判断
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 行、列数
	rows, cols := len(grid), len(grid[0])
	// 上下左右 与当前位置的offset
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// 判断当前位置是否在网格内
	inGrid := func(i, j int) bool {
		return i >= 0 && i < rows && j >= 0 && j < cols
	}

	var dfsMarking func(i, j int) int
	dfsMarking = func(i, j int) int {
		// terminator, 跳出网格 或 当前位置不是陆地 '1'，即被mark过了
		if !inGrid(i, j) || 1 != grid[i][j] {
			return 0
		}

		// mark 为'2'，且递归把他的上下左右已经子子孙孙都mark
		grid[i][j] = 2
		area := 1
		for _, direct := range directions {
			area += dfsMarking(i+direct[0], j+direct[1])
		}
		return area
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 遍历网格，求 maxArea
	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res = max(res, dfsMarking(i, j))
		}
	}
	return res
}
