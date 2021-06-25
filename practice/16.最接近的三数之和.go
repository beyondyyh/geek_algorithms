package practice

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
// 双指针法，i->[0..n-3], left->[i+1..right], right->[n-1..i]
// 时间复杂度：O(nlogn) + O(n^2) = O(n^2)
// 空间复杂度：O(nlogn)
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			// 绝对值越小越接近target，更新res
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}

			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}
	}
	return res
}

// @lc code=end
