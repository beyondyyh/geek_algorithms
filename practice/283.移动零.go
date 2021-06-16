package practice

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start

// 方法一：双指针
// i->0..n-1, j->0..，i指针先走不断+1，如果nums[i] != 0 则赋值给j，同时j++，
// i遍历到数组尾部之后，此时已经把非零数字全部挪到前面了，最后判断如果j<i，则手动把剩下的位补0
func moveZeroes1(nums []int) {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	// 最后判断是否需要补0
	for j < i {
		nums[j] = 0
		j++
	}
}

// 双指针大法2
// 跟上面一样，只不过当 nums[i] != 0 时，直接交换 j 和 i，为啥可以？
// 因为 i 走的快在前面，j 在后面，nums[i]赋值给nums[j]之后，j指针也往后这一步了，即使此时j 不为0那么j和i一定是重叠的，如果不重叠那么j上肯定是0
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}

// @lc code=end
