package practice

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=449 lang=golang
 *
 * [449] 序列化和反序列化二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	Sep string // 分隔符
}

// func Constructor() Codec {
func NewCodec() Codec {
	return Codec{Sep: ","}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	return strings.Join(c.preorder(root), c.Sep)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	vals := strings.Split(data, c.Sep)
	return c.rebuildTree(vals, 0, len(vals)-1)
}

// 前序遍历：根->左->右
func (c *Codec) preorder(root *TreeNode) []string {
	vals := []string{}
	if root == nil {
		return vals
	}
	vals = append(vals, strconv.Itoa(root.Val))
	vals = append(vals, c.preorder(root.Left)...)
	vals = append(vals, c.preorder(root.Right)...)
	return vals
}

// 通过preorder 和 BST特点（BST中序遍历是严格升序数组）重建二叉树
func (c *Codec) rebuildTree(vals []string, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	// preorder 的第一个元素即为根节点
	rootVal, _ := strconv.Atoi(vals[start])
	// 由于BST满足：左子树 > rootVal > 右子树，因此：遍历vals数组找到第一个大于rootVal的数的位置 记为index 作为分界点，则满足：
	// 左子树部分为：vals[start+1 : index-1]， 右子树部分为：vals[index : end]。同理可得，递归构建左右子树
	index := end + 1 // 边界值处理，便于满足递归终止条件
	for i := start + 1; i <= end; i++ {
		intVal, _ := strconv.Atoi(vals[i])
		if intVal > rootVal {
			index = i
			break
		}
	}

	root := &TreeNode{Val: rootVal}
	// 递归构建左右子树
	root.Left = c.rebuildTree(vals, start+1, index-1)
	root.Right = c.rebuildTree(vals, index, end)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end
