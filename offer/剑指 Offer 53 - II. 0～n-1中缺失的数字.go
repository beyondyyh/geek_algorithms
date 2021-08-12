package offer

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
// 示例 1:
// 输入: [0,1,3]
// 输出: 2
// 示例 2:
// 输入: [0,1,2,3,4,5,6,7,9]
// 输出: 8
// @lc: https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

// 方法一：位运算，异或
// x^x= 0, x^0=x, a^b^c=a^c^b
// 时间复杂度：O(n)
func missingNumber(nums []int) int {
	res, n := 0, len(nums)
	for i := 0; i < n; i++ {
		res ^= nums[i] ^ i
	}
	return res ^ n
}

// 方法二：有序数组找某个数，首先想到二分查找
// 左子数组：nums[i] = i；
// 右子数组：num[i] != i;
// 缺失的数字等于 “右子数组的首位元素” 对应的索引
// 时间复杂度：O(logn)
func missingNumber2(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
