package week04

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
// 示例:
// 输入: n = 4, k = 2
// 输出:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]
// @leetcode: https://leetcode-cn.com/problems/combinations

// 方法一：递归、回溯
// 时间复杂度：回溯或递归的时间复杂度最坏的情况为：O(n^k)，但如果剪枝剪的好的话会大大降低时间复杂度，所以不必过于纠结
func combine(n, k int) [][]int {
	var res [][]int
	if k <= 0 || n < k {
		return res
	}

	var dfsCombine func(int, []int)
	dfsCombine = func(start int, path []int) {
		// 1. terminator 递归终止条件
		if len(path) == k {
			// res = append(res, path) // 错误的写法。。
			// 重点理解这里：向res结果集追加结果时最好copy一份path，不然会发生不可预知的错误
			res = append(res, append([]int{}, path...))
			return
		}
		// 2. process current logic
		// for i := start; i <= n; i++ { // 不剪枝
		for i := start; i <= n-(k-len(path))+1; i++ { // 剪枝
			// 向路径里添加一个数
			path = append(path, i)
			// 3. drill down，递推下探
			dfsCombine(i+1, path)
			// 4. revert states，也叫回溯
			// 重点理解这里：深度优先遍历有回头的过程，因此递归之前做了什么，递归之后需要做相同操作的逆向操作
			path = path[:len(path)-1]
		}
	}

	dfsCombine(1, []int{})
	return res
}
