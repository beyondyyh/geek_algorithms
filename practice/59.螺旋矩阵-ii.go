package practice

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	// 生成 n*n 的矩阵，初始值都为0
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 定义 上下左右 边界
	up, down, left, right := 0, n-1, 0, n-1
	num, tar := 1, n*n
	for num <= tar {
		// 1.从左往右遍历，并更新上边界
		for i := left; i <= right; i++ {
			matrix[up][i] = num
			num++
		}
		up++
		// 2.从上往下遍历，并更新右边界
		for i := up; i <= down; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// 3.从右往左遍历，并更新下边界
		for i := right; i >= left; i-- {
			matrix[down][i] = num
			num++
		}
		down--
		// 4.从下往上遍历，并更新左边界
		for i := down; i >= up; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

// @lc code=end
