package week04

// 给定一个没有重复数字的序列，返回其所有可能的全排列。
// 示例:
// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]
// @leetcode: https://leetcode-cn.com/problems/permutations

// ------------------------------------------------------------------------------------------------------------------------
// 回溯通用模板
// func backtrack(选择的起始位置, 路径) {
// 		if 满足结束条件:
// 			result.add(路径)
// 			return
// 		for i := 起始位置; i < 列表长度; i++ {
// 			做选择，路径.append(列表[i])
// 			backtrack(i+1, 路径)
// 			撤销选择
// 		}
// }
// ------------------------------------------------------------------------------------------------------------------------

// 回溯，子集、组合类题目首先想到用回溯，然后套用模板
func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	// 标记已经访问过的元素
	visited := make([]bool, len(nums))

	// backtrack 回溯思想
	// - path 每步选择之后的中间结果
	var backtrack func([]int)
	backtrack = func(path []int) {
		// 只有长度相等时，才是全排列
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...)) // 注意：需要拷贝
		}
		// 选择+标记，处理结果、再撤销选择+撤销标记
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			path = append(path, nums[i])
			backtrack(path)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{})
	return res
}
