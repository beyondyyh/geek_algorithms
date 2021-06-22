package practice

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start

// // --- stack 基本实现
// type Stack []interface{}

// func (s *Stack) Len() int           { return len(*s) }
// func (s *Stack) IsEmpty() bool      { return s.Len() == 0 }
// func (s *Stack) Push(x interface{}) { *s = append(*s, x) }                                    // 将x入栈，size++
// func (s *Stack) Pop() interface{}   { x := (*s)[s.Len()-1]; *s = (*s)[:s.Len()-1]; return x } // 弹出栈顶元素，size--
// func (s *Stack) Peek() interface{}  { return (*s)[s.Len()-1] }                                // 查看栈顶元素，size不变

// 方法一：具有最近相似性，用stack
// 遍历字符串s，遇到左括号时：将对应的右括号入栈，遇到右括号时：如果栈为空 或 栈顶元素不等于当前右括号，说明不匹配可以提前返回false
// 最后看栈是否为空，为空说明都匹配上了
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid(s string) bool {
	// 奇数长度，肯定不匹配
	if len(s)&1 == 1 { // n%2==1
		return false
	}

	dict := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := &Stack{}
	for _, char := range []byte(s) {
		switch char {
		case '(', '[', '{': // 遇到左括号时：将对应的右括号入栈
			stack.Push(dict[char])
		default: // 遇到右括号：此时栈为空(说明前面没有左括号) 或 栈顶元素与当前右括号不相等，均无效，直接返回fase
			if stack.IsEmpty() || stack.Pop().(byte) != char {
				return false
			}
		}

	}
	// 最后看栈是否为空
	return stack.IsEmpty()
}

// @lc code=end
