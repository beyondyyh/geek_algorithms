package week04

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
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
// @leetcode: https://leetcode-cn.com/problems/number-of-islands

// 方法一：深度优先搜索
// 时间复杂度：O(m*n) m,n分别为行和列
// 空间复杂度：O(m*n)，在最坏情况下，整个网格均为陆地，深度优先搜索的深度达到m*n
func numIslands(grid [][]byte) int {
	// 边界值判断
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 行、列数
	rows, cols := len(grid), len(grid[0])
	// 上下左右 与当前位置的offset
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 判断当前位置是否在网格内
	inGrid := func(i, j int) bool {
		return i >= 0 && i < rows && j >= 0 && j < cols
	}

	// dfsMarking 的作用是如果当前位置是'1'，则把它的前后左右都标记为'0'，
	// 以此类推，递归把他的子子孙孙都标记为'0'，也就是把相邻的陆地都夷为平地，
	// 只保留一个'1'，然后陆地个数count加一
	var dfsMarking func(int, int) // 谨遵泛型递归四部曲
	dfsMarking = func(i, j int) {
		// 1. terminaort 递归终止条件，i、j越界或当前位置不是'1'
		if !inGrid(i, j) || '1' != grid[i][j] {
			return
		}

		// 2. process current logic
		grid[i][j] = '0'
		// 3. drill down 递推下探将它的上下左右以及子子孙孙的上下左右都做dfsMarking操作
		// dfsMarking(i-1, j) // 上
		// dfsMarking(i+1, j) // 下
		// dfsMarking(i, j-1) // 左
		// dfsMarking(i, j+1) // 右

		// 以上4行代码可以简化，预先定义一个方向偏离量数组，防止循环里写即可
		for _, direct := range directions {
			dfsMarking(i+direct[0], j+direct[1])
		}

		// 4. revert states, nothing todo
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
