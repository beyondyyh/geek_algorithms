package week07

// 127. 单词接龙
// 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：
// 序列中第一个单词是 beginWord 。
// 序列中最后一个单词是 endWord 。
// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典 wordList 中的单词。
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。
// 示例 1：
// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
// 输出：5
// 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
// @lc: https://leetcode-cn.com/problems/word-ladder/

// 单向BFS
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	// setep1：先将 wordList 放到哈希表里，便于判断某个单词是否在 wordList 里
	wordSet := make(map[string]struct{})
	for _, wd := range wordList {
		wordSet[wd] = struct{}{}
	}
	// 如果 endWord 不存在于 wordList 中，则无法转换
	if _, has := wordSet[endWord]; !has {
		return 0
	}
	delete(wordSet, beginWord) // beginWord从set中删除，便于后续操作

	// step2：图的广度优先遍历，必须使用队列和表示是否访问过的 visited 哈希表
	q := &Queue{}
	q.Push(beginWord)
	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}

	// step3：开始广度优先遍历，包含起点，因此初始化的时候步数为 1
	step := 1
	for !q.IsEmpty() {
		size := q.Len()
		for i := 0; i < size; i++ {
			// 依次遍历当前队列中的单词
			curWord := q.Pop().(string)
			// 如果 curWord 能够修改 1 个字符与 endWord 相同，则返回 step + 1
			if changeWordByOneLetter(curWord, endWord, q, visited, wordSet) {
				return step + 1
			}
		}
		step++
	}
	return 0
}

// 方法二：双向广度优先遍历
// 已知目标顶点的情况下，可以分别从起点和目标顶点（终点）执行广度优先遍历，直到遍历的部分有交集。这种方式搜索的单词数量会更小一些；
// 更合理的做法是，每次从单词数量小的集合开始扩散；
// 这里 beginSet 和 endSet 交替使用，等价于单向 BFS 里使用队列，每次扩散都要加到总的 visited 里。
// @ref：https://leetcode-cn.com/problems/word-ladder/solution/yan-du-you-xian-bian-li-shuang-xiang-yan-du-you-2/
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	// setep1：先将 wordList 放到哈希表里，便于判断某个单词是否在 wordList 里
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}
	// 如果 endWord 不存在于 wordList 中，则无法转换
	if _, has := wordSet[endWord]; !has {
		return 0
	}

	// step2：已经访问过的 word 添加到 visited 哈希表里
	visited := make(map[string]struct{})
	// 分别用左边和右边扩散的哈希表代替单向 BFS 里的队列，它们在双向 BFS 的过程中交替使用
	beginSet, endSet := make(map[string]struct{}), make(map[string]struct{})
	beginSet[beginWord], endSet[endWord] = struct{}{}, struct{}{}

	// step3：执行双向 BFS，左右交替扩散的步数之和为所求
	step := 1
	for len(beginSet) > 0 && len(endSet) > 0 {
		// 优先选择小的哈希表进行扩散，考虑到的情况更少，此处保证 beginSet 是较小的集合
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		// 这里，保证 beginSet 是相对较小的集合，nextSets 在扩散完成以后，会成为新的 beginSet
		nextSet := make(map[string]struct{})
		for word := range beginSet {
			if changeWordByOneLetter2(word, endSet, visited, wordSet, nextSet) {
				return step + 1
			}
		}
		// 原来的 beginSet 废弃，从 nextSet 开始新的双向 BFS
		beginSet = nextSet
		step++
	}
	return 0
}

// changeWordByOneLetter2 尝试对 word 修改每一个字符，看看是不是能落在 endSet 中，扩展得到的新的 word 添加到 nextSet 里
func changeWordByOneLetter2(word string, endSet, visited, wordSet, nextSet map[string]struct{}) bool {
	for i := 0; i < len(word); i++ {
		for k := 'a'; k <= 'z'; k++ {
			if byte(k) == word[i] {
				continue
			}
			nextWord := word[:i] + string(k) + word[i+1:]
			if _, has := wordSet[nextWord]; has {
				if _, has := endSet[nextWord]; has {
					return true
				}
				if _, has := visited[nextWord]; !has {
					nextSet[nextWord] = struct{}{}
					visited[nextWord] = struct{}{}
				}
			}
		}
	}
	return false
}

// changeWordByOneLetter 尝试对 curWord 修改每一个字符，看看是不是能与 endWord 匹配
func changeWordByOneLetter(curWord, endWord string, q *Queue, visited, wordSet map[string]struct{}) bool {
	for i := 0; i < len(endWord); i++ {
		for k := 'a'; k <= 'z'; k++ {
			if byte(k) == curWord[i] {
				continue
			}
			// 替换 curWord 的第 i 个索引位置的字符
			nextWord := curWord[:i] + string(k) + curWord[i+1:]
			if _, has := wordSet[nextWord]; has {
				if nextWord == endWord {
					return true
				}
				// 注意：添加到队列以后，必须马上标记为已经访问
				if _, has := visited[nextWord]; has {
					q.Push(nextWord)
					visited[nextWord] = struct{}{}
				}
			}
		}
	}
	return false
}
