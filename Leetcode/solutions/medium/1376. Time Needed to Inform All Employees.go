package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/time-needed-to-inform-all-employees/
func init() {
	Solutions[1376] = func() {
		fmt.Println("Input: n = 1, headID = 0, manager = [-1], informTime = [0]")
		fmt.Println("Output:", numOfMinutes(1, 0, []int{-1}, []int{0}))
		fmt.Println("Input: n = 6, headID = 2, manager = [2,2,-1,2,2,2], informTime = [0,0,1,0,0,0]")
		fmt.Println("Output:", numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
	}
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adjList := make([][]int, n)
	for i := 0; i < n; i++ {
		adjList[i] = []int{}
	}

	for i := 0; i < n; i++ {
		if manager[i] != -1 {
			adjList[manager[i]] = append(adjList[manager[i]], i)
		}
	}

	maxTime := math.MinInt
	queue := [][2]int{{headID, 0}}

	for len(queue) != 0 {
		employeePair := queue[0]
		queue = queue[1:]
		parent := employeePair[0]
		time := employeePair[1]
		maxTime = max(maxTime, time)

		for _, adjacent := range adjList[parent] {
			queue = append(queue, [2]int{adjacent, time + informTime[parent]})
		}
	}
	return maxTime
}
