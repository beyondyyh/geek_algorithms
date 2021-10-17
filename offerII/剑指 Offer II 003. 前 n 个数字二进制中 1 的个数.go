package offerII

// 剑指 Offer II 003. 前 n 个数字二进制中 1 的个数
// 给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。

// 示例 1:
// 输入: n = 2
// 输出: [0,1,1]
// 解释:
// 0 --> 0
// 1 --> 1
// 2 --> 10

// 示例 2:
// 输入: n = 5
// 输出: [0,1,1,2,1,2]
// 解释:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101

// 位运算：n&(n-1) = 0，去掉n的二进制最低位
func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}
