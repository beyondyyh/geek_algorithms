package week04

import (
	"beyondyyh/geek_algorithms/kit"
)

type (
	TreeNode = kit.TreeNode
	ListNode = kit.ListNode
)

var (
	null = kit.NULL
	NULL = kit.NULL
)

type Stack []interface{}

func (s *Stack) Len() int           { return len(*s) }
func (s *Stack) IsEmpty() bool      { return s.Len() == 0 }
func (s *Stack) Push(x interface{}) { *s = append(*s, x) }
func (s *Stack) Pop() interface{}   { x := (*s)[s.Len()-1]; *s = (*s)[0 : s.Len()-1]; return x }
func (s *Stack) Peek() interface{}  { return (*s)[s.Len()-1] }

type Queue []interface{}

func (q *Queue) Len() int           { return len(*q) }
func (q *Queue) IsEmpty() bool      { return q.Len() == 0 }
func (q *Queue) Push(x interface{}) { *q = append(*q, x) }
func (q *Queue) Pop() interface{}   { x := (*q)[0]; *q = (*q)[1:]; return x }
