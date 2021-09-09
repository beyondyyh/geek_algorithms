package offer

import (
	"math"
	"strconv"
)

// 剑指 Offer 17. 打印从1到最大的n位数
// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
// 示例 1:
// 输入: n = 1
// 输出: [1,2,3,4,5,6,7,8,9]
// @lc: https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/

// 0-9的n位全排列
// 整体思路：
// 1. 为了避免数字开头出现0，先把首位first固定，first取值范围为1~9
// 2. 用digit表示要生成的数字的位数，本题要从1位数一直生成到n位数，对每种数字的位数都生成一下首位，所以有个双重for循环
// 3. 生成首位之后进入递归生成剩下的 digit-1 位数，从0~9中取值
// 4. 递归的终止条件为已经生成了digit位的数字，即index == digit，将此时的数num转为int加到结果res中
// 时间复杂度：O(10^n)，从 1 到 10^n-1 都遍历了一遍
// 空间复杂度：O(n)，递归栈的最大深度
func printNumbers(n int) []int {
	count := 0 // 第count个数
	res := make([]int, int(math.Pow10(n)-1))

	// 递归生成剩下的 digit-1 位数，从0~9中取值
	var dfs func(int, []byte, int)
	dfs = func(index int, num []byte, digit int) {
		if index == digit {
			res[count], _ = strconv.Atoi(string(num))
			count++
			return
		}
		for i := '0'; i <= '9'; i++ {
			num[index] = byte(i)
			dfs(index+1, num, digit+1)
		}
	}

	for digit := 1; digit < n+1; digit++ {
		for first := '1'; first <= '9'; first++ {
			num := make([]byte, digit)
			num[0] = byte(first)
			dfs(1, num, digit)
		}
	}
	return res
}
