package week02

import "sort"

// 给定两个数组，编写一个函数来计算它们的交集。
// 示例 1：
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
// @leetcode: https://leetcode-cn.com/problems/intersection-of-two-arrays-ii

// 说明：
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
// 我们可以不考虑输出结果的顺序。
// 进阶：
// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

// 方法一：将数组排序后，利用双指针法进行遍历，如果nums1[i]=nums2[j]则push到结果集里，
// 否则移动较小元素的指针
// 时间复杂度：O(mlogm+nlogn)，2次快排平均时间复杂度为 O(NlogN) + 2遍历O(n)，忽略常数和只保留最高指数项故为O(NlogN)
// 空间复杂度：O(min(m,n))
func intersect1(nums1 []int, nums2 []int) []int {
	// sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	// sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
	sort.Ints(nums1)
	sort.Ints(nums2)

	n1, n2 := len(nums1), len(nums2)
	ans := make([]int, 0, min(n1, n2))
	for i, j := 0, 0; i < n1 && j < n2; {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法二：借助hashmap
// 将较短数组的值存入map m[元素]元素个数，遍历另一个较长数组，如果m[num]>0说明num在2个数组都存在则加入结果集，同时个数减一
// 时间复杂度：O(m+n)
// 空间复杂度：O(min(m,n))
func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect2(nums2, nums1)
	}
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}

	ans := make([]int, 0, len(m))
	for _, num := range nums2 {
		if m[num] > 0 {
			ans = append(ans, num)
			m[num]--
		}
	}
	return ans
}
