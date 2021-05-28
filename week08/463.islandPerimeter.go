package week08

// 463. 岛屿的周长
// 给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
// 网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
// 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
// @lc: https://leetcode-cn.com/problems/island-perimeter/

func islandPerimeter(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	// 是否在区域内
	inGrid := func(i, j int) bool {
		return i >= 0 && i < rows && j >= 0 && j < cols
	}

	// 上下左右 offset
	directions := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	var dfsMarking func(int, int) int
	dfsMarking = func(i, j int) int {
		// 走到网格边上了，返回，对应一条黄色的边
		if !inGrid(i, j) {
			return 1
		}
		// 遇到水了，返回，对应一条黄色的边
		if 0 == grid[i][j] {
			return 1
		}
		// 不是陆地，不需要marking，同时也跟周长没关系
		if 1 != grid[i][j] {
			return 0
		}

		// marking
		grid[i][j] = 2
		var area int
		for _, direct := range directions {
			x, y := i+direct[0], j+direct[1]
			area += dfsMarking(x, y)
		}
		return area
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 题目 只有一个岛屿只需要mark一次
			if 1 == grid[i][j] {
				return dfsMarking(i, j)
			}
		}
	}
	return 0
}
