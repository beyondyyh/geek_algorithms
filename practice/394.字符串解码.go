package practice

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	num, res := 0, ""
	nums, strs := &Stack{}, &Stack{}
	for _, c := range []byte(s) {
		switch {
		case c == '[': // 将 '[' 前的数字压入nums栈内，字母字符串压入strs栈内
			nums.Push(num)
			strs.Push(res)
			num, res = 0, ""
		case c == ']': // 遇到 ']' 时，操作与之相配的 '[' 之间的字符，使用分配律
			var tmp string
			times := nums.Pop().(int) // Pop不能放到循环里
			for j := 0; j < times; j++ {
				tmp += res
			}
			res = strs.Pop().(string) + tmp
		case c >= '0' && c <= '9':
			num = num*10 + int(c-'0')
		default:
			res += string(c)
		}
		// fmt.Printf("c:%v, num:%d res:%s nums:%+v, strs:%+v\n", string(c), num, res, *nums, *strs)
	}
	return res
}

// @lc code=end
