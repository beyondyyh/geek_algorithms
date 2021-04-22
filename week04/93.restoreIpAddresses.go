package week04

import (
	"strconv"
	"strings"
)

// 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 示例 1：
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// @leetcode: https://leetcode-cn.com/problems/restore-ip-addresses

// 回溯思想典型案例，难点在于递归终止条件 和 剪枝逻辑
func restoreIpAddresses(s string) []string {
	size := len(s)
	// 如果长度不够，不搜索
	if size < 4 || size > 12 {
		return nil
	}

	// validIPSegment 判断s的子区间 [start, end) 是否为一个有效的IP段
	// 顺便转成字符串，方便后续直接操作，是：返回IP段，否：返回空
	var validIPSegment = func(s string, start, end int) string {
		length := end - start
		// 长度大于1位的时候，不能以0开头，如：0.1.2.2合法，01.1.2.2就不合法，就这么个意思
		if length > 1 && s[start] == '0' {
			return ""
		}

		// 截取片段
		segment := string(s[start:end])
		if segInt, err := strconv.Atoi(segment); err != nil || segInt > 255 {
			return ""
		}
		return segment
	}

	res := []string{}
	// backtrack 回溯
	// - pos	截取IP段的起始位置
	// - remain	剩下的IP段数量，包含当前正在截取的这段
	// - path 	截取的IP段的结果栈，回溯的中间状态
	var backtrack func(int, int, []string)
	backtrack = func(pos, remain int, path []string) {
		// 1.terminator, 如果截取的起始位置已经到达字符串尾部了，则本次梦境结束
		if pos == size {
			// 剩下的还需要截取的IP段数量为0，说明已经4段了，满足要求
			if remain == 0 {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		// 剪枝逻辑：如果剩余的字符总数大于 剩下的可供截取的IP段的数量的最大值(假设每段都是3位)，则不满足条件，本次梦境结束
		if size-pos > remain*3 {
			return
		}

		// 每一个结点可以选择截取的方法只有3种：截1位、截2位、截3位
		for i := pos; i < pos+3; i++ {
			// 如果截取的起始位置到达字符串尾部了，则退出
			if i >= size {
				break
			}
			// 判断截取的这一片段是否是合法的IP段
			ipSegment := validIPSegment(s, pos, i+1)
			// 如果是，则选择、处理、回撤 三部曲
			if "" != ipSegment {
				path = append(path, ipSegment)
				backtrack(i+1, remain-1, path)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0, 4, []string{})
	return res
}

// 暴力法，非常暴力，牛逼克拉斯
func restoreIpAddresses1(s string) []string {
	res := []string{}
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				for d := 1; d < 4; d++ {
					if a+b+c+d == len(s) {
						n1, _ := strconv.Atoi(s[:a])
						n2, _ := strconv.Atoi(s[a : a+b])
						n3, _ := strconv.Atoi(s[a+b : a+b+c])
						n4, _ := strconv.Atoi(s[a+b+c:])
						if n1 <= 255 && n2 <= 255 && n3 <= 255 && n4 <= 255 {
							segments := append([]string{}, strconv.Itoa(n1), strconv.Itoa(n2), strconv.Itoa(n3), strconv.Itoa(n4))
							IP := strings.Join(segments, ".")
							if len(IP) == len(s)+3 {
								res = append(res, IP)
							}
						}
					}
				}
			}
		}
	}
	return res
}
