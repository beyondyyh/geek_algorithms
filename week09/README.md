# 布隆过滤器、LRU缓存、排序算法、高级动态规划

## Bloom filter

**原理：**
布隆过滤器（Bloom Filter）的核心实现是一个超大的位数组和K个哈希函数。将每个要添加的元素通过K个哈希函数，得到对应于位数组上的k个位置，然后都标记为1。主要用于判断一个元素一定不存在的场景，有一定的误判率。

**主要使用场景：**

- 网络爬虫里对url去重，或记录已经访问过的url集合，防止重复访问；
- 反垃圾邮件，从10亿个邮箱里判断某些邮箱是否是垃圾邮箱，同理：垃圾短信；
- 缓存击穿，将确定存在于db的记录放到布隆过滤器里，当黑客等大量访问不存在的记录时快速返回，防止将请求都达到db上把服务打挂。

**添加元素：**

- 将要查询的元素通过k个哈希函数，得到位数组上的k个位置
- 将这k个位置标记为1

**查询元素：**

- 将要查询的元素通过k个哈希函数，得到位数组上的k个位置，
- `如果这 k 个位置中有一个为0，则肯定不存在集合中`，一般应用场景都是使用这个特性。
- 如果这 k 个位置全部为1，则表示极大概率存在集合中

## LRU cache

- [基于 双链表+hashmap 代码实现](146.LRUCache.go)
- [基于 标准库 `container/list` 代码实现](146.LRUCache-2.go)

## 排序算法

**相关概念：**

- **稳定：** 如果a原本在b前面，而a=b，排序之后a仍然在b的前面。
- **不稳定：** 如果a原本在b的前面，而a=b，排序之后 a 可能会出现在 b 的后面。
- **时间复杂度：** 对排序数据的总的操作次数。反映当n变化时，操作次数呈现什么规律。
- **空间复杂度：** 是指算法在计算机内执行时所需存储空间的度量，它也是数据规模n的函数。

### 十大经典排序算法

