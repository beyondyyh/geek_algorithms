package offer

import (
	"strconv"
	"strings"
)

// 剑指 Offer 46. 把数字翻译成字符串
// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

// 示例 1:
// 输入: 12258
// 输出: 5
// 解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
// @lc: https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

// dp状态定义：dp[i]表示num[0..i] 即前i个字符的翻译方法数
// dp转移方程：
//	dp[i] = dp[i-1] + dp[i-2] 前面两个数在10–25之间，可以单独翻译也可放在一起翻译
//	dp[i] = dp[i-1] 前面两个数不在1–25之间，只能翻译自己，方法数没有增加。
// dp初始化：dp[0] = dp[1] = 1
// 最终结果：dp[n]
func translateNum(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		tmp := str[i-2 : i] // 前闭后开区间
		if strings.Compare(tmp, "10") >= 0 && strings.Compare(tmp, "25") <= 0 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
