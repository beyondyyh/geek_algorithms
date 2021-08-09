package week07

import (
	"strings"
)

// 51. N 皇后
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
// 示例 1：
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
// @lc: https://leetcode-cn.com/problems/n-queens/

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
			cols[col], pie[row+col], na[row-col] = true, true, true    // update the flags
			dfsMarking(row+1, append(path, col))                       // drill down
			cols[col], pie[row+col], na[row-col] = false, false, false // revert state
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
