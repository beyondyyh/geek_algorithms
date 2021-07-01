package practice

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
// 滑动窗口 + 数组
func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	if n < m {
		return nil
	}

	res := []int{}
	cnt_s, cnt_p := [26]int{}, [26]int{}
	for i := 0; i < m; i++ {
		cnt_p[p[i]-'a']++
		cnt_s[s[i]-'a']++
	}

	// 第一个窗口形成时如果能匹配上，则加入res
	if cnt_s == cnt_p {
		res = append(res, 0)
	}

	// 滑动窗口，窗口区间：[0, m) -> (i-m, i] -> [i-m+1, i]
	for i := m; i < n; i++ {
		cnt_s[s[i-m]-'a']-- // 左侧滑出窗口的元素去掉
		cnt_s[s[i]-'a']++   // 右侧进入窗口的原始加入
		// 新窗口形成时 cnt_s 和 cnt_p 是否相等
		if cnt_s == cnt_p {
			res = append(res, i-m+1)
		}
	}
	return res
}

// @lc code=end

// 思路二：滑动窗口 + 双指针
func findAnagrams2(s string, p string) []int {
	n, m := len(s), len(p)
	if n < m {
		return nil
	}

	res := []int{}
	cnt_s, cnt_p := [26]int{}, [26]int{}
	for i := 0; i < m; i++ {
		cnt_p[p[i]-'a']++
	}

	left, right := 0, 0
	// 右窗口开始不断向右移动
	for ; right < n; right++ {
		curRight := s[right] - 'a'
		// 将右窗口当前访问到的元素个数加1
		cnt_s[curRight]++
		// 当前窗口数组中 curRight 比 cnt_p 数组中对应元素的个数要多的时候就该移动左窗口指针
		for cnt_s[curRight] > cnt_p[curRight] {
			curLeft := s[left] - 'a'
			// 将左窗口当前访问到的元素个数减1
			cnt_s[curLeft]--
			left++
		}
		if right-left+1 == m {
			res = append(res, left)
		}
	}
	return res
}
