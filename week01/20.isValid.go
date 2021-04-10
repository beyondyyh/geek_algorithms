package week01

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。左括号必须以正确的顺序闭合。
// 示例 1：
// 输入：s = "()[]{}"
// 输出：true
// @leetcode: https://leetcode-cn.com/problems/valid-parentheses

// 最近相似性，标准的用栈
// 方法一：遍历字符串将左括号入栈，当遇到右括号时出栈栈顶元素并判断是否相等，不相等肯定不匹配，同理可得，最后判断栈是否为空
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid1(s string) bool {
	dict := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := &Stack{}
	for _, c := range []byte(s) {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
		} else {
			// 以下三种情况均为无效：
			// 1. 当前元素不是有效的右括号
			// 2. 当前元素是右括号但是栈为空，即没有与之相匹配的左括号
			// 3. 弹出的栈顶元素不能与当前元素进行匹配
			val, exists := dict[c]
			if !exists || stack.IsEmpty() || stack.Pop().(byte) != val {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

// 与方法一类似，只不过把右括号入栈
func isValid2(s string) bool {
	stack := &Stack{}
	for _, c := range s {
		switch {
		case c == '(':
			stack.Push(')')
		case c == '[':
			stack.Push(']')
		case c == '{':
			stack.Push('}')
		default: // 如果栈为空则说明前面都没push成功，即有无效的字符
			if stack.IsEmpty() || c != (stack.Pop()).(rune) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
