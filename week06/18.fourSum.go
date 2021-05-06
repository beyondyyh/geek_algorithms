package week06

import (
	"sort"
)

// 18. 四数之和
// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
// 注意：答案中不可以包含重复的四元组。
// 示例 1：
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// @lc: https://leetcode-cn.com/problems/4sum/

//---------------------------------------------------------------
// 来来来，再复习一下“回溯”通用模板
// func backtrack(选择的起始位置pos, 路径) {
// 		if 满足结束条件:
// 			result.add(路径)
// 			return（return一定不能搞丢了）
// 		for i := 起始位置pos; i < 列表长度; i++ {
// 			做选择，路径.append(列表[i])
// 			backtrack(i+1, 路径)
// 			撤销选择
// 		}
// }
//---------------------------------------------------------------

// 方法一：回溯
// 时间复杂度：指数级，因为有剪枝不太好评估
// 空间复杂度：O(log(n))
func fourSum1(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}

	// 排个序
	sort.Ints(nums)

	res := [][]int{}
	var backtracking func(int, int, []int)
	backtracking = func(index, diff int, path []int) {
		// fmt.Printf("index:%d diff:%d path:%v\n", index, diff, path)
		// terminator
		if diff == 0 && len(path) == 4 {
			res = append(res, append([]int{}, path...))
			return
		}
		// 选择、处理、回撤
		for i := index; i < n; i++ {
			// 剪枝1：如果剩余可选的数字数量小于 (4-path.size())，说明剩下的凑不够4个了则剪掉，直接返回
			if n-i < (4 - len(path)) {
				return
			}
			// 剪枝2：从第二轮(i>index)循环开始，如果当前数字和前一个相同，则进入下次遍历，去重逻辑
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			// leftNum 表示选择当前数字后，最大还能选几个
			leftNum := 3 - len(path)
			// 剪枝3：如果当前数字加上下一个数字乘以leftNum 大于 diff，由于nums升序，后面的数无论怎么选都会大于diff，也即当前数字太大了而且后续会更大，因此可以直接返回
			if i < n-1 && nums[i]+nums[i+1]*leftNum > diff {
				return
			}
			// 剪枝4：如果当前数字加上nums最后一个数字（最大）小于 diff，由于nums升序，后面的数无论怎么选都会小于diff，也即当前选择的数字太小了，需要进入下一层循环
			// 注意：这里跟剪枝3略有不同，由于nums是升序的，如果当前层数字都太小了，可以继续进入下一层循环(i+1)继续搜索，因此是continue，而不是return
			if i < n-1 && nums[i]+nums[n-1]*leftNum < diff {
				continue
			}
			// drill down & revert
			path = append(path, nums[i])
			backtracking(i+1, diff-nums[i], path)
			path = path[:len(path)-1]
		}
	}

	backtracking(0, target, []int{})
	return res
}

// 方法二：排序+双指针法，左右夹逼
// 类似于三数之和，外层多一次遍历
// 时间复杂度：O(n^3)
// 空间复杂度：O(log(n)) 快排引入的额外存储
func fourSum2(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}

	sort.Ints(nums)
	res := [][]int{}
	// abcd四个指针，a->0..n-4, b->a+1..n-3, c->b+1..n-2 d->n-1..c，c和d两端夹逼
	// 首先c和d两端夹逼，当相遇之后，移动b，直到b没法移动之后，再移动a
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
				// fmt.Printf("a:%d b:%d c:%d d:%d sum:%d\n", a, b, c, d, sum)
				if sum < target {
					c++
				} else if sum > target {
					d--
				} else {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					for c < d && nums[c+1] == nums[c] { // c去重
						c++
					}
					for c < d && nums[d-1] == nums[d] { // d去重
						d--
					}
					// 继续两端夹逼，直到相遇后，进入下一轮
					c++
					d--
				}
			}
		}
	}
	return res
}
