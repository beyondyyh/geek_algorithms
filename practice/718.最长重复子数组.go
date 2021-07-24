package practice

/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
// 方法一：暴力枚举
// i->0..len(nums1)-1, j->0..len(nums2)-1, 如果nums1[i] == nums2[j]则，长度至少为1，不断向后扩散直到不相等
// 时间复杂度：O(n^3)
// 空间复杂度：O(1)
func findLength1(nums1 []int, nums2 []int) int {
	var res int
	m, n := len(nums1), len(nums2)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0
			for i+k < m && j+k < n && nums1[i+k] == nums2[j+k] {
				k++
			}
			res = max(res, k)
		}
	}
	return res
}

// 方法二：动态规划
// dp状态定义：dp[i][j]表示长度为i，以nums1[i-1]结尾的子数组 和 长度为j，以nums2[j-1]结尾的子数组，二者的最大公共后缀子数组的长度
// dp转移方程：
// 	nums1[i-1] == nums2[j-1], dp[i][j] = dp[i-1][j-1]+1
//	nums1[i-1] != nums2[j-1], dp[i][j] = 0
// dp初始化：初始化为0
// 最终结果：dp数组中的最大值，可以在迭代的过程中直接保存最大值即可
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func findLength(nums1, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	var res int
	// i=0或j=0的 base case，初始化时已经包括
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			res = max(dp[i][j], res)
		}
	}
	return res
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
