package week01

import "sort"

// 方法一：先将nums2 append到nums1中，再利用快排思想排序，简单好理解
// 时间复杂度：快排平均时间复杂度 O(m+n)log(m+n)
// 空间复杂度：快排平均空间复杂度 O(m+n)log(m+n)
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2) // func copy(dst, src []Type) int
	sort.Ints(nums1)
}

// 方法二：双指针法
// 分别初始化i，j个指针都从头部开始遍历，把较小值赋值给 sorted，同时自身指针后移一位，
// 直到其中一个走完，然后将剩下一个还未走完的直接追加给 sorted 即可
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n) 引入了额为存储 sorted
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	sorted := make([]int, 0, m+n)
	for {
		// nums1遍历结束，直接将nums2剩下的项目追加到sorted，结束
		if i == m {
			sorted = append(sorted, nums2[j:]...)
			break
		}
		// nums2遍历结束，直接将nums1剩下的项追加到sorted，结束
		if j == n {
			sorted = append(sorted, nums1[i:]...)
			break
		}

		// 将二者较小值追加到sorted，同时自身指针后移一位
		if nums1[i] <= nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	copy(nums1, sorted) // 题目要求直接改变nums1的长度，所以需要多次数组拷贝
}

// 方法三：逆向双指针，方法二的优化版本
// 方法二之所以用一个 sorted 数组来临时保存，是因为如果直接合并nums2到nums1中则会覆盖还没遍历到的部分。
// 优化思路：可以将nums2先放到nums1后面空的部分这样直接覆盖不会影响结果，因此可以指针从后往前遍历，每次取2者之中较大的放进nums1的最后面。
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func merge3(nums1 []int, m int, nums2 []int, n int) {
	// i:nums1的尾部，j:nums2的尾部，tail:合并的尾部
	i, j, tail := m-1, n-1, m+n-1
	for j >= 0 {
		if i < 0 || nums1[i] <= nums2[j] {
			nums1[tail] = nums2[j]
			j--
		} else {
			nums1[tail] = nums1[i]
			i--
		}
		tail--
	}
}
