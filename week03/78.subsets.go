package week03

// index 就是循环到第几层
// 逻辑：每一层选 或者 不选这个数
// list 存储的是中间结果
// ans 存储最终结果

// 方法一：迭代法
// 思路：先把空数组[]加到结果集res里，遍历都把之前已经放入res的子集合都拿出再追加上当前元素形成新的子集合
// 时间复杂度：O(n * 2^n)
// 空间复杂度：O(n) 即构造子集使用的临时数组 newset 的空间代价
func subsets1(nums []int) [][]int {
	res := [][]int{} // 初始化结果集
	if len(nums) == 0 {
		return res
	}

	res = append(res, []int{}) // 把空数组放入结果集
	for _, num := range nums {
		for _, subset := range res {
			newset := make([]int, len(subset), len(subset)+1)
			copy(newset, subset) // copy subset to newset
			newset = append(newset, num)
			res = append(res, newset)
		}
	}
	return res
}

// 方法二：递归、回溯
func subsets2(nums []int) [][]int {
	res := [][]int{} // 存放最终结果集
	if len(nums) == 0 {
		return res
	}

	dfs(&res, nums, []int{}, 0)
	return res
}

// dfs 递归、回溯
// - res 	最终的结果集
// - nums	input数组
// - subset 每一层逻辑的中间结果
// - index 	当前第几层，也就是遍历到第几个元素
func dfs(res *[][]int, nums []int, subset []int, index int) {
	// 1. terminator 终止条件
	if index == len(nums) {
		*res = append(*res, subset)
		return
	}

	// 2. process current logic & drill down
	// 2.1>. 不选当前index的元素，subset不需要变更
	dfs(res, nums, subset, index+1)
	// 2.2>. 选当前index的元素，把当前元素加入subset
	newsubset := append(subset, nums[index])
	dfs(res, nums, newsubset, index+1)

	// 3. revert current states，清理当前层
	// newsubset是局部变量，所以不需要revert
}

// 方法三：方法二的优雅写法
// 时间复杂度：O(n * 2^n)，一共 2^n种状态树，每种状态需要O(n)的时间来构建子集
// 空间复杂度：O(n)，临时slice newsubset的空间代价为O(n)
func subsets3(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	// dfs
	// - index 	当前第几层，也就是遍历到第几个元素
	// - subset	每一层逻辑的中间结果
	var dfs func(int, []int)
	dfs = func(index int, subset []int) {
		// 1. terminator 递归终止条件
		if index == len(nums) {
			res = append(res, append([]int{}, subset...))
			return
		}

		// 2. 处理当前层逻辑 & drill down
		// 2.1> 不选择当前元素
		dfs(index+1, subset)
		// 2.2> 选择当前元素
		newsubset := append(subset, nums[index])
		dfs(index+1, newsubset)

		// 3. revert states, subset是局部变量不需要revert
	}

	dfs(0, []int{})
	return res
}
