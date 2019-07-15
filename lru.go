package lru

import (
	"container/list"
)

// LRUCache indicates a LRUCache instance
type LRUCache struct {
	Capacity int
	L        *list.List
	EMap     map[int]*list.Element
}

// kvPair indicates <key, value> as the Value of a *list.Element
type kvPair struct {
	Key   int
	Value int
}

// NewLRUCache gets a new LRUCache instance 
func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		Capacity: capacity,
		L:        list.New(),
		EMap:     make(map[int]*list.Element, capacity),
	}

	return lru
}

// Length returns the existing items count
func (lru *LRUCache) Length() int {
	return lru.L.Len()
}

// Get return the value by key
func (lru *LRUCache) Get(key int) int {
	if v, ok := lru.EMap[key]; ok {
		lru.L.MoveToFront(v)
		return v.Value.(*kvPair).Value
	}
	return -1
}

// Put adds <key, value> into LRUCache
func (lru *LRUCache) Put(key int, value int) {
	if v, ok := lru.EMap[key]; ok {
		// Move to front when key already exists
		currentFront := lru.L.Front()
		lru.L.MoveToFront(v)
		lru.L.MoveToBack(currentFront)
	} else {
		// If already full, evicts the back element
		if lru.Length() >= lru.Capacity {
			back := lru.L.Back()
			delete(lru.EMap, back.Value.(*kvPair).Key)
			lru.L.Remove(back)
		}
		lru.L.PushFront(&kvPair{Key: key, Value: value})
		lru.EMap[key] = lru.L.Front()
	}
}
