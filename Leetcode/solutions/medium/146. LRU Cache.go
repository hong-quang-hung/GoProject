package medium

import (
	"container/list"
	"fmt"
)

// Reference: https://leetcode.com/problems/lru-cache/
func init() {
	Solutions[146] = func() {
		fmt.Println("Input:")
		fmt.Println("['LRUCache', 'put', 'put', 'get', 'put', 'get', 'put', 'get', 'get', 'get']")
		fmt.Println("[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]")
		fmt.Println("Output:")

		lRUCache := LRUCacheConstructor(2)
		lRUCache.Put(1, 1)
		lRUCache.Put(2, 2)
		fmt.Println("lRUCache.Get(1)", "return", lRUCache.Get(1))
		lRUCache.Put(3, 3)
		fmt.Println("lRUCache.Get(2)", "return", lRUCache.Get(2))
		lRUCache.Put(4, 4)
		fmt.Println("lRUCache.Get(1)", "return", lRUCache.Get(1))
		fmt.Println("lRUCache.Get(3)", "return", lRUCache.Get(3))
		fmt.Println("lRUCache.Get(4)", "return", lRUCache.Get(4))
	}
}

type LRUCache struct {
	cache    map[int]*list.Element
	linklist *list.List
	capacity int
}

func LRUCacheConstructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*list.Element, capacity),
		linklist: list.New(),
		capacity: capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}

	elem := lru.cache[key]
	lru.linklist.MoveToFront(elem)
	return elem.Value.([]int)[1]
}

func (lru *LRUCache) Put(key int, value int) {
	if elem, ok := lru.cache[key]; ok {
		lru.linklist.Remove(elem)
		newelem := lru.linklist.PushFront([]int{key, value})
		lru.cache[key] = newelem
		return
	}

	if len(lru.cache) == lru.capacity {
		elem := lru.linklist.Back()
		v := lru.linklist.Remove(elem)
		delete(lru.cache, v.([]int)[0])
	}

	newelem := lru.linklist.PushFront([]int{key, value})
	lru.cache[key] = newelem
}
