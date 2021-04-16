package week03

import "fmt"

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 示例 1：
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// @leetcode: https://leetcode-cn.com/problems/generate-parentheses

// // 泛型递归模板四部曲
// func recursion(level int) {
// 	// 1. terminator 终止条件
// 	if level <= MAX {
// 		return
// 	}
// 	// 2. 处理当前层逻辑
// 	process current logic
// 	// 3. drill down 下探到下一层
// 	recursion(level+1)
// 	// 4. reverse states, 清除当前层状态，if needed
// }

// 方法一：递归、剪枝
func generateParenthesis(n int) []string {
	var res []string
	generate(0, 0, n, "", &res)
	return res
}

// generate
// left: 	已经使用的左括号的个数
// right: 	已经使用的右括号的个数
// n: 		代表生成括号的对数，题目的n
// s: 		当前拼成的字符串
// res: 	保存的结果集
func generate(left, right, n int, s string, res *[]string) {
	// terminator 终止条件
	// 已经使用的左括号个数=n，且右括号个数为也为n
	if left == n && right == n {
		*res = append(*res, s)
		fmt.Printf("left:%d right:%d s:%s res:%+v\n", left, right, s, res)
		return
	}

	// process current logic && drill down
	// 1. 什么时候可以加左括号？ 左括号随后可以加，只要不超过n即可
	// 2. 什么时候可以加右括号？ 前面有左括号的情况，也就是说左括号使用的个数大于右括号的情况，且右括号使用个数小于n，
	// 但是由于n > left > right，所以 right < n一定成立，就不用写冗余代码了
	//
	// 也就是著名的 “剪枝” 思想
	if left < n {
		// drill down，下探到下一层逻辑
		generate(left+1, right, n, s+"(", res)
	}
	if left > right {
		// drill down，下探到下一层逻辑
		generate(left, right+1, n, s+")", res)
	}

	// revese states
	// nothing to do
}
