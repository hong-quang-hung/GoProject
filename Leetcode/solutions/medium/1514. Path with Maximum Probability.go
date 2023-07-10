package medium

import (
	"container/heap"
	"fmt"
)

func init() {
	Solutions[1514] = Leetcode_Max_Probability
}

// Reference: https://leetcode.com/problems/path-with-maximum-probability/
func Leetcode_Max_Probability() {
	fmt.Println("Input:  n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2")
	fmt.Println("Output:", maxProbability(3, S2SoSliceInt("[[0,1],[1,2],[0,2]]"), []float64{0.5, 0.5, 0.2}, 0, 2))
	fmt.Println("Input:  n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2")
	fmt.Println("Output:", maxProbability(3, S2SoSliceInt("[[0,1],[1,2],[0,2]]"), []float64{0.5, 0.5, 0.3}, 0, 2))
	fmt.Println("Input:  n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2")
	fmt.Println("Output:", maxProbability(3, S2SoSliceInt("[[0,1]]"), []float64{0.5}, 0, 2))
}

type pair struct {
	to  int
	val float64
}

type maxProbabilityHeap []*pair

func (h maxProbabilityHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h maxProbabilityHeap) Len() int           { return len(h) }
func (h maxProbabilityHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxProbabilityHeap) Empty() bool        { return len(h) == 0 }
func (h *maxProbabilityHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *maxProbabilityHeap) Push(i interface{}) {
	*h = append(*h, i.(*pair))
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	g := make([][]pair, n)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], pair{to: edges[i][1], val: succProb[i]})
		g[edges[i][1]] = append(g[edges[i][1]], pair{to: edges[i][0], val: succProb[i]})
	}

	h := new(maxProbabilityHeap)
	visited := make([]bool, n)
	dist := make([]float64, n)

	heap.Push(h, &pair{to: start, val: 0})
	dist[start] = 1
	for !h.Empty() {
		cur := heap.Pop(h).(*pair)
		if cur.to == end {
			return dist[end]
		}

		if visited[cur.to] {
			continue
		}

		visited[cur.to] = true
		for _, edge := range g[cur.to] {
			if visited[edge.to] {
				continue
			}

			next := dist[cur.to] * edge.val
			if next > dist[edge.to] {
				dist[edge.to] = next
				heap.Push(h, &pair{edge.to, next})
			}
		}
	}
	return 0
}
