package week03

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 示例 1：
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//  1	3	1
//	1	5 	1
// 	4 	2 	1
// 输出：7 解释：因为路径 1→3→1→1→1 的总和最小。
// 同理可得：最大和为：12 解释：因为路径 1->3->5->2->1 的总和最大。
// @leetcode: https://leetcode-cn.com/problems/minimum-path-sum

// 动态规划
// 当前节点(坐标:x,y)只能从上面节点(坐标:x-1,y)或左边节点(x,y-1)到达，因此总和最小应该是从二者值较小的点走过来
// 故function: f(x,y) = min(f(x-1,y), f(x,y-1)) + grid[x][y]
// 时间复杂度：O(m*n)
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 初始化：f(0,0) = A[0][0], f(i,0) = sum(0,0 -> i,0)、 f(0,y) = sum(0,0 -> 0,y)
	for x := 1; x < len(grid); x++ {
		grid[x][0] = grid[x][0] + grid[x-1][0]
	}
	for y := 1; y < len(grid[0]); y++ {
		grid[0][y] = grid[0][y] + grid[0][y-1]
	}

	// 遍历求解 f[x][y] 表示 x,y 到0,0的和的最小值
	for x := 1; x < len(grid); x++ {
		for y := 1; y < len(grid[0]); y++ {
			grid[x][y] = min(grid[x-1][y], grid[x][y-1]) + grid[x][y]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
