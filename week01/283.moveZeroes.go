package week01

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 示例:
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 说明:
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
// @leetcode https://leetcode-cn.com/problems/move-zeroes

// 方法一：双指针法，遍历数组会经常用到，特别是有序数组
// 思路：j保存非零元素下标，i遍历数组，当i位置!=0则赋值给j同时j后移一位，直到i遍历结束，最后把j<n的位置都补上0即可
// 时间复杂度：>O(n)
// 空间复杂度：O(1)
func moveZeroes1(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}

// 方法二：一次遍历，借助快排思想
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func moveZeroes2(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		// 当前元素!=0，就把其交换到左边，等于0的交换到右边
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
