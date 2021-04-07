package week01

// 给你一个有序数组nums，请你 原地 删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// @Leetcode: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array

// 方法一：对于有序的数组，双指针法（也叫快慢指针法）遍历能解决大部分case，应该先往双指针上面靠
// 题解：慢指针i，快指针j，i,j:=0,1 快指针j先走，
// 如果nums[j] = nums[i]说明是重复的项，i不动，j继续往前走，直到nums[j] != nums[i]说明重复运动结束，此时i走一步，同时把j的值赋值给i
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	i := 0
	for j := 1; j < n; j++ { // 快指针走到数组结尾则遍历结束
		if nums[j] != nums[i] { // 相等时说明是重复项，不需要处理，快指针继续走，直到不等才需要赋值
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
