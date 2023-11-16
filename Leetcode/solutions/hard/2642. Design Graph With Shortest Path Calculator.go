package hard

import "fmt"

// https://leetcode.com/problems/design-graph-with-shortest-path-calculator/
func init() {
	Solutions[2642] = func() {
		fmt.Println("Input:")
		fmt.Println("[\"Graph\", \"shortestPath\", \"shortestPath\", \"addEdge\", \"shortestPath\"]")
		fmt.Println("[[4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]], [3, 2], [0, 3], [[1, 3, 4]], [0, 3]]")
		fmt.Println("OutPut:")
		g := GraphConstructor(4, S2SoSliceInt("[[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]"))
		fmt.Println("g.shortestPath(3, 2)", "// return", g.ShortestPath(3, 2))
		fmt.Println("g.shortestPath(0, 3)", "// return", g.ShortestPath(0, 3))
		g.AddEdge([]int{1, 3, 4})
		fmt.Println("g.addEdge([1, 3, 4])", "// return", "null")
		fmt.Println("g.shortestPath(0, 3)", "// return", g.ShortestPath(0, 3))
	}
}

type Graph struct {
	g    [][][2]int
	h    *Heap[[2]int]
	n    int
	MAX  int
	dist []int
}

func GraphConstructor(n int, edges [][]int) Graph {
	g := Graph{
		g:    make([][][2]int, n),
		h:    NewHeap[[2]int](CompareEdges),
		n:    n,
		MAX:  n * int(1e6),
		dist: make([]int, n),
	}
	for _, e := range edges {
		g.AddEdge(e)
	}
	return g
}

func (g *Graph) AddEdge(e []int) {
	g.g[e[0]] = append(g.g[e[0]], [2]int{e[1], e[2]})
}

func (g *Graph) ShortestPath(s int, e int) int {
	for i := 0; i < g.n; i++ {
		g.dist[i] = g.MAX
	}
	g.dist[s] = 0
	g.h.Clear()
	g.h.Push([2]int{0, s})
	for !g.h.Empty() {
		x := g.h.Top()
		g.h.Pop()
		if g.dist[x[1]] < x[0] {
			continue
		}
		for _, y := range g.g[x[1]] {
			if g.dist[y[0]] > g.dist[x[1]]+y[1] {
				g.dist[y[0]] = g.dist[x[1]] + y[1]
				g.h.Push([2]int{g.dist[y[0]], y[0]})
			}
		}
	}
	if g.dist[e] == g.MAX {
		return -1
	}
	return g.dist[e]
}

type Heap[T any] struct {
	a       []T
	sz      int
	compare func(a, b T) int
}

func NewHeap[T any](comparator func(a, b T) int) *Heap[T] {
	return &Heap[T]{
		a:       make([]T, 0),
		sz:      0,
		compare: comparator,
	}
}

func (h *Heap[T]) Push(x T) {
	if len(h.a) == h.sz {
		h.a = append(h.a, x)
	} else {
		h.a[h.sz] = x
	}
	h.sz++
	h.up(h.sz - 1)
}

func (h *Heap[T]) Pop() {
	h.sz--
	h.a[0] = h.a[h.sz]
	h.down(0)
}

func (h *Heap[T]) Empty() bool {
	return h.sz == 0
}

func (h *Heap[T]) Clear() {
	h.sz = 0
}

func (h *Heap[T]) Top() T {
	return h.a[0]
}

func (h *Heap[T]) up(i int) {
	for i != 0 && h.compare(h.a[i], h.a[(i-1)>>1]) == -1 {
		h.a[i], h.a[(i-1)>>1] = h.a[(i-1)>>1], h.a[i]
		i = (i - 1) >> 1
	}
}

func (h *Heap[T]) down(i int) {
	for i<<1+1 < h.sz {
		j := i<<1 + 1
		if j+1 < h.sz && h.compare(h.a[j], h.a[j+1]) == 1 {
			j++
		}
		if h.compare(h.a[i], h.a[j]) != 1 {
			break
		}
		h.a[i], h.a[j] = h.a[j], h.a[i]
		i = j
	}
}

func (h *Heap[T]) Print() {
	fmt.Println("Size:", h.sz)
	fmt.Println("Elem:", h.a[:h.sz])
}

func CompareInts(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func CompareEdges(a, b [2]int) int {
	return CompareInts(a[0], b[0])
}
