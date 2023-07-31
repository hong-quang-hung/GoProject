package medium

import (
	"container/heap"
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/
func init() {
	Solutions[1353] = func() {
		fmt.Println("Input: events = [[1,2],[2,3],[3,4]]")
		fmt.Println("Output:", maxEvents(S2SoSliceInt("[[1,2],[2,3],[3,4]]")))
		fmt.Println("Input: events = [[1,2],[2,3],[3,4],[1,2]]")
		fmt.Println("Output:", maxEvents(S2SoSliceInt("[[1,2],[2,3],[3,4],[1,2]]")))
		fmt.Println("Input: events = [[1,4],[4,4],[2,2],[3,4],[1,1]]")
		fmt.Println("Output:", maxEvents(S2SoSliceInt("[[1,4],[4,4],[2,2],[3,4],[1,1]]")))
	}
}

func maxEvents(events [][]int) int {
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	h := new(MinHeap)
	eventsIndex, res := 0, 0
	for currentDay := 1; currentDay <= 100000; currentDay++ {
		for !h.Empty() && h.Peek() < currentDay {
			heap.Pop(h)
		}

		for eventsIndex < len(events) && events[eventsIndex][0] == currentDay {
			heap.Push(h, events[eventsIndex][1])
			eventsIndex++
		}

		if !h.Empty() {
			heap.Pop(h)
			res++
		}
	}
	return res
}
