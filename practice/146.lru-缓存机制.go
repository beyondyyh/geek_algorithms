package practice

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type DListNode struct {
	key, value int
	prev, next *DListNode
}

type LRUCache struct {
	size       int                // 当前元素个数
	capacity   int                // 初始容量
	cache      map[int]*DListNode // hashmap存当前的kv pairs
	head, tail *DListNode         // 虚拟头尾节点
}

func NewLRUCache(capacity int) LRUCache {
	head, tail := &DListNode{}, &DListNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DListNode, capacity),
		head:     head,
		tail:     tail,
	}
}

// Get 查询
// 1. key不存在，返回-1
// 2. key存在则返回value，并将节点移动到头部
func (lru *LRUCache) Get(key int) int {
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}

	lru.moveToHead(node)
	return node.value
}

// Put 放入
// 1. 放之前key不存在，则需要新生成一个节点插入头部、更新hash、且size++，此时存在2种case
//	a. 容量够用，直接返回即可
//	b. 容量超限，则需要删除尾部节点、删除hash中对应的key、同时size--
// 2. 放之前key已经存在，则修改key对应节点的value，并且将其移动到头部，size不变
func (lru *LRUCache) Put(key, value int) {
	if _, ok := lru.cache[key]; !ok {
		node := &DListNode{key: key, value: value}
		lru.insertToHead(node)
		lru.cache[key] = node
		lru.size++
		if lru.size > lru.capacity { // case 1-b
			last := lru.tail.prev
			lru.removeNode(last)
			delete(lru.cache, last.key)
			lru.size--
		}
		return
	}

	// key存在，修改key对应节点的value、移动到头部
	node := lru.cache[key]
	node.value = value
	lru.moveToHead(node)
}

// moveToHead 将一个节点移动到头部
func (lru *LRUCache) moveToHead(node *DListNode) {
	lru.removeNode(node)
	lru.insertToHead(node)
}

// removeNode 删除一个节点
func (lru *LRUCache) removeNode(node *DListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// insertToHead 头部插入一个节点
func (lru *LRUCache) insertToHead(node *DListNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node // 3,4两步顺序不能反
	lru.head.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
