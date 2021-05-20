package week07

// 36. 有效的数独
// 请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
// 数独部分空格内已填入了数字，空白格用 '.' 表示
// @lc: https://leetcode-cn.com/problems/valid-sudoku/

// 定义三个map[byte]bool数组，每个数组中有9个map，分别是row、col和box的数独
// 遍历如果在row、col和box中都未出现过则继续下一个，否则返回False，直到遍历结束返回True
// box的索引：boxIndex = (row/3)*3 + col/3
func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]bool, 9)
	col := make([]map[byte]bool, 9)
	box := make([]map[byte]bool, 9)
	// 初始化，都设为false
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool, 9) // 没必要指定make的size，可能填不满就判断为Fase退出了
		col[i] = make(map[byte]bool, 9)
		box[i] = make(map[byte]bool, 9)
	}
	// 遍历
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			char := board[i][j]
			if char == '.' {
				continue
			}
			// 在row、col、box任何一个中存在，返回False
			boxIndex := (i/3)*3 + j/3
			if row[i][char] || col[j][char] || box[boxIndex][char] {
				return false
			}
			// 分别向row、col、box中添加改数
			row[i][char] = true
			col[j][char] = true
			box[boxIndex][char] = true
		}
	}
	return true
}
