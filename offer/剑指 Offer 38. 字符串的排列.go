package offer

import "sort"

// 剑指 Offer 38. 字符串的排列
// 输入一个字符串，打印出该字符串中字符的所有排列。
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
// 示例:
// 输入：s = "abc"
// 输出：["abc","acb","bac","bca","cab","cba"]
// @lc: https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/

// dfs+回溯，类似数组全排列，[46.全排列](https://leetcode-cn.com/problems/permutations)
// 子集、排列 等问题首先要想到 dfs+回溯
func permutation(s string) []string {
	res, n := []string{}, len(s)
	if n == 0 {
		return res
	}

	// 先排序，便于去重
	sb := []byte(s)
	sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
	s = string(sb)

	// 标记已访问过的字符
	visited := make([]bool, n)
	var backtracking func(string)
	backtracking = func(path string) {
		// terminator，长度等于n的才是全排列字串
		if len(path) == n {
			res = append(res, path)
			return
		}
		// 选择，处理，回撤
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			// 当前字符等于前一个字符： 有两种情况：
			// 1. s[i-1] 没用过 说明回溯到了同一层 此时接着用num[i] 则会与 同层用num[i-1] 重复
			// 2. s[i-1] 用过了 说明此时在num[i-1]的下一层 相等不会重复
			if i > 0 && s[i] == s[i-1] && !visited[i-1] {
				continue
			}

			visited[i] = true    // mark
			path += string(s[i]) // process
			backtracking(path)   // drill down
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking("")
	return res
}
