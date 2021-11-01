package offerII

type ListNode struct {
	Val  int
	Next *ListNode
}

type DListNode struct {
	Val        int
	Next, Prev *ListNode
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack []interface{}

func (s *Stack) Len() int           { return len(*s) }
func (s *Stack) Push(x interface{}) { *s = append(*s, x) }
func (s *Stack) Pop() interface{}   { x := (*s)[s.Len()-1]; *s = (*s)[:s.Len()-1]; return x }
func (s *Stack) IsEmpty() bool      { return s.Len() == 0 }

type Queue []interface{}

func (q *Queue) Len() int           { return len(*q) }
func (q *Queue) Push(x interface{}) { *q = append(*q, x) }
func (q *Queue) Pop() interface{}   { x := (*q)[0]; *q = (*q)[1:]; return x }
func (q *Queue) IsEmpty() bool      { return q.Len() == 0 }
