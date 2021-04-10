package week02

import "strconv"

// 写一个程序，输出从 1 到 n 数字的字符串表示。
// 1. 如果 n 是3的倍数，输出“Fizz”；
// 2. 如果 n 是5的倍数，输出“Buzz”；
// 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
// @leetcode: https://leetcode-cn.com/problems/fizz-buzz
// 示例：n=6
// 返回：["1","2","Fizz","4","FizzBuzz","6"]

// 方法一： 字符串连接
// 对于 FizzBuzz 来说，只需要判断两个条件就可以了，而不需要像暴力法中那样判断三个条件。
// 同样的，对于 FizzBuzzJazz，只需要判断三个条件就可以了
// 举个例子，现在需要判断 15，步骤将会是下面这样的：
// 条件 1： 15 % 3 == 0, num_ans_str = "Fizz"
// 条件 2： 15 % 5 == 0, num_ans_str += "Buzz"
// => num_ans_str = "FizzBuzz"
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func fizzBuzz1(n int) []string {
	ans := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		divideBy3 := (i%3 == 0)
		divideBy5 := (i%5 == 0)

		var numstr string
		if divideBy3 {
			numstr += "Fizz"
		}
		if divideBy5 {
			numstr += "Buzz"
		}
		if numstr == "" {
			numstr = strconv.Itoa(i)
		}
		ans = append(ans, numstr)
	}
	return ans
}

// 方法二：针对方法1更优雅的写法，用map维护条件1-n
func fizzBuzz2(n int) []string {
	ans := make([]string, 0, n)
	dict := map[int]string{
		3: "Fizz",
		5: "Buzz",
		// and more...
	}
	for i := 1; i <= n; i++ {
		var numstr string
		for key, val := range dict {
			if i%key == 0 {
				numstr += val
			}
		}
		if numstr == "" {
			numstr += strconv.Itoa(i)
		}
		ans = append(ans, numstr)
	}
	return ans
}
