package medium

import (
	"container/heap"
	"fmt"

	"leetcode.com/Leetcode/types"
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

type SmallestInfiniteSet struct {
	heap   *types.MinHeap
	remove map[int]bool
	curVal int
}

func (s *SmallestInfiniteSet) init() {
	s.remove = make(map[int]bool)
	s.heap = new(types.MinHeap)
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
