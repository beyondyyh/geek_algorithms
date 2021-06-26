package practice

import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}

	sort.Ints(nums)
	res := [][]int{}
	var a, b, c, d int
	for a = 0; a <= n-4; a++ {
		if a > 0 && nums[a] == nums[a-1] { // a去重
			continue
		}
		for b = a + 1; b <= n-3; b++ {
			if b > a+1 && nums[b] == nums[b-1] { // b去重
				continue
			}
			c, d = b+1, n-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum < target {
					c++
				} else if sum > target {
					d--
				} else {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					for c < d && nums[c+1] == nums[c] { // c去重，后移一位
						c++
					}
					for c < d && nums[d] == nums[d-1] { // d去重，前移一位
						d--
					}
					// 正常逻辑，c和d继续两端夹逼，直到相遇后，进入下一轮
					c++
					d--
				}
			}
		}
	}
	return res
}

// @lc code=end
