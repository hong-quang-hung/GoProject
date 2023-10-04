package easy

import "fmt"

// Reference: https://leetcode.com/problems/132-pattern/
func init() {
	Solutions[706] = func() {
		fmt.Println("Input:")
		fmt.Println("['MyHashMap', 'put', 'put', 'get', 'get', 'put', 'get', 'remove', 'get']")
		fmt.Println("[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]")
		fmt.Println("Output:")
		myHashMap := MyHashMapConstructor()
		fmt.Println("myHashMap.put(1, 1)")
		myHashMap.Put(1, 1)
		fmt.Println("myHashMap.put(2, 2)")
		myHashMap.Put(2, 2)
		fmt.Println("myHashMap.get(1)", "// return", myHashMap.Get(1))
		fmt.Println("myHashMap.get(3)", "// return", myHashMap.Get(3))
		fmt.Println("myHashMap.put(2, 1)")
		myHashMap.Put(2, 1)
		fmt.Println("myHashMap.get(2)", "// return", myHashMap.Get(2))
		fmt.Println("myHashMap.remove(2)")
		myHashMap.Remove(2)
		fmt.Println("myHashMap.get(2)", "// return", myHashMap.Get(2))
	}
}

type hashing func(n int) int

type MyHashMap struct {
	h1 hashing
	h2 hashing
	a  []int
	b  []int
	n  int
}

func MyHashMapConstructor() MyHashMap {
	a := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		a[i] = -1
	}

	return MyHashMap{
		h1: func(k int) int { return k % 10000 },
		h2: func(k int) int { return 1 + k%5 },
		a:  a,
		b:  make([]int, 10000),
		n:  10000,
	}
}

func (m *MyHashMap) Put(key int, value int) {
	for i := 0; i < len(m.a); i++ {
		k := (m.h1(key) + i*m.h2(key)) % m.n
		if m.a[k] == -1 || m.a[k] == key {
			m.a[k] = key
			m.b[k] = value
			break
		}
	}
}

func (m *MyHashMap) Get(key int) int {
	for i := 0; i < len(m.a); i++ {
		k := (m.h1(key) + i*m.h2(key)) % m.n
		if m.a[k] == key {
			return m.b[k]
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	for i := 0; i < len(m.a); i++ {
		k := (m.h1(key) + i*m.h2(key)) % m.n
		if m.a[k] == key {
			m.a[k] = -1
			m.b[k] = -1
			break
		}
	}
}
