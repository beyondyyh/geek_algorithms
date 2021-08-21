package offer

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

// 示例：
// 输入：nums = [1,2,3,4]
// 输出：[1,3,2,4]
// 注：[3,1,2,4] 也是正确的答案之一。
// @lc: https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

// 双指针left，right，分别从两端夹逼
// left如果为奇数一直往右移，right如果为偶数左移，知道二者相遇
// 否则交换left和right .
// 重复上述操作，直到 left == right 退出
// 时间复杂度：O(n)
func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && (nums[left]&1 == 1) { // left为奇数，一直右移
			left++
		}
		for left < right && (nums[right]&1 == 0) { // right为偶数，一直左移
			right--
		}
		nums[left], nums[right] = nums[right], nums[left] // 否则交换
		// 交换之后left必定是奇数，right必定是偶数，多移动一次减少遍历
		left++
		right--
	}
	return nums
}
