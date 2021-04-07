package week01

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 进阶：
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
// 示例 :
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右旋转 1 步: [7,1,2,3,4,5,6]
// 向右旋转 2 步: [6,7,1,2,3,4,5]
// 向右旋转 3 步: [5,6,7,1,2,3,4]
// @leetcode：https://leetcode-cn.com/problems/rotate-array

// 方法一：多次翻转，也是最容易理解的方法，先将整个数组翻转，然后在再分别翻转前 k 个后面部分
// 时间复杂度：O(n)，n/2, k/2, n-k/2，n为数组长度
// 空间复杂度：O(1)，没有引入额外的存储
func rotate1(nums []int, k int) {
	k %= len(nums)    // 这个小心思比较巧妙，可以防止当k=len(nums)时数组越界的情况
	reverse(nums)     // [7,6,5,4,3,2,1]
	reverse(nums[:k]) // [5,6,7,4,3,2,1]
	reverse(nums[k:]) // [5,6,7,1,2,3,4]
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

// 方法二：使用额外的数组
// 遍历nums，将 num[i] 赋值给newNums的 (i+k) % len(nums) 最终的位置
// 时间复杂度：O(n)，n为数组的长度
// 空间复杂度：O(n)
func rotate2(nums []int, k int) {
	n := len(nums)
	newNums := make([]int, n)
	for i := 0; i < n; i++ {
		newNums[(i+k)%n] = nums[i] // 取模防止数组越界
	}
	copy(nums, newNums) // copy(dst, src []Type) int
}
