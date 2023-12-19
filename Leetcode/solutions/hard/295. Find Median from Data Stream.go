package hard

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/find-median-from-data-stream/
func init() {
	Solutions[295] = func() {
		fmt.Println(`Input:`)
		fmt.Println(`["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]`)
		fmt.Println(`[[], [1], [2], [], [3], []]`)
		fmt.Println(`Output:`)

		medianFinder := MedianFinderConstructor()
		medianFinder.AddNum(1)
		medianFinder.AddNum(2)
		fmt.Println(`medianFinder.findMedian();`, `// return`, medianFinder.FindMedian())
		medianFinder.AddNum(3)
		fmt.Println(`medianFinder.findMedian();`, `// return`, medianFinder.FindMedian())
	}
}

type MedianFinder struct {
	small *MaxHeap
	large *MinHeap
}

func MedianFinderConstructor() MedianFinder {
	m := new(MedianFinder)
	m.small = new(MaxHeap)
	m.large = new(MinHeap)
	return *m
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.small, num)
	if m.small.Len() != 0 && m.large.Len() != 0 {
		if m.small.Peek() > m.large.Peek() {
			heap.Push(m.large, heap.Pop(m.small).(int))
		}
	}

	if m.small.Len() > m.large.Len()+1 {
		heap.Push(m.large, heap.Pop(m.small).(int))
	}

	if m.large.Len() > m.small.Len()+1 {
		heap.Push(m.small, heap.Pop(m.large).(int))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.small.Len() > m.large.Len() {
		return float64(m.small.Peek())
	}

	if m.small.Len() < m.large.Len() {
		return float64(m.large.Peek())
	}

	return (float64(m.small.Peek()) + float64(m.large.Peek())) / float64(2)
}
