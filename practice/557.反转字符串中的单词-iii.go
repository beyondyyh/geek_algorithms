package practice

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	// 反转sb的[start,end]区间
	reverse := func(sb []byte, start, end int) {
		for ; start < end; start, end = start+1, end-1 {
			sb[start], sb[end] = sb[end], sb[start]
		}
	}

	// 字节数组
	sb := []byte(s)
	// start, end 表示每个单词的开始和结束位置；n 表示字符串长度
	start, end, n := 0, 0, len(sb)
	for start < n {
		for end < n && sb[end] != ' ' {
			end++
		}
		// 注意：end此时指向的是空格的位置，所以这里是 end-1
		// 至此，[start,end-1]表示一个完整的单词区间，反转之
		reverse(sb, start, end-1)
		// 更新 start,end 去找下一个单词
		start, end = end+1, end+1
	}
	return string(sb)
}

// @lc code=end
