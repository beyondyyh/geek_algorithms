package offer

// 剑指 Offer 56 - II. 数组中数字出现的次数 II
// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

// 示例 1：
// 输入：nums = [3,4,3,3]
// 输出：4

// 示例 2：
// 输入：nums = [9,1,7,9,7,9,7]
// 输出：1
// @lc: https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/

// 解题思路：位运算
// x&1：得到x最低位值； x&(~x) = 0：x与上x取反，结果为0；x&(x-1)：去掉x最低位的1
// 1. x^1 可获取最低位的值(0或者1)，再配合左移操作，可以获得 x 的从低到高所有二进制位上的值；
// 2. 建立一个长度为 32 的数组 counts，通过以上方法（x&1, x >>= 1）可计算出所有数字的各二进制位的 1 的出现次数。
// 3. 遍历将 counts 各元素对 3 求余，则结果会得到 “只出现一次的数字Y” 的各二进制位。
// 4. 利用 左移操作 和 或运算 ，可将 countscounts 数组中各二进位的值恢复到数字 res 上

// 时间复杂度：O(n)，其中n为数组 nums 的长度；遍历数组占用 O(n) ，每轮中的常数个位运算操作占用O(1) 。
// 空间复杂度：O(1)，数组 counts 长度恒为 32，占用常数大小的额外空间
func singleNumberII(nums []int) int {
	counts := [32]int{}
	// 1.统计所有位上的1个数总和
	for _, num := range nums {
		for j := 0; j < 32; j++ {
			counts[j] += num & 1 // 统计当前位上1的个数，num&1会得到num最低位的1
			num >>= 1            // 去掉最低位
		}
	}

	// 2.恢复res
	res, m := 0, 3
	for i := 0; i < 32; i++ {
		res <<= 1
		res |= counts[31-i] % m
	}
	return res
}
