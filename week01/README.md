# week01 数组、链表、跳表、栈、队列、优先队列、双端队列

## 数据结构

### 数组

- 内存中连续的空间
- 存相同类型的数据
- 可以通过下标快速查询数据，时间复杂度 `O(1)`
- 增删元素会改变内存空间的大小，需要重新分配内存空间，`O(n)`

**缺点：** 增删元素频繁时数组效率不高

### 链表

- 用零散的内存空间存储数据
- 每个数据元素都包含一个指向下一个数据元素的内存地址指针
- 链表非连续存储，查找数据只能遍历, 复杂度 `O(n)`
- 增删元素效率高，只需要修改指针 `O(1)`

**缺点：** 查找元素效率不高

### 跳表

- 由于链表查询效率不高，跳表在链表的基础上 `升维` 增加了索引
- 实现了可以二分查找的链表，查找效率 `O(logn)`

**缺点：** 查找效率提升带来了增删效率的降低，增删节点都需要更新索引，时间复杂度变成: `O(logn)`，空间复杂度：`O(n)`

**工业中应用：** redis的zset，LRU缓存

### 栈

- 线性表, 数据按照线性组织存放的，每个元素的前面只能有一个前驱数据元素，后面也只能有一个后继数据元素
- 数组和链表都是线性表，栈是在线性表基础上加上了操作限制条件: 删除时 `“后进先出”`
- 栈不需要随机访问，也不需要在中间添加、删除元素，所以可以使用数组实现，也可以使用链表实现

**工业中应用：** 用栈来管理程序运行过程中的方法调用，栈顶元素始终是当前正在执行的方法的工作区

### 队列

- 一种操作受限的线性表
- 栈是后进先出, 队列是 `先进先出`，跟排队一样

**工业中应用：**

- 优先队列，带优先级的queue，如任务调度系统，总是会先执行优先级高的任务
- 双端队列，更高级的用法

## week01学习总结

todos

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- | ------ |
[1. 两数之和](https://leetcode-cn.com/problems/two-sum/)|[1.twoSum.go](1.twoSum.go)|S|hashmap、diff递减|
[20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)|[20.isValid.go](20.isValid.go)|S|栈|
[21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)|[21.mergeTwoLists.go](21.mergeTwoLists.go)|S|递归、双指针|
[23. 合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/)|[23.mergeKLists.go](23.mergeKLists.go)|H|小根堆、优先队列|
[24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)|[24.swapPairs.go](24.swapPairs.go)|M|递归、迭代|
[26. 删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)|[26.removeDuplicates.go](26.removeDuplicates.go)|S|双指针|
[61. 旋转链表](https://leetcode-cn.com/problems/rotate-list/)|[61.rotateRight.go](61.rotateRight.go)|M|双指针、取模|
[66. 加一](https://leetcode-cn.com/problems/plus-one/)|[66.plusOne.go](66.plusOne)|S|x|
[88. 合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)|[88.merge.go](88.merge.go)|S|双指针|
[92. 反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii/)|[92.reverseBetween.go](92.reverseBetween.go)|M|双指针|
[155. 最小栈](https://leetcode-cn.com/problems/min-stack/)|[155.minStack.go](155.minStack.go)|S|辅助栈|
[189. 旋转数组](https://leetcode-cn.com/problems/rotate-array/)|[189.rotate.go](189.rotate.go)|S|取模、交换|
[283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)|[283.moveZeroes.go](283.moveZeroes.go)|S|双指针、快排思想|
[641. 设计双端循环队列](https://leetcode-cn.com/problems/design-circular-deque/)|[641.MyCircularDeque.go](641.MyCircularDeque.go)|M|双链表|
