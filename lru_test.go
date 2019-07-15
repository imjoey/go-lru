package lru

import "testing"

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	if lru.Get(1) != 1 { // returns 1
		t.FailNow()
	}
	lru.Put(3, 3)         // evicts key 2
	if lru.Get(2) != -1 { // returns -1 (not found)
		t.FailNow()
	}
	lru.Put(4, 4)         // evicts key 1
	if lru.Get(1) != -1 { // returns -1 (not found)
		t.FailNow()
	}
	if lru.Get(3) != 3 { // returns 3
		t.FailNow()
	}
	if lru.Get(4) != 4 { // returns 4
		t.FailNow()
	}
}

// LRUCache cache = new LRUCache( 2 /* capacity */ );

// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4
