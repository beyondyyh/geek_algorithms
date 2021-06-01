package week09

// 146. LRU 缓存机制
// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
// 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
// 示例：
// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
// @lc: https://leetcode-cn.com/problems/lru-cache/

// 时间复杂度：对于 Get 和 Put 都是 O(1)
// 空间复杂度：O(n) n为capcity

// Double-Linked-List node
type DListNode struct {
	key, value int
	prev, next *DListNode
}

type LRUCache struct {
	size       int                // 大小
	capacity   int                // 容量
	cache      map[int]*DListNode // hashmap存当前的kv pairs
	head, tail *DListNode         // 头尾哑结点
}

func NewLRUCache(capacity int) *LRUCache {
	head := &DListNode{key: -1, value: -1} // 头哑结点
	tail := &DListNode{key: -1, value: -1} // 尾哑结点
	head.next = tail
	tail.prev = head

	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DListNode, capacity),
		head:     head,
		tail:     tail,
	}
	return lru
}

// Get 查询
// 1. key不存在，返回-1
// 2. key存在则返回value，并将节点移动到头部
func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}

	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

// Put 放入
// 1. 放之前key不存在，则在头部插入一个新结点，此时需要看当前容量是否满了，有2种情况：
//	a. 容量够用，在头部插入之后，直接返回即可；
//	b. 容量满了，在头部插入之后，还需要删除尾部的结点，注意：删除节点后，需要同步删除hashmap中的key；
// 2. 放之前key已经存在，则修改key对应的value值，同时把该结点移动到头部
func (lru *LRUCache) Put(key, value int) {
	if _, ok := lru.cache[key]; !ok { // case 1
		node := &DListNode{key: key, value: value}
		lru.cache[key] = node
		lru.insertToHead(node)
		lru.size++
		if lru.size > lru.capacity { // case b
			last := lru.removeTail()
			delete(lru.cache, last.key)
			lru.size--
		}
	} else { // case 2
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

// moveToHead 将结点移动到头部，操作：先删除再插入到头部
func (lru *LRUCache) moveToHead(node *DListNode) {
	lru.removeNode(node)
	lru.insertToHead(node)
}

// removeNode 删除一个结点
func (lru *LRUCache) removeNode(node *DListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// insertToHead 插入到头部
func (lru *LRUCache) insertToHead(node *DListNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// removeTail 删除最后一个节点，并返回该节点
func (lru *LRUCache) removeTail() *DListNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}
