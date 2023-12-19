package medium

import "fmt"

// Reference: https://leetcode.com/problems/keys-and-rooms/
func init() {
	Solutions[841] = func() {
		fmt.Println(`Input: rooms = [[1],[2],[3],[]]`)
		fmt.Println(`Output:`, canVisitAllRooms(S2SoSliceInt(`[[1],[2],[3],[]]`)))
		fmt.Println(`Input: rooms = [[1,3],[3,0,1],[2],[0]]`)
		fmt.Println(`Output:`, canVisitAllRooms(S2SoSliceInt(`[[1,3],[3,0,1],[2],[0]]`)))
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	stack, visited := []int{}, make([]bool, n)
	for _, key := range rooms[0] {
		if key != 0 {
			stack = append(stack, key)
		}
	}

	visited[0] = true
	count := 1
	for len(stack) > 0 {
		key := stack[0]
		stack = stack[1:]
		if visited[key] {
			continue
		}

		visited[key] = true
		count++
		for _, nextKey := range rooms[key] {
			if !visited[nextKey] {
				stack = append(stack, nextKey)
			}
		}
	}
	return count == n
}
