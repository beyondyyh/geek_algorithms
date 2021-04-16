package week03

import (
	"beyondyyh/geek_algorithms/kit"
	"fmt"
	"strconv"
)

type (
	TreeNode = kit.TreeNode
)

var (
	null = kit.NULL
	NULL = kit.NULL
)

type Istack []interface{}

func (s *Istack) Len() int {
	return len(*s)
}

func (s *Istack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Istack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Istack) Pop() interface{} {
	x := (*s)[s.Len()-1]
	*s = (*s)[0 : s.Len()-1]
	return x
}

func (s *Istack) Peek() interface{} {
	return (*s)[s.Len()-1]
}

func (s *Istack) print() {
	var ele []interface{}
	for _, e := range *s {
		if item, ok := e.(item); ok {
			val := "nil"
			color, node := item[0].(int), item[1].(*TreeNode)
			if node != nil {
				val = strconv.Itoa(node.Val)
			}
			ele = append(ele, [2]interface{}{color, val})
		}
	}
	fmt.Printf("%+v\n", ele)
}
