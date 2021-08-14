package offer

// 剑指 Offer 04. 二维数组中的查找
// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
// 示例:
// 现有矩阵 matrix 如下：
// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
// 给定 target = 5，返回 true。
// 给定 target = 20，返回 false。
// @lc: https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

// 对角线查找，从右上角往左下角查找
// 时间复杂度：O(m+n) m,n表示矩阵的行列数，最多循环m+n次，比如从右上角查找时，target刚好位于左下角
// 空间复杂度：O(1)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	// 从右上角开始查找
	i, j := 0, cols-1
	for i < rows && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target > matrix[i][j] {
			i++
		} else {
			j--
		}
	}
	return false
}
