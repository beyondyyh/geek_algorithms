package practice

/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
// 手撕快排
// 思路：分治的思想，选择一个轴点pivot（一般选第一个），遍历剩下的元素将比pivot小的交换到前面，一轮结束后将pivot放到排序后应该在的位置
// 递归调用
func sortArray(nums []int) []int {
	var quickSort func(arr []int, left, right int) []int
	quickSort = func(arr []int, left, right int) []int {
		if len(arr) <= 1 {
			return arr
		}
		// 分区操作，返回轴点索引下标
		// 在数组arr的子区间 [left, right] 执行 partition 操作，返回 arr[left] 排序以后应该在的位置
		// 在遍历过程中保持循环不变量的语义
		// 1、[left + 1, j] < arr[left]
		// 2、(j, i] >= arr[left]
		// 3、交换 arr[left]和arr[j]
		partition := func(arr []int, left, right int) int {
			j := left
			pivot := arr[left] // 选取第一个元素作为轴点
			for i := left + 1; i <= right; i++ {
				if arr[i] < pivot { // 小于 pivot 的元素都被交换到前面
					j++
					arr[j], arr[i] = arr[i], arr[j]
				}
			}
			// 在之前遍历的过程中，满足 [left + 1, j] < pivot，并且 (j, i] >= pivot
			arr[j], arr[left] = arr[left], arr[j]
			// 交换以后 [left, j - 1] < pivot, nums[j] = pivot, [j + 1, right] >= pivot
			return j
		}

		// recursion
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
		return arr
	}

	left, right := 0, len(nums)-1
	return quickSort(nums, left, right)
}

// @lc code=end
