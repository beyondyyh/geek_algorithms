package daily

// 46. 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// @lc: https://leetcode-cn.com/problems/permutations/

// 方法一：dfs+回溯
func permute1(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	res := [][]int{}
	visited := make([]bool, n)
	// dfs+回溯
	// - path 每步选择之后的中间结果
	var backtracking func([]int)
	backtracking = func(path []int) {
		// terminator
		if len(path) == n {
			res = append(res, path) // 注意这里，path有时需要是深拷贝 res = append(res, append([]int{}, path...))
			return
		}
		// 选择、处理、回撤
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}

			visited[i] = true            // marking
			path = append(path, nums[i]) // choice
			backtracking(path)           // backtrack
			visited[i] = false           // revert
			path = path[:len(path)-1]    // revert
		}
	}
	backtracking([]int{})

	return res
}
