package week09

import "container/list"

// 基于go标准库 container/list 实现

type item struct {
	val int
	ele *list.Element
}

type LRUCache2 struct {
	capacity int
	lruMap   map[int]*item
	lruList  *list.List
}

func NewLRUCache2(capacity int) *LRUCache2 {
	return &LRUCache2{
		capacity: capacity,
		lruMap:   make(map[int]*item, capacity),
		lruList:  list.New(),
	}
}

func (lru *LRUCache2) Get(key int) int {
	if _, ok := lru.lruMap[key]; !ok {
		return -1
	}

	node := lru.lruMap[key]
	lru.lruList.MoveToFront(node.ele)
	return node.val
}

func (lru *LRUCache2) Put(key, value int) {
	if node, ok := lru.lruMap[key]; ok {
		node.val = value
		lru.lruList.MoveToFront(node.ele)
		return
	}

	// 头部插入
	ele := lru.lruList.PushFront(key)
	lru.lruMap[key] = &item{
		val: value,
		ele: ele,
	}

	// 容量超限
	for lru.lruList.Len() > lru.capacity {
		last := lru.lruList.Back()
		delete(lru.lruMap, last.Value.(int)) // list里存的是key->node
		lru.lruList.Remove(last)
	}
}
