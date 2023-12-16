package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/course-schedule-ii/
func init() {
	Solutions[210] = func() {
		fmt.Println("Input: numCourses = 2, prerequisites = [[1,0]]")
		fmt.Println("Output:", findOrder(2, S2SoSliceInt("[[1,0]]")))
		fmt.Println("Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]")
		fmt.Println("Output:", findOrder(4, S2SoSliceInt("[[1,0],[2,0],[3,1],[3,2]]")))
	}
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}

	fmt.Println(g)

	res := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		res[i] = i
	}

	sort.Slice(res, func(i, j int) bool {
		return len(g[i]) > len(g[j])
	})

	return res
}
