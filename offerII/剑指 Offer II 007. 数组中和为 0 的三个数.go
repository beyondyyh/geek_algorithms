package offerII

import "sort"

// 剑指 Offer II 007. 数组中和为 0 的三个数
// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？请找出所有和为 0 且 不重复 的三元组。
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 示例 2：
// 输入：nums = []
// 输出：[]
// 示例 3：
// 输入：nums = [0]
// 输出：[]

// 双指针两端夹逼
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	// 先排个序，把相同的数字放在一起，便于去重撒
	sort.Ints(nums) // O(nlogn)
	res := [][]int{}
	// O(n^2)
	var i, j, k int
	for k = 0; k <= n-3; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] { // a去重
			continue
		}
		i, j = k+1, n-1
		for i < j { // 两端夹逼
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
				for i < j && nums[j] == nums[j-1] { // c去重
					j--
				}
				i++
				j--
			}
		}
	}
	return res
}
