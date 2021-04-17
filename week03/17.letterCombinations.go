package week03

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 示例 1：
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// @leetcode: https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number

// 方法一：傻傻的递归，类似于括号生成(https://leetcode-cn.com/problems/generate-parentheses/)
func letterCombinations1(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}

	// 数字->字符的映射
	dict := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// i: 输入的digits的下标，当前遍历到第几层
	// s: 当前组成的字符串，临时变量
	var combinations func(int, string)
	combinations = func(i int, s string) {
		// 1. terminator 终止条件
		if i == len(digits) {
			res = append(res, s)
			return
		}
		// 2. process current logic
		letters := dict[rune(digits[i])] // 当前数字对应字母
		for j := 0; j < len(letters); j++ {
			// 3. drill down
			combinations(i+1, s+string(letters[j]))
		}
		// 4. revert, nothing todo
	}
	combinations(0, "")
	return res
}
