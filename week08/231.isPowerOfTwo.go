package week08

// 231. 2的幂
// 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
// 示例 1:
// 输入: 1
// 输出: true
// 解释: 20 = 1
// @lc: https://leetcode-cn.com/problems/power-of-two/

// 首先2的幂为x，一定是x的二进制位中有且仅有一个1
// n&(n-1) 表示消除n的二进制最低位的1，由于2的幂二进制位有且仅有一个1，那么消除之后就会变为0了
// 时间复杂度：O(1)
// 空间复杂度：O(1)
func isPowerOfTwo(n int) bool {
	return (n > 0) && n&(n-1) == 0
}
