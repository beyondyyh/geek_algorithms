package practice

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
// 思路：将 map[nums[i]] = i 放入hashmap，遍历nums计算出diff，
// 看diff在不在map中。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if j, exists := m[target-num]; exists {
			return []int{i, j}
		}
		m[num] = i
	}
	return []int{}
}

// @lc code=end
