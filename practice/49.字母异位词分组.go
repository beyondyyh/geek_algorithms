package practice

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
// 类似 242.有效的字母异位词 ，map的可以可以是数组 map[[26]int]xxx
// 时间复杂度：O(n(k+26))，n 是 strs 中的字符串的数量，k 是 字符串的的最大长度
// 空间复杂度：O(n(k+26))
func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, c := range str {
			cnt[c-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	res := make([][]string, 0, len(mp))
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

// @lc code=end
