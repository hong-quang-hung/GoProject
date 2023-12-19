package medium

import "fmt"

// Reference: https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/
func init() {
	Solutions[1743] = func() {
		fmt.Println(`Input: adjacentPairs = [[2,1],[3,4],[3,2]]`)
		fmt.Println(`Output:`, restoreArray(S2SoSliceInt(`[[2,1],[3,4],[3,2]]`)))
		fmt.Println(`Input: adjacentPairs = [[4,-2],[1,4],[-3,1]]`)
		fmt.Println(`Output:`, restoreArray(S2SoSliceInt(`[[4,-2],[1,4],[-3,1]]`)))
		fmt.Println(`Input: adjacentPairs = [[100000,-100000]]`)
		fmt.Println(`Output:`, restoreArray(S2SoSliceInt(`[[100000,-100000]]`)))
	}
}

func restoreArray(adjacentPairs [][]int) []int {
	neighbours := make(map[int][]int, len(adjacentPairs)+1)
	for i := range adjacentPairs {
		neighbours[adjacentPairs[i][0]] = append(neighbours[adjacentPairs[i][0]], adjacentPairs[i][1])
		neighbours[adjacentPairs[i][1]] = append(neighbours[adjacentPairs[i][1]], adjacentPairs[i][0])
	}

	rootNode := 0
	for i, nodes := range neighbours {
		if len(nodes) == 1 {
			rootNode = i
			break
		}
	}

	prev := rootNode
	output := []int{rootNode}
	rootNode = neighbours[rootNode][0]
	for len(neighbours[rootNode]) == 2 {
		output = append(output, rootNode)
		if neighbours[rootNode][0] != prev {
			prev = rootNode
			rootNode = neighbours[rootNode][0]
		} else {
			prev = rootNode
			rootNode = neighbours[rootNode][1]
		}
	}
	return append(output, rootNode)
}
