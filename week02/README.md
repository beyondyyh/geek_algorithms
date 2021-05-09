# week02 哈希表、树、二叉树、二叉搜索树、堆、二叉堆、图

## 数据结构

### 哈希表

- 哈希表也叫散列表，是根据关键码值(Key value)而直接进行访问的数据结构，查找删除效率 `O(1)`
- 它通过散列函数(Hash Function)将key映射到表中一个位置，存放记录的数组叫 作哈希表(或散列表)

**工业应用：** 缓存(LRU Cache)、键值对存储(redis)等

### 树

- Linked-List 是特殊化的Tree
- Tree 是特殊化的 Graph

### 二叉树

**定义：**

```golang
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

**二叉树遍历：**

- 前序(Pre-order):根-左-右
- 中序(In-order):左-根-右
- 后序(Post-order):左-右-根

**示例代码：**

```golang
func Preorder(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    res = append(res, root.Val)
    res = append(res, Preorder(root.Left)...)
    res = append(res, Preorder(root.Right)...)
    return res
}

func Inorder(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    res = Inorder(root.Left)
    res = append(res, root.Val)
    res = append(res, Inorder(root.Right)...)
    return res
}

func Postorder(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }

    res = append(res, Postorder(root.Left)...)
    res = append(res, Postorder(root.Right)...)
    res = append(res, root.Val)
    return res
}
```

### `二叉搜索树 Binary Search Tree`

二叉搜索树，也称二叉排序树、有序二叉树(Ordered Binary Tree)、排序二叉树(Sorted Binary Tree)，是指 `一棵空树` 或者具有下列性质的二叉树:

1. 左子树上所有结点的值均小于它的根结点的值;
2. 右子树上所有结点的值均大于它的根结点的值;
3. 以此类推: `左、右子树也分别为二叉查找树`。 （这点很重要!）

**重要：**

- 中序遍历 即为升序排列
- 查找、插入、删除的平均时间复杂度为 `O(logN)`

## 堆 Heap

- 堆是可以迅速的找到一堆数据中的最大值或最小值的数据结构。
- 根节点最大的叫`大顶堆` 或 大根堆，根节点最小的叫`小顶堆` 或小根堆
- 常见的堆有 二叉堆、斐波那契堆等
- 时间复杂度：
  - 找最大 or 最小：O(1)
  - 删除最大 or 最小：O(logN)
  - 插入 or 创建：O(logN) or O(1)

各类高级语言里，堆其实是个接口，如golang的 `container/heap`

**工业应用：** 优先队列等

## 二叉堆 Binary Heap

- 通过完全二叉树实现（注意不是二叉搜索树，二叉搜索树的中序遍历是有序的，堆无序），`所谓完全二叉树：二叉树中的每一级节点除了最下面的一层，都是满的，也就是都有左右2个子节点。`
- 如二叉堆（大顶堆）特点：
  - 是一颗完全树
  - 树中任意节点的值总是 >= 其左右子节点的值
- 时间复杂度：
  - 找最大 or 最小：O(1)
  - 删除最大 or 最小：O(logN)

**重点：** 由于堆只需要找最大值or最小值不需要访问下面的节点，所以不需要用链表，用一个以为数组即可实现堆则满足：

- 数组第一个元素放堆顶元素
- 任意第k个节点的左子节点索引为：(2*k + 1)，右子节点索引为：(2*k + 2)，父节点索引为：floor((k-1)/2)

**插入元素x操作（大顶堆）：**

- 新元素一律先插入到堆的尾部，heap.size()++, heap[size()] = x
- 依次从堆尾向上调整整个堆的结构，（一直到根即可），内部叫 `heapfiyUp`，即新元素与它的父节点进行比较，如果比父节点大则进行交换，依次往上，直到父节点比它小，最差到堆顶，时间复杂度为：`O(logN)`

**删除堆顶操作（大顶堆）：**

- 将堆尾部元素替换到堆顶，堆顶被替换掉即被删除，heap.size()--, heap[0]=heap[size()-1]
- 依次从堆顶向下调整整个堆的结构，（一直到堆尾即可），内部叫 `heapifyDown`，即拿堆顶元素依次与它的左右子节点进行比较，找到比它大的最大的子节点，然后进行交换，直到它子节点都大，最差到堆尾，时间复杂度为：`O(logN)`

**工业应用：** 优先队列等

**高频题目：**

- [最小的K个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)
- [滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)
- [前K个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)

**参考链接：**

- [维基百科：堆（Heap）](https://en.wikipedia.org/wiki/Heap_(data_structure))
- [堆的Java实现代码](https://shimo.im/docs/Lw86vJzOGOMpWZz2/read)

## 图

面试中并不多，如岛屿和陆地个数等。

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- | ------ |
[1. 两数之和](https://leetcode-cn.com/problems/two-sum/description/)|[1.twoSum.go](1.twoSum.go)|S|hashmap、diff递减|
[40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)|[40.getLeastNumbers.go](40.getLeastNumbers.go)|S|大根堆|
[剑指Offer49.丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)|[49.nthUglyNumber.go](49.nthUglyNumber.go)|M|DP、小根堆|
[59. 滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)|[59.maxSlidingWindow.go](59.maxSlidingWindow.go)|H|大顶堆、双堆队列|
[84. 柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)|[84.largestRectangleArea.go](84.largestRectangleArea.go)|H|暴力枚举、单调栈|
[242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/description/)|[242.isAnagram.go](242.isAnagram.go)|S|-|
[347. 前K个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)|[347.topKFrequent.go](347.topKFrequent.go)|M|hashmap计数、小根堆|
[N 叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/)|[589.preorder.go](589.preorder.go)|S|树的遍历|
