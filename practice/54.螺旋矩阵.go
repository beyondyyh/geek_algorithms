/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 行、列数
	rows, cols := len(matrix), len(matrix[0])
	// 存放结果集
	res := make([]int, 0, rows*cols)
	// 定义 上下左右 边界
	up, down, left, right := 0, rows-1, 0, cols-1
	// 不断遍历
	for {
		// 1.从左向右遍历上面的行
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++ // 上边界下移一位，如果重设后的上边界大于现有的下边界，则退出
		if up > down {
			break
		}

		// 2.从上往下遍历右边的列
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right-- // 右边界左移一位，如果重设后的右边界小于现有的左边界，则退出
		if right < left {
			break
		}

		// 3.从右往左遍历下边的行
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down-- // 下边界上移一位，如果重设后的下边界小于现有的上边界，则退出
		if down < up {
			break
		}

		// 4.从下往上遍历左边的列
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++ // 左边界右移一位，如果重设后的左边界大于右边界，则退出
		if left > right {
			break
		}
	}
	return res
}

// @lc code=end