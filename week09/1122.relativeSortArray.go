package week09

import "sort"

// 1122. 数组的相对排序
// 给你两个数组，arr1 和 arr2，
// arr2 中的元素各不相同
// arr2 中的每个元素都出现在 arr1 中
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
// 示例：
// 输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// 输出：[2,2,2,1,4,3,3,9,6,7,19]
// 提示：
// 1 <= arr1.length, arr2.length <= 1000
// 0 <= arr1[i], arr2[i] <= 1000
// arr2 中的元素 arr2[i] 各不相同
// arr2 中的每个元素 arr2[i] 都出现在 arr1 中
// @lc: https://leetcode-cn.com/problems/relative-sort-array/

// 方法一：计数排序
// 时间复杂度：O(n+m) n为arr1的长度，m为arr2 的长度。arr的长度为1001，常数
// 空间复杂度：O(n) n为arr1的长度
func relativeSortArray1(arr1 []int, arr2 []int) []int {
	// 题目提示arr2最大值小于1000，可以开辟一个长1001的数组存 数组arr1 的值和出现的次数
	counter := make([]int, 1001)
	for _, num1 := range arr1 {
		counter[num1]++
	}

	res := make([]int, 0, len(arr1))
	// 遍历arr2，将counter中 的arr2的值加入结果集，直到计数清零
	for _, v := range arr2 {
		for ; counter[v] > 0; counter[v]-- {
			res = append(res, v)
		}
	}
	// 处理剩下的 不存在于arr2中的数
	for v, cnt := range counter {
		for ; cnt > 0; cnt-- {
			res = append(res, v)
		}
	}
	return res
}

// 方法二：自定义比较器，调用库函数排序
// 对于2个元素x,y，他们的比较有3种情况：
// 	1. 都存在于arr2中，比较下标，下标小的在前面；
//	2. 其中一个存在于arr2中，则存在于arr2的放前面；
//	3. 都不存在于arr2中，直接比较大小即可。
// 根据以上思路，需要一个 hashmap 存储arr2中元素和下标的对应关系
// 时间复杂度：O(NlogN + m)，n为arr1的长度，m为arr2 的长度，快排时间复杂度
// 空间复杂度：O(NlogN)，快排空间复杂度
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	m := make(map[int]int, len(arr2))
	for k, v := range arr2 {
		m[v] = k
	}

	// 使用库sort函数进行比较，定义比较器
	sort.Slice(arr1, func(i int, j int) bool {
		x, y := arr1[i], arr1[j]
		indexX, hasX := m[x]
		indexY, hasY := m[y]
		if hasX && hasY { // case 1
			return indexX < indexY
		} else if hasX || hasY { // case 2
			return hasX
		} else { // case 3
			return x < y
		}
	})
	return arr1
}
