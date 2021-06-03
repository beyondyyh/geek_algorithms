package daily

// 718. 最长重复子数组
// 给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
// 示例：
// 输入：
// A: [1,2,3,2,1]
// B: [3,2,1,4,7]
// 输出：3
// 解释：
// 长度最长的公共子数组是 [3, 2, 1] 。
// @lc: https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/

// 方法一：暴力枚举
// i->0..len(nums1)-1, j->0..len(nums2)-1, 只有
// 时间复杂度：O(n^3)
// 空间复杂度：O(n)
func findLength1(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
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

// 方法二：动态规划，bottom-up
// dp状态定义：dp[i][j] 表示 A数组从i位置到尾部（A[i:]） 和 B数组从j位置到尾部（B[j:]） 的最长公共前缀
// dp转移方程：
// 	A[i] == B[j]: dp[i][j] = dp[i+1][j+1]+1，那么 A[i:] 与 B[j:] 的最长公共前缀为 A[i+1:] 与 B[j+1:] 的最长公共前缀的长度加1
//	A[i] != B[j]: dp[i][j] = 0，则说明当前元素不相同，那么就没有公共前缀，即A[i:] 与 B[j:] 的最长公共前缀为0
// 由于dp[i][j] 的值从 dp[i+1][j+1]转移得到，所以需要倒过来，首先计算 dp[m-1][n-1]，最后计算 dp[0][0]
// 最终结果：dp数组中的最大值，可以在迭代的过程中直接保存最大值即可
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func findLength2(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)

	// dp初始化
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 递推
	res := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}
