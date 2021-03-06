package offerII

// 剑指 Offer II 004. 只出现一次的数字
// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
// 示例 1：
// 输入：nums = [2,2,3,2]
// 输出：3
// 示例 2：
// 输入：nums = [0,1,0,1,0,1,100]
// 输出：100

// 提示：
// 1 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次

// 解题思路：位运算
// 位运算： x&1：得到x最低位值； x&(~x) = 0：x与上x取反，结果为0；x&(x-1)：去掉x最低位的1
// 思路：x^1 可获取最低位的值(0或者1)，再配合左移操作，可以获得 x 的从低到高所有二进制位上的值；
// 1. 建立一个长度为 32 的数组 counts，通过以上方法（x&1, x >>= 1）可计算出所有数字的各二进制位的 1 的出现次数。
// 2. 遍历将 counts 各元素对 3 求余，则结果会得到 “只出现一次的数字Y” 的各二进制位。
// 3. 利用 左移操作 和 或运算 ，可将 counts 数组中各二进位的值恢复到数字 res 上

// 时间复杂度：O(n)，其中n为数组 nums 的长度；遍历数组占用 O(n) ，每轮中的常数个位运算操作占用O(1) 。
// 空间复杂度：O(1)，数组 counts 长度恒为 32，占用常数大小的额外空间
func singleNumber(nums []int) int {
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
		// counts[i] %= 3 // 得到 只出现一次的数字 的第 (31 - i) 位
		res <<= 1
		res |= counts[31-i] % m
	}
	return res
}
