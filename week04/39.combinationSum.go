package week04

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 说明：
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
// 示例 1：
// 输入：candidates = [2,3,6,7], target = 7,
// 所求解集为：
// [
//   [7],
//   [2,2,3]
// ]
// @leetcode: https://leetcode-cn.com/problems/combination-sum

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

// 回溯
// 可以套用`回溯通用模板`，问题在于如何去重？
// 由于每一个元素可以重复使用，下一轮搜索的起点依然是 i，不是 i+1
// 每一层携带到下层梦境的变量是 diff值
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}

	var backtrack func(int, int, []int)
	backtrack = func(pos int, diff int, path []int) {
		// 差值为0说明符合要求，追加结果，注意：复制保存
		if diff == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		// 选择、处理逻辑、回撤
		for i := pos; i < len(candidates); i++ {
			// 剪枝逻辑：如果下一个要添加的元素大于diff值，就不需要再继续了，直接进入下一层梦境
			if candidates[i] > diff {
				continue // 如果数组是升序的，那么其实可以break了，因为后续元素都比i大，更不合适了
			}
			path = append(path, candidates[i])
			// 注意：由于每一个元素可以重复使用，下一轮搜索的起点依然是 i，这里非常容易弄错
			// 注意：这里携带到下层梦境的变量是diff值
			backtrack(i, diff-candidates[i], path)
			// 回撤选择
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target, []int{})
	return res
}
