package offer

// 剑指 Offer 29. 顺时针打印矩阵
// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]

// 示例 2：
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
// @lc: https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	res := []int{}
	rows, cols := len(matrix), len(matrix[0])
	up, down, left, right := 0, rows-1, 0, cols-1
	for {
		// 1. 从左向右遍历上面的行
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		// 2. 从上向下遍历右边的列
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		// 3. 从右向左遍历下面的行
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}

		// 4. 从下向上遍历左边的列
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
