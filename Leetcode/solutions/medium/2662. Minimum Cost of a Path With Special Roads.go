package medium

import (
	"container/heap"
	"fmt"
)

func init() {
	Solutions[2662] = Leetcode_Minimum_Cost
}

// Reference: https://leetcode.com/problems/minimum-cost-of-a-path-with-special-roads/
func Leetcode_Minimum_Cost() {
	fmt.Println("Input: start = [1,1], target = [4,5], specialRoads = [[1,2,3,3,2],[3,4,4,5,1]]")
	fmt.Println("Output:", minimumCost([]int{1, 1}, []int{4, 5}, S2SoSliceInt("[[1,2,3,3,2],[3,4,4,5,1]]")))
	fmt.Println("Input: start = [3,2], target = [5,7], specialRoads = [[3,2,3,4,4],[3,3,5,5,5],[3,4,5,6,6]]")
	fmt.Println("Output:", minimumCost([]int{3, 2}, []int{5, 7}, S2SoSliceInt("[[3,2,3,4,4],[3,3,5,5,5],[3,4,5,6,6]]")))
}

type minimumCostHeap [][]int

func (h minimumCostHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h minimumCostHeap) Len() int           { return len(h) }
func (h minimumCostHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minimumCostHeap) Empty() bool        { return len(h) == 0 }
func (h *minimumCostHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *minimumCostHeap) Push(i interface{}) {
	*h = append(*h, i.([]int))
}

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	m, h := make(map[[2]int]int), new(minimumCostHeap)
	res := target[0] - start[0] + target[1] - start[1]
	heap.Push(h, []int{0, start[0], start[1]})
	for h.Len() > 0 {
		elem := heap.Pop(h).([]int)
		curDist, x, y := elem[0], elem[1], elem[2]
		for _, road := range specialRoads {
			pair := [2]int{road[2], road[3]}
			if _, ok := m[pair]; !ok {
				m[pair] = road[2] - start[0] + road[3] - start[1]
			}

			d := abs(road[0]-x) + abs(road[1]-y) + curDist + road[4]
			if m[pair] > d {
				m[pair] = d
				heap.Push(h, []int{m[pair], road[2], road[3]})
			}
		}
	}

	for _, road := range specialRoads {
		res = min(res, m[[2]int{road[2], road[3]}]+target[0]-road[2]+target[1]-road[3])
	}
	return res
}
