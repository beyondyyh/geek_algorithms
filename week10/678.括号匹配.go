package week10

// 678. 有效的括号字符串
// 给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：

// 任何左括号 ( 必须有相应的右括号 )。
// 任何右括号 ) 必须有相应的左括号 ( 。
// 左括号 ( 必须在对应的右括号之前 )。
// * 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。
// 一个空字符串也被视为有效字符串。

// 栈，时间复杂度O(n)，空间复杂度O(n)
// 思路：
// 在有星号的情况下，需要两个栈分别存储左括号和星号。从左到右遍历字符串，进行如下操作:
// 1.如果遇到左括号，则将当前下标存入左括号栈。
// 2.如果遇到星号，则将当前下标存入星号栈。
// 3.如果遇到右括号，则需要有一个左括号或星号和右括号匹配，由于星号也可以看成右括号或者空字符串，因此当前的右括号应优先和左括号匹配，没有左括号时和星号匹配：
// 	a.如果左括号栈不为空，则从左括号栈弹出栈顶元素；
// 	b.如果左括号栈为空且星号栈不为空，则从星号栈弹出栈顶元素；
// 	c.如果左括号栈和星号栈都为空，则没有字符可以和当前的右括号匹配，返回 false

// 遍历结束之后，左括号栈和星号栈可能还有元素。为了将每个左括号匹配，需要将星号看成右括号，且每个左括号必须出现在其匹配的星号之前。
// 当两个栈都不为空时，每次从左括号栈和星号栈分别弹出栈顶元素，对应左括号下标和星号下标，判断是否可以匹配，匹配的条件是左括号下标小于星号下标，如果左括号下标大于星号下标则返回 false

// 最终判断左括号栈是否为空。如果左括号栈为空，则左括号全部匹配完毕，剩下的星号都可以看成空字符串，此时 ss 是有效的括号字符串，返回true。
// 如果左括号栈不为空，则还有左括号无法匹配，此时 ss 不是有效的括号字符串，返回 false
func checkValidString(s string) bool {
	var leftStk, starStk []int
	for i, ch := range s {
		if ch == '(' {
			leftStk = append(leftStk, i)
		} else if ch == '*' {
			starStk = append(starStk, i)
		} else if len(leftStk) > 0 {
			leftStk = leftStk[:len(leftStk)-1]
		} else if len(starStk) > 0 {
			starStk = starStk[:len(starStk)-1]
		} else {
			return false
		}
	}
	i := len(leftStk) - 1
	for j := len(starStk) - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if leftStk[i] > starStk[j] {
			return false
		}
	}
	return i < 0
}
