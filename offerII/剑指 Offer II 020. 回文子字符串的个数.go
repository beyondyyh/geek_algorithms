package offerII

// 剑指 Offer II 020. 回文子字符串的个数
// 给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
// 示例 1：
// 输入：s = "abc"
// 输出：3
// 解释：三个回文子串: "a", "b", "c"
// 示例 2：
// 输入：s = "aaa"
// 输出：6
// 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
// @lc: https://leetcode-cn.com/problems/a7VOhD/

// 中心扩展法，以1个或2个字符为中心分别向两边扩散
// 此题与求最长回文子串类似，@lc: https://leetcode-cn.com/problems/longest-palindromic-substring/
// 时间复杂度：O(n^2) 遍历一次枚举中线点，每次以中心点向两边扩散
func countSubstrings(s string) int {
	if s == "" {
		return 0
	}
	// 以left和right为中心分别向两边扩散，找到有效的回文字串，返回回文串的个数
	// left == right 表示以一个字符为中心扩散，得到的回文串是奇数长度
	// left != right 表示以两个字符为中心扩散，得到的回文串是偶数长度
	expandAroundCenter := func(s string, left, right int) int {
		var num int
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			num++
		}
		return num
	}

	res := 0
	// 遍历字符串s，以每1个或2个字符为中心扩展
	for i := 0; i < len(s); i++ {
		res += expandAroundCenter(s, i, i) + // 1个字符为中心
			expandAroundCenter(s, i, i+1) // 2个字符为中心
	}
	return res
}
