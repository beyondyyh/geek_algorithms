package week04

// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 示例 1：
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
// @leetcode: https://leetcode-cn.com/problems/search-a-2d-matrix

// 从右上角查找
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	r, c := 0, col-1
	for r < row && c >= 0 {
		if target == matrix[r][c] {
			return true
		} else if target < matrix[r][c] {
			c--
		} else {
			r++
		}
	}
	return false
}
