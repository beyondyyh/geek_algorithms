package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

// --- stack 基本实现
type Stack []interface{}

func (s *Stack) Len() int           { return len(*s) }
func (s *Stack) IsEmpty() bool      { return s.Len() == 0 }
func (s *Stack) Push(x interface{}) { *s = append(*s, x) }                                    // 将x入栈，size++
func (s *Stack) Pop() interface{}   { x := (*s)[s.Len()-1]; *s = (*s)[:s.Len()-1]; return x } // 弹出栈顶元素，size--
func (s *Stack) Peek() interface{}  { return (*s)[s.Len()-1] }                                // 查看栈顶元素，size不变

// // --- Queue 基本实现
type Queue []interface{}

func (q *Queue) Len() int           { return len(*q) }
func (q *Queue) IsEmpty() bool      { return q.Len() == 0 }
func (q *Queue) Push(x interface{}) { *q = append(*q, x) }                    // 将x入队，size++
func (q *Queue) Pop() interface{}   { x := (*q)[0]; *q = (*q)[1:]; return x } // 弹出队头元素，size--
func (q *Queue) Peek() interface{}  { return (*q)[0] }                        // 查看队头元素，size不变
