package practice

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start

// 方法一：排序+双指针法 两端夹逼，此题最优解
// abc三个指针，a->0..n-3, b->a+1..c, c->n-1..b
// a先保持不动，b和c两端加逼，直到相遇之后，再往后移动a，以此类推直到数组尾部
// 时间复杂度：O(n^2)
// 空间复杂度：O(log(n)) 快排的空间复杂度
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := [][]int{}
	var k, i, j int
	for k = 0; k <= n-3; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] { // a去重
			continue
		}
		i, j = k+1, n-1
		for i < j { // i和j两端夹逼
			sum := nums[k] + nums[i] + nums[j]
			if sum < 0 {
				i++
			} else if sum > 0 {
				j--
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] { // b去重
					i++
				}
				for i < j && nums[j] == nums[j-1] { // b去重
					j--
				}
				i++
				j--
			}
		}
	}
	return res
}

// @lc code=end
