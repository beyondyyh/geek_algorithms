package offer

import (
	"sort"
	"strconv"
	"strings"
)

// 剑指 Offer 45. 把数组排成最小的数
// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

// 示例 1:
// 输入: [10,2]
// 输出: "102"

// 示例 2:
// 输入: [3,30,34,5,9]
// 输出: "3033459"
// @lc: https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/

// 设数组 nums 中任意两数字的字符串为x、y ，则规定排序判断规则为：
// 1. 若拼接字符串：x + y > y + x，则 x “大于” y；
// 2. 反之，若 x + y < y + x，则 x “小于” y；
// x “小于” y 代表：排序完成后，数组中 x 应在 y 左边；“大于” 则反之。
// 根据以上规则，套用任何排序方法对 nums 执行排序即可
// 时间复杂度：O(nlog(n))
func minNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] < strs[j]+strs[i]
	})
	return strings.Join(strs, "")
}

// 分而治之
func quickSort(nums []int, left, right int) {
	// 分而治之，选择一个轴点povit，遍历剩下的元素将比povit小的交换到前面
	// 一次分治后，povit所在位置即为排序后最终的位置，返回povit的索引
	partition := func(nums []int, left, right int) int {
		povit := nums[left]
		j := left
		for i := left + 1; i <= right; i++ {
			if nums[i] < povit {
				j++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[left], nums[j] = nums[j], nums[left]
		return j
	}

	if len(nums) <= 1 {
		return
	}
	if left < right {
		povit := partition(nums, left, right)
		quickSort(nums, left, povit-1)
		quickSort(nums, povit+1, right)
	}
}
