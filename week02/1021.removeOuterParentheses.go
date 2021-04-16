package week02

// 示例 1：
// 输入："(()())(())"
// 输出："()()()"
// 解释：
// 输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
// 删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。
// @leetcode: https://leetcode-cn.com/problems/remove-outermost-parentheses

// 时间复杂度：O(n)
func removeOuterParentheses(S string) string {
	var ans string
	level := 0
	for _, c := range S {
		if c == ')' {
			level--
		}
		if level >= 1 {
			ans += string(c)
		}
		if c == '(' {
			level++
		}
	}
	return ans
}
