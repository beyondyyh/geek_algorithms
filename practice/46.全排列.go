package practice

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
// 全排列 问题，典型的回溯思想

// -----------------------------------------------------------------------------------
// 回溯通用模板
// func backtracking(选择的起始位置, 走过的路径) {
// 	  if 满足递归终止条件 {
// 		result.add(走过的路径)
// 		return
// 	  }
// 	  for i := 选择的起始位置; i < 选择列表长度; i++ {
// 		做选择，路径.append(列表[i])
// 		backtracking(i+1, 路径)
// 		撤销选择
// 	  }
// }
// -----------------------------------------------------------------------------------

// 因为是全排列，每次都从0开始选择，pos不用传
func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	// 标记已经访问过的元素
	visited := make([]bool, len(nums))

	var backtracking func([]int)
	backtracking = func(path []int) {
		// terminator，长度等于n才是全排列
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		// 做选择、处理、撤销
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			path = append(path, nums[i])
			backtracking(path)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	backtracking([]int{})
	return res
}

// @lc code=end
