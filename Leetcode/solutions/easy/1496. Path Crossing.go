package easy

import "fmt"

// Reference: https://leetcode.com/problems/path-crossing/
func init() {
	Solutions[1496] = func() {
		fmt.Println(`Input: path = "NES"`)
		fmt.Println(`Output:`, isPathCrossing("NES"))
		fmt.Println(`Input: path = "NNSWWEWSSESSWENNW"`)
		fmt.Println(`Output:`, isPathCrossing("NNSWWEWSSESSWENNW"))
	}
}

func isPathCrossing(path string) bool {
	key := [2]int{0, 0}
	m := make(map[[2]int]interface{})
	m[key] = struct{}{}
	for _, ch := range path {
		switch ch {
		case 'N':
			key[1]++
		case 'E':
			key[0]++
		case 'S':
			key[1]--
		case 'W':
			key[0]--
		}

		if _, ok := m[key]; ok {
			return true
		} else {
			m[key] = struct{}{}
		}
	}
	return false
}
