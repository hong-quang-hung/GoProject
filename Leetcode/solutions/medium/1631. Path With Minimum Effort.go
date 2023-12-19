package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/path-with-minimum-effort/
func init() {
	Solutions[1631] = func() {
		fmt.Println(`Input: heights = [[1,2,2],[3,8,2],[5,3,5]]`)
		fmt.Println(`Output:`, minimumEffortPath(S2SoSliceInt(`[[1,2,2],[3,8,2],[5,3,5]]`)))
		fmt.Println(`Input: heights = [[1,2,3],[3,8,4],[5,3,5]]`)
		fmt.Println(`Output:`, minimumEffortPath(S2SoSliceInt(`[[1,2,3],[3,8,4],[5,3,5]]`)))
	}
}

type minimumEffortPathHeap [][]int

func (h minimumEffortPathHeap) Len() int           { return len(h) }
func (h minimumEffortPathHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h minimumEffortPathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minimumEffortPathHeap) Empty() bool        { return len(h) == 0 }
func (h *minimumEffortPathHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *minimumEffortPathHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	h := new(minimumEffortPathHeap)
	visited := make([][]bool, len(heights))
	for i := range visited {
		visited[i] = make([]bool, len(heights[0]))
	}

	heap.Push(h, []int{0, 0, 0})
	for h.Len() > 0 {
		top := heap.Pop(h).([]int)
		if visited[top[0]][top[1]] {
			continue
		}

		visited[top[0]][top[1]] = true
		if top[0] == m-1 && top[1] == n-1 {
			return top[2]
		}

		for _, d := range BoardDirection {
			y, x := top[0]+d[0], top[1]+d[1]
			if x >= 0 && y >= 0 && x < n && y < m && !visited[y][x] {
				heap.Push(h, []int{y, x, max(top[2], abs(heights[y][x]-heights[top[0]][top[1]]))})
			}
		}
	}
	return 0
}
