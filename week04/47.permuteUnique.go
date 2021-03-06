package week04

import "sort"

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
// 示例 1：
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
//  [1,2,1],
//  [2,1,1]
// ]
// @leetcode https://leetcode-cn.com/problems/permutations-ii

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

// 回溯+剪枝，套用模板
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	sort.Ints(nums)

	// 标记已经访问过的元素，位置
	visited := make([]bool, len(nums))
	// 回溯+剪枝
	var backtrack func([]int)
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...)) // 注意：需要拷贝
			return
		}
		// 选择、处理逻辑、回撤选择
		for i := 0; i < len(nums); i++ {
			// 元素已经添加过，跳过
			if visited[i] {
				continue
			}
			// 当前值等于前一个值： 两种情况：
			// 1. nums[i-1] 没用过 说明回溯到了同一层 此时接着用num[i] 则会与 同层用num[i-1] 重复
			// 2. nums[i-1] 用过了 说明此时在num[i-1]的下一层 相等不会重复
			if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			visited[i] = true
			path = append(path, nums[i])
			// drill down
			backtrack(path)
			// revert states
			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{})
	return res
}
