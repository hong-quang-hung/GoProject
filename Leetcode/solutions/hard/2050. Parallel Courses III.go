package hard

import "fmt"

// Reference: https://leetcode.com/problems/parallel-courses-iii/
func init() {
	Solutions[2050] = func() {
		fmt.Println(`Input: n = 3, relations = [[1,3],[2,3]], time = [3,2,5]`)
		fmt.Println(`Output:`, minimumTime(3, S2SoSliceInt(`[[1,3],[2,3]]`), []int{1, 2, 3, 2}))
		fmt.Println(`Input: n = 5, relations = [[1,5],[2,5],[3,5],[3,4],[4,5]], time = [1,2,3,4,5]`)
		fmt.Println(`Output:`, minimumTime(5, S2SoSliceInt(`[[1,5],[2,5],[3,5],[3,4],[4,5]]`), []int{1, 2, 3, 4, 5}))
	}
}

func minimumTime(n int, relations [][]int, time []int) int {
	graph := make([][]int, n)
	inDegrees := make([]int, n)
	for _, r := range relations {
		graph[r[0]-1] = append(graph[r[0]-1], r[1]-1)
		inDegrees[r[1]-1]++
	}

	totalTime := make([]int, n)
	queue := []int{}
	for v := range graph {
		if inDegrees[v] == 0 {
			queue = append(queue, v)
			totalTime[v] = time[v]
		}
	}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, w := range graph[v] {
			totalTime[w] = max(totalTime[w], totalTime[v]+time[w])
			inDegrees[w]--
			if inDegrees[w] == 0 {
				queue = append(queue, w)
			}
		}
	}

	res := 0
	for _, t := range totalTime {
		res = max(res, t)
	}
	return res
}
