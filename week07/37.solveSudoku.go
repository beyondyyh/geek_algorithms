package week07

// 37. 解数独
// 编写一个程序，通过填充空格来解决数独问题。
// 数独的解法需 遵循如下规则：
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
// 数独部分空格内已填入了数字，空白格用 '.' 表示。
// @lc: https://leetcode-cn.com/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}

	solve(board)
}

func solve(board [][]byte) bool {
	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == '.' {
				// 尝试填入1-9，没填入1个数就判断一下是否有效，有效则下一步，无效则backtracking
				var c byte
				for c = '1'; c <= '9'; c++ {
					if isValid(board, i, j, c) { // 尝试放入c后如果棋盘有效，则put it here
						board[i][j] = c // Put c for this cell

						// 放入c之后，递归 drill down
						if solve(board) {
							return true
						} else {
							board[i][j] = '.' // 回撤
						}

					} // end isValid
				} // end for 1-9

				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		// check列，即每一行上的第col列是否为出现过c
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		// check行，即第row行上的每一列是否为出现过c
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		// check 3*3 block
		boxRow, boxCol := (row/3)*3+i/3, (col/3)*3+i%3
		if board[boxRow][boxCol] != '.' && board[boxRow][boxCol] == c {
			return false
		}
	}
	return true
}
