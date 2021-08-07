package practice

import "strings"

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	if n < 1 {
		return [][]string{}
	}

	// 每次遍历row时，之前的皇后所能攻击的位置（列、撇、捺）
	type flag map[int]bool
	cols, pie, na := make(flag), make(flag), make(flag)

	// results 表示所有有效的放置位置
	results := [][]int{}
	// 1. 遍历行，查看列、撇、捺位置是否存在攻击
	var dfsMarking func(int, []int)
	dfsMarking = func(row int, path []int) {
		// terminator
		if row >= n {
			results = append(results, append([]int{}, path...))
			return
		}

		// 遍历当前行上的每一列，依次看列、撇、捺上是否存在攻击
		for col := 0; col < n; col++ {
			if cols[col] || pie[row+col] || na[row-col] {
				continue
			}

			// update the flags
			cols[col], pie[row+col], na[row-col] = true, true, true
			path = append(path, col)

			// drill down
			dfsMarking(row+1, path)

			// revert state
			cols[col], pie[row+col], na[row-col] = false, false, false
			path = path[:len(path)-1]
		}
	}

	res := [][]string{}
	// 2. 根据results生成棋盘
	generateBoard := func() {
		for _, path := range results {
			borad := []string{}
			for _, pos := range path {
				row := strings.Repeat(".", pos) + "Q" + strings.Repeat(".", n-1-pos)
				borad = append(borad, row)
			}
			res = append(res, borad)
		}
	}

	// main logic
	dfsMarking(0, []int{})
	generateBoard()

	return res
}

// @lc code=end
