package week07

// 433. 最小基因变化
// 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
// 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
// 现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。
// 注意：
// 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
// 如果一个起始基因序列需要多次变化，那么它每一次变化之后的基因序列都必须是合法的。
// 假定起始基因序列与目标基因序列是不一样的。
// 示例 1：
// start: "AACCGGTT"
// end:   "AACCGGTA"
// bank: ["AACCGGTA"]
// 返回值: 1

// ---------------------------------------------
// 回溯通用模板：
// result = []
// func backtracking(选择列表，路径) {
// 		if 满足结束条件:
//			result.add(路径)
//			return
// 		for 选择 in 选择列表：
//			做选择
//			backtracking(选择列表, 路径)
//			撤销选择
// }
// ---------------------------------------------

// DFS+回溯，
func minMutation1(start, end string, bank []string) int {
	m := map[byte][]string{
		'A': {"T", "G", "C"}, 'C': {"T", "G", "A"},
		'T': {"A", "G", "C"}, 'G': {"T", "A", "C"},
	}
	// 返回 s 在 bank 中的索引位置，不存在返回 -1
	indexOf := func(s string, bank []string) int {
		for i, v := range bank {
			if s == v {
				return i
			}
		}
		return -1
	}
	// 合法性判断，end 不在 bank 里，无法变异
	if indexOf(end, bank) == -1 {
		return -1
	}

	minCount := len(bank) + 1
	visited := make([]bool, len(bank))
	// dfs回溯，套用回溯标准模板
	// - 当前生成的 curr
	var backtracking func(string, int)
	backtracking = func(curr string, count int) {
		// terminator
		if curr == end {
			minCount = min(minCount, count)
			return
		}
		// 做选择、处理、撤销选择
		for i := 0; i < len(curr); i++ { // 遍历当前gene，逐个替换字符生成 next，看是否在 bank 里
			for _, s := range m[curr[i]] { // 遍历当前字符所对应的碱基
				next := curr[:i] + s + curr[i+1:] // 生成新的基因
				if idx := indexOf(next, bank); idx != -1 && !visited[idx] {
					visited[idx] = true         // marking
					backtracking(next, count+1) // drill down
					visited[idx] = false        // revert state
				}
			}
		}
	}
	backtracking(start, 0)

	if minCount > len(bank) {
		return -1
	}
	return minCount
}

// 单向BFS，队列
func minMutation2(start, end string, bank []string) int {
	m := map[byte][]string{
		'A': {"T", "G", "C"}, 'C': {"T", "G", "A"},
		'T': {"A", "G", "C"}, 'G': {"T", "A", "C"},
	}
	// 返回 s 在 bank 中的索引位置，不存在返回 -1
	indexOf := func(s string, bank []string) int {
		for i, v := range bank {
			if s == v {
				return i
			}
		}
		return -1
	}
	// 合法性判断，end 不在 bank 里，无法变异
	if indexOf(end, bank) == -1 {
		return -1
	}

	step := 0
	visited := make([]bool, len(bank))
	queue := []string{start}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			if curr == end {
				return step
			}
			// 遍历当前gene，逐个替换字符生成 next gene，看是否在 bank 里
			for j := 0; j < len(curr); j++ {
				for _, s := range m[curr[j]] {
					next := curr[:j] + s + curr[j+1:]
					if idx := indexOf(next, bank); idx != -1 && !visited[idx] {
						queue = append(queue, bank[idx]) // offer
						visited[idx] = true              // marking
					}
				}
			}
			step++
			queue = queue[1:]
		}
	}
	return -1
}

// 双向BFS，此题最优解
// 这是一个典型的广度优先搜索问题，和之前的单词接龙问题是一样的。[单词接龙 题解](https://leetcode-cn.com/problems/word-ladder/solution/dan-xiang-bfs-shuang-xiang-bfs-golangban-p1en/)
// 因为是求最短的转换路径问题，那么广度搜索是优于深度搜索的。
// 这里的做法就是将起点和终点分别加入两个集合当中，然后从集合元素少的一端开始搜索，这样能够减少搜索量，直到两个集合产生了交集，那么就结束搜索。
func minMutation3(start, end string, bank []string) int {
	// setep1：先将 bank 放到哈希表里，便于判断某个 newGene 是否在 bank 里
	bankSet := make(map[string]struct{}, len(bank))
	for _, v := range bank {
		bankSet[v] = struct{}{}
	}
	// 合法性校验，基因库中不包含目标，无法转换
	if _, has := bankSet[end]; !has {
		return -1
	}

	// step2：已经访问过的 gene 添加到 visited 哈希表里
	visited := make(map[string]struct{})
	// 分别用左边和右边扩散的哈希表代替单向 BFS 里的队列，它们在双向 BFS 的过程中交替使用
	startSet, endSet := make(map[string]struct{}), make(map[string]struct{})
	startSet[start], endSet[end] = struct{}{}, struct{}{}

	geneBase := []byte{'A', 'C', 'G', 'T'}
	// step3：执行双向 BFS，左右交替扩散的步数之和为所求
	step := 1
	for len(startSet) > 0 && len(endSet) > 0 {
		// 优先选择小的哈希表进行扩散，考虑到的情况更少，此处保证 startSet 是较小的集合
		if len(startSet) > len(endSet) {
			startSet, endSet = endSet, startSet
		}

		// 这里，保证 startSet 是相对较小的集合，nextSet 在扩散完成以后，会成为新的 startSet
		nextSet := make(map[string]struct{})
		for curr := range startSet {
			// 遍历当前字符串，逐个替换字符生成 next 字符串，看是否在 bank 里
			for i := 0; i < len(curr); i++ {
				for _, char := range geneBase {
					if curr[i] == char { // 相同的跳过
						continue
					}
					newGene := curr[:i] + string(char) + curr[i+1:]
					if _, has := bankSet[newGene]; has { // 有效的基因
						if _, has := endSet[newGene]; has { // 在 endSet 中找到了新基因，变异完成
							return step
						}
						if _, has := visited[newGene]; !has { // 在 endSet 中未找到，继续变异
							nextSet[newGene] = struct{}{}
							visited[newGene] = struct{}{}
						}
					}
				}
			}
		}

		// 原来的 startSet 废弃，从 nextSet 开始新的双向 BFS
		startSet = nextSet
	}

	return -1
}
