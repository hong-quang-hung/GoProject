package medium

import "fmt"

// Reference: https://leetcode.com/problems/course-schedule/
func init() {
	Solutions[207] = func() {
		fmt.Println("Input: numCourses = 2, prerequisites = [[1,0]]")
		fmt.Println("Output:", canFinish(2, S2SoSliceInt("[[1,0]]")))
		fmt.Println("Input: numCourses = 2, prerequisites = [[1,0],[0,1]]")
		fmt.Println("Output:", canFinish(2, S2SoSliceInt("[[1,0],[0,1]]")))
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree, adjList := make([]int, numCourses), make(map[int][]int)
	for _, p := range prerequisites {
		indegree[p[1]]++
		adjList[p[0]] = append(adjList[p[0]], p[1])
	}

	q := make([]int, 0)
	for i := range indegree {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	cntVisited := 0
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		cntVisited++
		if cntVisited > numCourses {
			return false
		}

		for _, next := range adjList[node] {
			indegree[next]--
			if indegree[next] == 0 {
				q = append(q, next)
			}
		}
	}
	return cntVisited == numCourses
}
