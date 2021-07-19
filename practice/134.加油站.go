package practice

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
// 方法一：保留枚举，想象成一个环，环上每个点表示添加的油量，每条边表示消耗的油量，题目意思是：问我们从哪个节点出发，还可以回到该节点。只能顺时针方向走。
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func canCompleteCircuit1(gas []int, cost []int) int {
	n := len(gas)
	// 考虑从每一个点出发，进行枚举
	for i := 0; i < n; i++ {
		j := i           // j表示从i出发最远能到达第j个点
		remain := gas[i] // 到达i点时剩余的油量
		// 当前剩余的油能否到达下一个点
		for remain-cost[j] >= 0 {
			// 减去花费的加上新的点的补给
			// 因为是个环形，所以要求余，特殊case是最后一个点的情况
			remain = remain - cost[j] + gas[(j+1)%n]
			j = (j + 1) % n
			// j回到了i，表示可以绕环一周回到原点
			if j == i {
				return i
			}
		}
	}
	// 任何点都不可以回到出发点
	return -1
}

// 方法二：暴力优化
// 思路：如果从 i 点出发 最远能到达 j，那么[i+1,j)之间的点一定也是不能绕一周的。
// 求证，反证法，因为：假设 i+1 能绕一圈，则i+1一定能到达j+1，而 i 一定能到i+1，所以i能到达j+1，跟上面的i最远到达j相悖，所以不成立。
// 所以：下一次的 i 外层循环不需要从 i+1 开始了，可以直接从 j+1 开始，大大减少了外层的循环。特殊情况是到达末尾的时候，会回到 0
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		j := i           // j表示从i出发最远能到达第j个点
		remain := gas[i] // 到达i点剩余的油量
		// 当前剩余的油能否到达下一个点
		for remain-cost[j] >= 0 {
			remain = remain - cost[j] + gas[(j+1)%n]
			j = (j + 1) % n
			if j == i {
				return i
			}
		}
		// 最远距离绕到了 i 的前面，所以 i 后边的都不可能绕一圈了
		if j < i {
			return -1
		}
		// i 直接跳到 j，外层 for 循环执行 i++，相当于从 j+1 开始考虑
		i = j
	}
	return -1
}

// @lc code=end
