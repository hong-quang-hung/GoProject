package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/smallest-number-in-infinite-set/
func Leetcode_Smallest_Infinite_Set() {
	fmt.Println("Input:")
	fmt.Println("['SmallestInfiniteSet', 'addBack', 'popSmallest', 'popSmallest', 'popSmallest', 'addBack', 'popSmallest', 'popSmallest', 'popSmallest']")
	fmt.Println("[[], [2], [], [], [], [1], [], [], []]")
	fmt.Println("Output:")

	smallestInfiniteSet := SmallestInfiniteSetConstructor()
	smallestInfiniteSet.AddBack(2)
	fmt.Println("smallestInfiniteSet.popSmallest() :", smallestInfiniteSet.PopSmallest())
	fmt.Println("smallestInfiniteSet.popSmallest() :", smallestInfiniteSet.PopSmallest())
	fmt.Println("smallestInfiniteSet.popSmallest() :", smallestInfiniteSet.PopSmallest())
	smallestInfiniteSet.AddBack(1)
	fmt.Println("smallestInfiniteSet.popSmallest() :", smallestInfiniteSet.PopSmallest())
	fmt.Println("smallestInfiniteSet.popSmallest() :", smallestInfiniteSet.PopSmallest())
	fmt.Println("smallestInfiniteSet.popSmallest() :", smallestInfiniteSet.PopSmallest())
}

type smallestInfiniteHeap []int

func (h smallestInfiniteHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h smallestInfiniteHeap) Len() int           { return len(h) }
func (h smallestInfiniteHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h smallestInfiniteHeap) Empty() bool        { return len(h) == 0 }
func (h *smallestInfiniteHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *smallestInfiniteHeap) Peek() int {
	r := (*h)[(*h).Len()-1]
	return r
}
func (h *smallestInfiniteHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

type SmallestInfiniteSet struct {
	heap   *smallestInfiniteHeap
	remove map[int]bool
	curVal int
}

func (s *SmallestInfiniteSet) init() {
	s.remove = make(map[int]bool)
	s.heap = new(smallestInfiniteHeap)
	s.curVal = 1
}

func SmallestInfiniteSetConstructor() SmallestInfiniteSet {
	s := SmallestInfiniteSet{}
	s.init()
	return s
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	res := s.curVal
	if s.heap.Len() > 0 && s.heap.Peek() < res {
		res = heap.Pop(s.heap).(int)
	} else {
		s.curVal++
	}

	s.remove[res] = true
	return res
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if !s.remove[num] {
		return
	}

	s.remove[num] = false
	heap.Push(s.heap, num)
}
