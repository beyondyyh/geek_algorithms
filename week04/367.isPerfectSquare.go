package week04

// 367. 有效的完全平方数
// 给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
// 进阶：不要 使用任何内置的库函数，如  sqrt 。
// 示例 1：
// 输入：num = 16
// 输出：true
// 示例 2：
// 输入：num = 14
// 输出：false

// 二分查找法
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	left, right := 2, num/2
	for left <= right {
		mid := left + (right-left)/2
		guess := mid * mid
		if guess == num {
			return true
		} else if guess > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
