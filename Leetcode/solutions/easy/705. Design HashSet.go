package easy

import "fmt"

// Reference: https://leetcode.com/problems/design-hashset/
func Leetcode_Design_HashSet() {
	fmt.Println("Input:")
	fmt.Println("['MyHashSet', 'add', 'add', 'contains', 'contains', 'add', 'contains', 'remove', 'contains']")
	fmt.Println("[[], [1], [2], [1], [3], [2], [2], [2], [2]]")
	fmt.Println("Output:")
	myHashSet := MyHashSetConstructor()
	myHashSet.Add(1)
	myHashSet.Add(2)
	fmt.Println("myHashSet.Contains(1)", "-->", myHashSet.Contains(1))
	fmt.Println("myHashSet.Contains(3)", "-->", myHashSet.Contains(3))
	myHashSet.Add(2)
	fmt.Println("myHashSet.Contains(2)", "-->", myHashSet.Contains(2))
	myHashSet.Remove(2)
	fmt.Println("myHashSet.Contains(2)", "-->", myHashSet.Contains(2))
}

type MyHashSet struct {
	hash []bool
}

func MyHashSetConstructor() MyHashSet {
	return MyHashSet{make([]bool, 1000001)}
}

func (h *MyHashSet) Add(key int) {
	h.hash[key] = true
}

func (h *MyHashSet) Remove(key int) {
	h.hash[key] = false
}

func (h *MyHashSet) Contains(key int) bool {
	return h.hash[key]
}
