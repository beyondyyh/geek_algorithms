package week06

import "sort"

// 15. 三数之和
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// @lc: https://leetcode-cn.com/problems/3sum/

// 方法一：排序+双指针法 两端夹逼，此题最优解
// abc三个指针，a->0..n-3, b->a+1..c, c->n-1..b
// a先保持不动，b和c两端加逼，直到相遇之后，再往后移动a，以此类推直到数组尾部
// 时间复杂度：O(n^2)
// 空间复杂度：O(log(n)) 快排的空间复杂度
func threeSum1(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := [][]int{}
	var a, b, c int
	for a = 0; a <= n-3; a++ {
		if nums[a] > 0 {
			break
		}
		if a > 0 && nums[a] == nums[a-1] { // a去重
			continue
		}
		b, c = a+1, n-1
		for b < c { // b和c两端夹逼
			sum := nums[a] + nums[b] + nums[c]
			if sum < 0 {
				b++
			} else if sum > 0 {
				c--
			} else {
				res = append(res, []int{nums[a], nums[b], nums[c]})
				for b < c && nums[b] == nums[b+1] { // b去重
					b++
				}
				for b < c && nums[c] == nums[c-1] { // b去重
					c--
				}
				b++
				c--
			}
		}
	}
	return res
}

// 回溯+剪枝，数量级大时可能超时
// 通用解法，不论是4sum，5sum，甚至nsum都可用此解法
// 时间复杂度：由于加了剪枝逻辑，所以不会评估
func threeSum2(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	// 先排序
	sort.Ints(nums)

	res := [][]int{}
	var backtracking func(int, int, []int)
	backtracking = func(pos, diff int, path []int) {
		// fmt.Printf("pos:%d diff:%d path:%v\n", pos, diff, path)
		// terminator
		if diff == 0 && len(path) == 3 {
			res = append(res, append([]int{}, path...))
			return
		}
		// 选择、处理、回撤
		for i := pos; i < n; i++ {
			// 剪枝1：如果剩下元素个数不足以凑够3个，则直接return
			if n-i < 3-len(path) {
				return
			}
			// 剪枝2：从第二轮(i>pos)循环开始，如果当前数字和前一个相同，则进入下次遍历，去重逻辑
			if i > pos && nums[i] == nums[i-1] {
				continue
			}
			// leftNum 表示选择当前数字后，最大还能选几个
			leftNum := 2 - len(path)
			// 剪枝3：如果当前数字加上下一个数乘以leftNum 大于 diff，由于nums递增则说明后续不管怎么选择都会大于diff，则直接return
			if i < n-1 && nums[i]+nums[i+1]*leftNum > diff {
				return
			}
			// 剪枝4：如果当前数字加上nums最后一个数乘以leftNum 小于 diff，由于nums递增则说明后续不管怎么选择都会小于diff，则继续下次循环
			if i < n-1 && nums[i]+nums[n-1]*leftNum < diff {
				continue
			}

			// 以上剪枝1..2是提前判断，并剪掉不合法的分支，并没有真正做选择
			// process current logic & drill down
			path = append(path, nums[i])
			backtracking(i+1, diff-nums[i], path)
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0, []int{})
	return res
}
