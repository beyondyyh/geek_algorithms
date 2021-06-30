package practice

import "container/list"

/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
// 将字母放入栈中，输出的时候就是反序了，然后遍历S遇到字母就弹栈，非字母直接append
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func reverseOnlyLetters(s string) string {
	isLetter := func(char byte) bool {
		return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
	}

	letters := list.New() // stack
	for _, char := range []byte(s) {
		if isLetter(char) {
			letters.PushBack(char)
		}
	}

	res := make([]byte, 0, len(s))
	for _, char := range []byte(s) {
		if isLetter(char) {
			res = append(res, letters.Back().Value.(byte)) // 弹栈
			letters.Remove(letters.Back())
		} else {
			res = append(res, char)
		}
	}
	return string(res)
}

// 维护一个指针 j 从后往前遍历字符串，当需要字母时就使用它。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func reverseOnlyLetters2(s string) string {
	isLetter := func(char byte) bool {
		return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
	}

	n, res := len(s), make([]byte, 0, len(s))
	i, j, sb := 0, n-1, []byte(s)
	for ; i < n; i++ {
		if isLetter(sb[i]) {
			for !isLetter(sb[j]) {
				j--
			}
			res = append(res, sb[j])
			j--
		} else {
			res = append(res, sb[i])
		}
	}
	return string(res)
}

// @lc code=end
