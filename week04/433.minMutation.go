package week04

// 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
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
// @leetcode: https://leetcode-cn.com/problems/minimum-genetic-mutation

// 方法一：BFS
func minMutation1(start string, end string, bank []string) int {
	m := map[byte][]string{
		'A': {"T", "G", "C"}, 'C': {"T", "G", "A"},
		'T': {"A", "G", "C"}, 'G': {"T", "A", "C"},
	}
	var indexOf = func(s string, bank []string) int {
		for i, v := range bank {
			if s == v {
				return i
			}
		}
		return -1
	}
	// 合法性校验
	if -1 == indexOf(end, bank) {
		return -1
	}

	count := 0
	visited := make([]bool, len(bank))
	queue := []string{start}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			curr := queue[i]
			if curr == end {
				return count
			}
			for j := 0; j < len(curr); j++ {
				for _, s := range m[curr[j]] {
					gene := curr[:j] + s + curr[j+1:]
					// fmt.Printf("j:%d s:%s tmp:%s, visited:%+v\n", j, s, gene, visited)
					if idx := indexOf(gene, bank); idx != -1 && !visited[idx] {
						queue = append(queue, bank[idx])
						visited[idx] = true
					}
				}
			}
		}
		count++
		queue = queue[1:]
	}

	return -1
}

// 方法二：DFS+回溯
func minMutation2(start string, end string, bank []string) int {
	m := map[byte][]string{
		'A': {"T", "G", "C"}, 'C': {"T", "G", "A"},
		'T': {"A", "G", "C"}, 'G': {"T", "A", "C"},
	}
	var indexOf = func(s string, bank []string) int {
		for i, v := range bank {
			if s == v {
				return i
			}
		}
		return -1
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 合法性校验
	if -1 == indexOf(end, bank) {
		return -1
	}

	minCount := len(bank) + 1
	visited := make([]bool, len(bank))
	var dfs func(string, int)
	dfs = func(path string, count int) {
		// terminator
		if path == end {
			minCount = min(minCount, count)
			return
		}
		// 选择、处理逻辑、撤销选择
		for i := 0; i < len(path); i++ {
			for _, s := range m[path[i]] {
				gene := path[:i] + s + path[i+1:]
				if idx := indexOf(gene, bank); idx != -1 && !visited[idx] {
					visited[idx] = true
					dfs(gene, count+1)
					visited[idx] = false
				}
			}
		}
	}
	dfs(start, 0)

	if minCount > len(bank) {
		return -1
	}

	return minCount
}
