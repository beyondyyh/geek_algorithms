package offer

// 剑指 Offer 39. 数组中出现次数超过一半的数字
// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// 示例 1:
// 输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
// 输出: 2
// @lc: https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/

// 方法一：hashmap
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func majorityElement(nums []int) int {
	m, n := map[int]int{}, len(nums)
	for _, num := range nums {
		m[num]++
		if m[num] > n/2 {
			return num
		}
	}
	return -1
}

// 方法二：比拼消耗
// 思路：遍历数组，从第一个元素开始计数，计数器初始为 1，遇到与当前元素相同的元素时，计数器加1，否则减 1（比拼消耗）。
// 当计数器为0时，重新从当前元素开始计数，重复上面步骤直至到达数组末尾。
func majorityElement2(nums []int) int {
	res, cnt := -1, 0
	for _, num := range nums {
		if cnt == 0 {
			res = num
			cnt = 1
		} else {
			if num == res {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return res
}
