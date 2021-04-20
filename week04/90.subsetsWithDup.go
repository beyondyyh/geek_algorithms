package week04

import (
	"sort"
)

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 示例 1：
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
// @leetcode: https://leetcode-cn.com/problems/subsets-ii

// ------------------------------------------------------------------------------------------------------------------------
// 回溯通用模板：
// result = []
// func backtrack(选择列表, 路径) {
//     if 满足结束条件:
//         result.add(路径)
//         return
//     for 选择 in 选择列表:
//         做选择
//         backtrack(选择列表, 路径)
//         撤销选择，这步贼拉重要
// }
// ------------------------------------------------------------------------------------------------------------------------

// 回溯，与子集类似，唯一就是多了重复元素，那么：
// 做选择之前就需要去重，怎么去重呢？可以将数组排序，遍历选择元素时如果再遇到重复元素，则跳过
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	sort.Ints(nums)

	// 回溯思想
	// start 下次添加到集合中的元素位置索引
	// path  临时结果集合(每次需要复制保存)
	var backtrack func(int, []int)
	backtrack = func(pos int, path []int) {
		// 将子结果添加到结果集里
		res = append(res, append([]int{}, path...))
		// 选择+剪枝、处理结果、再撤销选择
		for i := pos; i < len(nums); i++ {
			// 排序之后，如果再遇到重复元素，则不选择此元素
			if i != pos && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{})
	return res
}
