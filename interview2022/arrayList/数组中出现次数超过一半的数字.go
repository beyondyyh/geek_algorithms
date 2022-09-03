package arrayList

// 方法一：利用hashmap计数
// 时间复杂度：O(n)
// 空间复杂度：O(n)

// 方法二：统计计数，也叫选举法
// 思路：遍历数组，从第一个元素开始计数，计数器初始为 1，遇到与当前元素相同的元素时，计数器加1，否则减 1（比拼消耗）。
// 当计数器为0时，重新从当前元素开始计数，重复上面步骤直至到达数组末尾。
// 存在的问题：
// 如果数组中没有超过一半的数字，此时返回的是最后一个元素而不是-1
// 因此需要一个counting方法，统计某个元素在数组中是否真的超过一半，不改变时间复杂度
func majorityElement(nums []int) int {
	majority := votes(nums)
	// 通过计数计算下该众数是否超过一半
	votes := counting(nums, majority)
	if votes <= len(nums)/2 {
		majority = -1
	}
	return majority
}

// 经典的摩尔投票法则：返回数组中的那个众数
func votes(nums []int) int {
	res, votes := -1, 0
	for _, num := range nums {
		if votes == 0 {
			votes = 1
			res = num
		} else {
			if num == res {
				votes++
			} else {
				votes--
			}
		}
	}
	return res
}

// counting 计数：统计k在数组中出现的次数
func counting(nums []int, k int) int {
	votes := 0
	for _, num := range nums {
		if num == k {
			votes++
		}
	}
	return votes
}
