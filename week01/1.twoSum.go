package week01

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// @leetcode https://leetcode-cn.com/problems/two-sum

// 方法一：暴力枚举
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ { // 注意i的右边界到n-2就可以了没必要到n-1，因为j在i后面，需要给j留一个位置  <n-1 --> <=n-1
		diff := target - nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] == diff {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 方法二：利用hashmap，存储元素，方法一的优化版，在数组中找diff需要遍历，时间复杂度为O(n)，可以借助map实现时间复杂度为O(1)的查找
// nums[i] 和 target-nums[i] 一定存在与nums数组中，所以遍历nums可以把(target-nums[i])存入map，去check diff是否存在，存在则找到
// 题目要求返回数组下标，所以可以将k=diff, v=index存入map，便于操作
// 时间复杂度：O(n)，只需要遍历一次数据，n为数组长度
// 空间复杂度：O(n)，引入了额外空间map
func twoSum2(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, ok := hmap[num]; ok {
			return []int{i, j}
		}
		hmap[diff] = i
	}
	return nil
}
