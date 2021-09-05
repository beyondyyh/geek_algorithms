package offer

// 剑指 Offer 57 - II. 和为s的连续正数序列
// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

// 示例 1：
// 输入：target = 9
// 输出：[[2,3,4],[4,5]]

// 示例 2：
// 输入：target = 15
// 输出：[[1,2,3,4,5],[4,5,6],[7,8]]
// @lc: https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/

// 方法一：滑动窗口
// 两个指针left、right分别指向窗口的左右边界
// 1. 如果窗口内的和 大于target，说明窗口大了，left右移一步；
// 2. 如果窗口内的和 小于target，说明窗口下了，right右移一步；
// 3. 如果窗口内的和 等于target，说明找到一组，放入res。
// 题目要求至少包含2个数，所以left <= target/2 即可，再往后肯定不满足2数之和等于target了
// 时间复杂度：O(n)，其中 n = target
// 空间复杂度：O(1)
func findContinuousSequence(target int) [][]int {
	res := [][]int{}
	left, right := 1, 2 // 初始窗口的左右边界
	sum := left + right // 窗口内数字的和
	for left <= target/2 {
		if sum < target {
			right++
			sum += right
		} else if sum > target {
			sum -= left // left右移，需要减去左边滑出窗口之外的数
			left++
		} else { // 找到一组合法数
			tmp := make([]int, 0, right-left+1)
			for i := left; i <= right; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp) // 将合法数加入结果集
			// left、right同时后移一位
			sum -= left
			left++
			right++
			sum += right
		}
	}
	return res
}
