package daily

// 32. 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
// 示例 1：
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
// 示例 2：
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
// @lc: https://leetcode-cn.com/problems/longest-valid-parentheses/

// 方法一：暴力法，先遍历字符串s得的字串str，依次判断字串是否为有效的括号。
// 时间复杂度：O(n^3) 遍历得的所有字串时间复杂度为O(n^2)，验证一个字串是否为有效时间复杂度为O(n)
// 空间复杂度：O(n)
func longestValidParentheses1(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	var isValid func(string) bool
	isValid = func(str string) bool {
		stack := &Stack{}
		for i := 0; i < len(str); i++ {
			if str[i] == '(' {
				stack.Push(str[i])
			} else if !stack.IsEmpty() && stack.Top().(byte) == '(' {
				stack.Pop()
			} else {
				return false
			}
		}
		return stack.IsEmpty()
	}

	end := n - 1
	if n&1 == 0 {
		end = n
	}
	for i := end; i >= 0; i -= 2 {
		for j := 0; j < n-i+1; j++ {
			// fmt.Printf("i:%d j:%d str:%s\n", i, j, s[j:j+i])
			if isValid(s[j : j+i]) {
				return i
			}
		}
	}
	return 0
}

// 方法二：栈
// 对于遇到的每个 '(' ，将它的下标放入栈中
// 对于遇到的每个 ')' ，我们先弹出栈顶元素表示匹配了当前右括号：
// 	如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的「最后一个没有被匹配的右括号的下标」
// 	如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func longestValidParentheses2(s string) int {
	res := 0
	stack := &Stack{}
	stack.Push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				res = max(res, i-stack.Top().(int))
			}
		}
	}
	return res
}

// 方法三：左右遍历
func longestValidParentheses3(s string) int {
	left, right, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return res
}
