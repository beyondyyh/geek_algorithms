package week07

import "beyondyyh/geek_algorithms/kit"

type (
	Queue = kit.Queue
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
