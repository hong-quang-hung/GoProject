package medium

import "fmt"

// Reference: https://leetcode.com/problems/detonate-the-maximum-bombs/
func init() {
	Solutions[2101] = func() {
		fmt.Println(`Input: bombs = [[2,1,3],[6,1,4]]`)
		fmt.Println(`Output:`, maximumDetonation(S2SoSliceInt(`[[2,1,3],[6,1,4]]`)))
		fmt.Println(`Input: bombs = [[1,1,5],[10,10,5]]`)
		fmt.Println(`Output:`, maximumDetonation(S2SoSliceInt(`[[1,1,5],[10,10,5]]`)))
		fmt.Println(`Input: bombs = [[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]]`)
		fmt.Println(`Output:`, maximumDetonation(S2SoSliceInt(`[[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]]`)))
	}
}

func maximumDetonation(bombs [][]int) int {
	direct := make([][]int, len(bombs))
	for i, v0 := range bombs {
		for j, v1 := range bombs {
			if i == j {
				continue
			}

			if (v0[0]-v1[0])*(v0[0]-v1[0])+(v0[1]-v1[1])*(v0[1]-v1[1]) <= v0[2]*v0[2] {
				direct[i] = append(direct[i], j)
			}
		}
	}

	res := 0
	for i := range bombs {
		linker := make(map[int]struct{})
		maximumDetonationDFS(direct, i, make([]bool, len(bombs)), linker)
		if len(linker) > res {
			res = len(linker)
		}
	}
	return res
}

func maximumDetonationDFS(direct [][]int, index int, visited []bool, linker map[int]struct{}) {
	if visited[index] {
		return
	}

	visited[index] = true
	linker[index] = struct{}{}
	for _, bomb := range direct[index] {
		maximumDetonationDFS(direct, bomb, visited, linker)
	}
}
