package offer

// 剑指 Offer 31. 栈的压入、弹出序列
// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

// 示例 1：
// 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// 输出：true
// 解释：我们可以按以下顺序执行：
// push(1), push(2), push(3), push(4), pop() -> 4,
// push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

// 示例 2：
// 输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
// 输出：false
// 解释：1 不能在 2 之前弹出。
// @lc: https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

// 借助一个辅助栈，将pushed元素入栈，当栈顶元素等于poped[i]时，执行pop，最后看栈是否为空
// 题目说 pushed 是 popped 的排列，所以不用判断长度是否相等的情况
// 时间复杂度：O(n) pushed中每个元素最多入栈和出栈一次
// 空间复杂度：O(n)
func validateStackSequences(pushed []int, popped []int) bool {
	stack, i := []int{}, 0
	for _, num := range pushed {
		stack = append(stack, num) // Pop
		// !stack.IsEmpty() && stack.Peek() == poped[]i
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}
