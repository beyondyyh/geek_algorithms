package offer

import "sort"

// 剑指 Offer 61. 扑克牌中的顺子
// 从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

// 示例 1:
// 输入: [1,2,3,4,5]
// 输出: True

// 示例 2:
// 输入: [0,0,1,2,5]
// 输出: True
// @lc: https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/

// 方法一：一次遍历，排重 并 找到最大值和最小值，如果没有重复 且 max-min < 5则可以构成顺子
// 时间复杂度：O(n) = O(5) = O(1)
// 空间复杂度：O(n)
func isStraight(nums []int) bool {
	m := make(map[int]struct{})
	ma, mi := 0, 14
	for _, num := range nums {
		// 跳过大小王
		if num == 0 {
			continue
		}
		ma = max(ma, num)          // 最大牌
		mi = min(mi, num)          // 最小牌
		if _, has := m[num]; has { // 有重复牌，直接返回false
			return false
		}
		m[num] = struct{}{}
	}
	return ma-mi < 5 // 最大牌 - 最小牌 < 5 则可构成顺子
}

// 方法二：排序后，判重 和 统计大小王的个数，同时 max-min < 5则可以构成顺子
// 时间复杂度：O(nlog(n)) = O(5log(5)) = O(1)
// 空间复杂度：O(1)
func isStraight2(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 { // 统计大小王个数
			joker++
		} else if nums[i] == nums[i+1] { // 有重复牌
			return false
		}
	}
	return nums[4]-nums[joker] < 5 // 最大牌 - 最小牌 < 5 则可构成顺子
}