> [一篇介绍十大经典排序的文字，附动图](https://www.cnblogs.com/onepixel/p/7674659.html)

**十种常见排序算法可以分为两大类：**

- **比较类排序：** 通过比较来决定元素间的相对次序，由于其时间复杂度不能突破O(nlogn)，因此也称为非线性时间比较类排序。
- **非比较类排序：** 不通过比较来决定元素间的相对次序，它可以突破基于比较排序的时间下界，以线性时间运行，因此也称为线性时间非比较类排序。 

<img src="../imgs/sorts.png" width="750" />

> 重点关注 `快排`、`归并排序`、`堆排序`

### 快排

**时间复杂度：** `O(n log(n))`，最坏为：O(n^2)，不稳定排序

**空间复杂度：** `O(n log(n))`

**基本思想：** 分治的思想，通过一趟遍历将数据分割成两部分，其中一部分比另一部分小，进而递归对这两部分继续进行排序

**算法描述：**

- 首先选择一个轴点，称为基准（`pivot`）；
- 重新排序数列，所有元素比基准值小的放在基准前面，比基准值大的放在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（`partition`）操作；
- 递归调用（`recursion`），分别把小于基准和大于基准的数组排序；

**示例代码：**

```golang
func quickSort(arr []int, left, right int) {
	if len(arr) <= 1 {
		return
	}

	// 分区操作，返回轴点索引下标
	var partition func(arr []int, left, right int) int
	partition = func(arr []int, left, right int) int {
		pivot := left      // 选取第一个元素作为轴点
		index := pivot + 1 // 遍历的游标
		for i := index; i <= right; i++ {
			if arr[i] < arr[pivot] {
				arr[i], arr[index] = arr[index], arr[i] // 比轴点小的交换到前面
				index++
			}
		}
		arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
		return index - 1
	}

	// recursion
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
	return
}
```

### 归并排序

**时间复杂度：** `O(n log(n))`，稳定排序！

**空间复杂度：** `O(n)`，2-路归并的额外空间

**基本思想：** 分治的思想典型应用，不断切割数组成为左右两部分，直到只包含1个元素，此时可以认为是有序的，然后合并2个有序数组，将两个有序表合并成一个有序表的过程称为：`2-路归并`。

**算法描述：**

- 把长度为 n 的数组分成两个长度为 n/2 的子数组；
- 对这两个子序列分别采用归并排序；
- 将两个排序好的子数组合并成一个最终的排序数组。

**示例代码：**

```golang
func mergeSort(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr // 最后切割只剩下一个元素，可以认为是有序的
	}

	// 合并2个有序数组
	var _merge func(left, right []int) []int
	_merge = func(left, right []int) []int {
		m, n := len(left), len(right)
		res := make([]int, 0, m+n)

		var i, j int
		for i < m && j < n {
			if left[i] <= right[j] {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}
		if i < m {
			res = append(res, left[i:]...)
		}
		if j < n {
			res = append(res, right[j:]...)
		}
		return res
	}

  	// 切分成左右2部分，分别进行归并排序
	mid := n / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return _merge(left, right)
}
```

### 堆排序

**时间复杂度：** `O(n log(n))`，稳定排序！，每次弹出堆顶元素后，heapifyDown 的时间复杂度

**空间复杂度：** `O(1)`，利用堆的特性，不需要额外存储

**基本思想：** 堆是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点，同理可得所有子节点也满足该特性

**示例代码：**

```golang
func heapSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	res := make([]int, 0, n)
	// 构建小顶堆
	h := &iheap{arr}
	heap.Init(h)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

// 利用golang标准库实现堆，只需要实现 heap接口的 `Push, Pop`方法即可
type iheap struct {
	sort.IntSlice // 默认升序
}

// // 默认升序，倒序的话需要重写 sort.Less 比较器
// func (h *iheap) Less(i, j int) bool {
// 	return h.IntSlice[i] > h.IntSlice[j]
// }

// Push(), Pop() 实现container/heap 的接口，不仅要改变元素值，还要改变数组长度，所以receiver用指针
func (h *iheap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

// Pop, 删除堆顶元素
// 注意：此处的x并不是真正的堆顶元素，而是数组的最后一个元素用来替换堆顶的，Pop只是实现了标准库 `container/heap`的Pop接口，内部还有调整堆结构的一系列操作
// 想一想二叉堆的删除操作：
//   - 将堆尾部元素替换到堆顶，即数组最后一个元素替换第一个元素
//   - 然后依次从堆顶向下调整整个堆的结构，（一直到堆尾即可），`heapifyDown`
// go标准库源码：https://golang.org/src/container/heap/heap.go?s=2190:2223#L50
func (h *iheap) Pop() interface{} {
	n := h.IntSlice.Len()
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}
```

## 高级动态规划

### 动态规划、递归、分治的区别

- **关键点：** 动态规划 和 递归、或 分治没有根本上的区别，关键看有无最优子结构
- **共性：** 找到最近重复子问题
- **差异性：** 最优子结构、中途可以淘汰次优解（必须淘汰次优解）时间复杂度从指数级（泛型递归）降到多项式级（O(n)或O(n^2)）

### 动态规划的关键点

- 确定dp数组（dp table）以及下标的含义，dp的含义不同写出的转移方程也不同，一般bottom-up比较好理解；
- 最优子结构，找到子问题 dp[n] = best_of(dp[n-1], dp[n-2])，第n步的值是前面几个值的最佳值（可能是累加、最大、最小等）；
- 递推公式（美其名曰：状态转移方程或DP方程）；
- dp状态初始化；
- 确定遍历顺序，根据dp转移方程进行递推；
- 举例推导dp数组。

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[146. LRU 缓存机制](https://leetcode-cn.com/problems/lru-cache/)|[146.LRUCache.go](146.LRUCache.go)|M|DListNode+hashmap，标准库container/list|
|[1122. 数组的相对排序](https://leetcode-cn.com/problems/relative-sort-array/)|[1122.relativeSortArray.go](1122.relativeSortArray.go)|S|计数排序、自定义排序|
|[56. 合并区间](https://leetcode-cn.com/problems/merge-intervals/)|[56.merge.go](56.merge.go)|S|排序、合并区间|
|[300. 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)|[300.lengthOfLIS.go](300.lengthOfLIS.go)|M|dp|
|[746. 使用最小花费爬楼梯](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)|[746.minCostClimbingStairs.go](746.minCostClimbingStairs.go)|S|dp、爬楼梯|
