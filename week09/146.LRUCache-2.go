package week09

import "container/list"

// 基于go标准库 container/list 实现

type kvpair struct {
	key, value int
}

type LRUCache2 struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func NewLRUCache2(capacity int) *LRUCache2 {
	return &LRUCache2{
		capacity: capacity,
		cache:    make(map[int]*list.Element, capacity),
		list:     list.New(),
	}
}

func (lru *LRUCache2) Get(key int) int {
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}

	lru.list.MoveToFront(node)
	return node.Value.(*kvpair).value
}

// Put 放入
// 1. 放之前key不存在，则需要新生成一个节点插入头部、更新hash、且size++，此时存在2种case
//	a. 容量够用，直接返回即可
//	b. 容量超限，把尾节点移动到头部PushFront，同时删除hash中key对应节点
// 2. 放之前key已经存在，则修改key对应节点的value，并且将其移动到头部
func (lru *LRUCache2) Put(key, value int) {
	if _, ok := lru.cache[key]; !ok {
		item := &kvpair{key: key, value: value}
		node := lru.list.PushFront(item)
		lru.cache[key] = node
		if lru.list.Len() > lru.capacity {
			last := lru.list.Back()
			lru.list.Remove(last)
			delete(lru.cache, last.Value.(*kvpair).key)
		}
		return
	}

	node := lru.cache[key]
	node.Value.(*kvpair).value = value
	lru.list.MoveToFront(node)
}
