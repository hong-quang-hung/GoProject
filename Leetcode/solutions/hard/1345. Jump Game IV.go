package hard

import (
	"fmt"
)

func init() {
	Solutions[1345] = Leetcode_Min_Jumps
}

// Reference: https://leetcode.com/problems/jump-game-iv/
func Leetcode_Min_Jumps() {
	fmt.Println("Input: arr = [100,-23,-23,404,100,23,23,23,3,404]")
	fmt.Println("Output:", minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
	fmt.Println("Input: arr = [7,6,9,6,9,6,9,7]")
	fmt.Println("Output:", minJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}))
}

func minJumps(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}

	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		if graph[arr[i]] == nil {
			graph[arr[i]] = []int{}
		}
		graph[arr[i]] = append(graph[arr[i]], i)
	}

	visited, queue, steps := make([]bool, n), []int{0}, 0
	visited[0] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == n-1 {
				return steps
			}

			if cur-1 >= 0 && !visited[cur-1] {
				queue = append(queue, cur-1)
				visited[cur-1] = true
			}

			if cur+1 < n && !visited[cur+1] {
				queue = append(queue, cur+1)
				visited[cur+1] = true
			}

			if v, c := graph[arr[cur]]; c {
				for _, next := range v {
					if !visited[next] {
						queue = append(queue, next)
						visited[next] = true
					}
				}
				delete(graph, arr[cur])
			}
		}
		steps++
	}
	return -1
}
