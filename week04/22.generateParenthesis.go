package week04

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 示例 1：
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// @leetcode: https://leetcode-cn.com/problems/generate-parentheses

// 泛型递归思想
func generateParenthesis(n int) []string {
	res := []string{}
	if n <= 0 {
		return res
	}

	var generate func(int, int, string)
	generate = func(left, right int, s string) {
		// terminator
		if n == left && right == n {
			res = append(res, s)
			return
		}
		// process current logic
		if left < n {
			// drill down
			generate(left+1, right, s+"(")
		}
		if right < left {
			generate(left, right+1, s+")")
		}
	}

	generate(0, 0, "")
	return res
}
