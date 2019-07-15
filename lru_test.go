package lru

import "testing"

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(4)

	lru.Put(1, 1)        // [1,]
	lru.Put(2, 2)        // [2, 1]
	if lru.Get(1) != 1 { // returns 1, then [1, 2]
		t.Fatalf("Expected: %d, got: %d", 1, lru.Get(1))
	}
	lru.Put(3, 3)        // [3, 1, 2]
	if lru.Get(2) != 2 { // returns 2, then [2, 3, 1]
		t.Fatalf("Expected: %d, got: %d", 2, lru.Get(2))
	}
	lru.Put(4, 4)        // [4, 2, 3, 1]
	if lru.Get(1) != 1 { // returns 1, then [1, 4, 2, 3]
		t.Fatalf("Expected: %d, got: %d", 1, lru.Get(1))
	}
	lru.Put(5, 5)         // evicts 3, then [5, 1, 4, 2]
	if lru.Get(3) != -1 { // return -1 (not found)
		t.Fatalf("Expected: %d, got: %d", -1, lru.Get(3))
	}
	if lru.Get(4) != 4 { // returns 4, then [4, 5, 1, 2]
		t.Fatalf("Expected: %d, got: %d", 4, lru.Get(4))
	}

	// To make sure the result as expected
	result := []int{4, 5, 1, 2}
	index := 0
	for e := lru.L.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*kvPair).Value
		if ev != result[index] {
			t.Fatalf("LRUCache: Expected: %d at index[%d], got: %d ", result[index], index, ev)
		}
		index++
	}

}
