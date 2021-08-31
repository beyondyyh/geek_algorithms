package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type treeTraversal func(*TreeNode) []int

// run: go test -run Test_TreeNode
func Test_TreeNode(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name   string
		nums   []int         // ints数组
		method treeTraversal // 二叉树遍历方法
		expect []int         // 预期结果数组
	}{
		{
			name:   "x1-1",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			method: Preorder,
			expect: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:   "x1-2",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			method: PreorderIter,
			expect: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:   "x1-2-2",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			method: PreorderIter2,
			expect: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:   "x2-1",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			method: Inorder,
			expect: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name:   "x2-2",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			method: InorderIter,
			expect: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name:   "x2-2-2",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			method: InorderIter2,
			expect: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name:   "x3-1",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			method: Postorder,
			expect: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			name:   "x3-2",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			method: PostorderIter,
			expect: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for _, c := range cases {
		root := Ints2Tree(c.nums)
		assert.Equal(c.expect, c.method(root), c.name)
	}
}
