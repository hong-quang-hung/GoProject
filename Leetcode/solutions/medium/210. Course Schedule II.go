package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/course-schedule-ii/
func init() {
	Solutions[210] = func() {
		fmt.Println(`Input: numCourses = 2, prerequisites = [[1,0]]`)
		fmt.Println(`Output:`, findOrder(2, S2SoSliceInt(`[[1,0]]`)))
		fmt.Println(`Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]`)
		fmt.Println(`Output:`, findOrder(4, S2SoSliceInt(`[[1,0],[2,0],[3,1],[3,2]]`)))
	}
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
		indegree[p[0]]++
	}

	frontier := []int{}
	for i, v := range indegree {
		if v == 0 {
			frontier = append(frontier, i)
		}
	}

	res := []int{}
	for len(frontier) != 0 {
		cur := frontier[0]
		frontier = frontier[1:]
		res = append(res, cur)
		for _, v := range g[cur] {
			indegree[v]--
			if indegree[v] == 0 {
				frontier = append(frontier, v)
			}
		}
	}

	if len(res) == numCourses {
		return res
	}
	return []int{}
}
