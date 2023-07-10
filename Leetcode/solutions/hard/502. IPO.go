package hard

import (
	"container/heap"
	"fmt"
	"sort"
)

func init() {
	Solutions[502] = Leetcode_Find_Maximized_Capital
}

// Reference: https://leetcode.com/problems/ipo/
func Leetcode_Find_Maximized_Capital() {
	fmt.Println("Input: k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]")
	fmt.Println("Output:", findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
	fmt.Println("Input: k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]")
	fmt.Println("Output:", findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}))
	fmt.Println("Input: k = 3, w = 0, profits = [1,3,2,2], capital = [0,2,1,4]")
	fmt.Println("Output:", findMaximizedCapital(3, 0, []int{1, 3, 2, 2}, []int{0, 2, 1, 4}))
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	item := make([][2]int, n)
	for i := 0; i < n; i++ {
		item[i] = [2]int{capital[i], profits[i]}
	}

	sort.Slice(item, func(i, j int) bool {
		return item[i][0] < item[j][0]
	})

	i, h := 0, new(MaxHeap)
	for k > 0 {
		for i < n && w >= item[i][0] {
			heap.Push(h, item[i][1])
			i++
		}

		if h.Len() == 0 {
			break
		}

		w += heap.Pop(h).(int)
		k--
	}
	return w
}
