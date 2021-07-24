package practice

/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
// 基于一个事实： (rand_X() - 1) * Y + rand_Y() = rand_XY
// rand7()->[1,7], rand7()-1->[0,6], (rand7()-1)*7->[0,42], (rand7()-1)*7+rand7()->[1,49]=Rand49()
func rand10() int {
	var rand7 func() int // 去掉！！只是为了不报错
	///////////
	for {
		a, b := rand7(), rand7()
		// Rand49, [1,49]之间的均匀随机数
		num := (a-1)*7 + b
		// 说明生成的数在 [1,40] 之间，可以直接返回
		if num <= 40 {
			return num%10 + 1
		}

		// 说明第一次生成的在 [41,49] 之间，利用随机数再操作一遍
		a = num - 40 // [1,9]
		b = rand7()
		// Rand63(), [1,63]之间的均匀随机数
		num = (a-1)*7 + b
		// 说明生成的数在 [1,60] 之间，可以直接返回
		if num <= 60 {
			return num%10 + 1
		}

		// 说明第二次生成的在 [61,63] 之间
		a = num - 60 // [1,3]
		b = rand7()
		// Rand21, [1,21]之间的均匀随机数
		num = (a-1)*7 + b
		// 说明生成的数在 [1,20] 之间，可以直接返回
		if num <= 20 {
			return num%10 + 1
		}
	}
}

// @lc code=end
