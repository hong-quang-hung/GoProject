package hard

import "fmt"

// Reference: https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/
func init() {
	Solutions[1203] = func() {
		fmt.Println("Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]")
		fmt.Println("Output:", sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, S2SoSliceInt("[[],[6],[5],[6],[3,6],[],[],[]]")))
		fmt.Println("Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]")
		fmt.Println("Output:", sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, S2SoSliceInt("[[],[6],[5],[6],[3],[],[4],[]]")))
	}
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	groups, inDegrees := make([][]int, n+m), make([]int, n+m)
	for i, g := range group {
		if g > -1 {
			g += n
			groups[g] = append(groups[g], i)
			inDegrees[i]++
		}
	}

	for i, ancestors := range beforeItems {
		gi := group[i]
		if gi == -1 {
			gi = i
		} else {
			gi += n
		}

		for _, ancestor := range ancestors {
			ga := group[ancestor]
			if ga == -1 {
				ga = ancestor
			} else {
				ga += n
			}

			if gi == ga {
				groups[ancestor] = append(groups[ancestor], i)
				inDegrees[i]++
			} else {
				groups[ga] = append(groups[ga], gi)
				inDegrees[gi]++
			}
		}
	}

	res := []int{}
	for i, d := range inDegrees {
		if d == 0 {
			sortItemsDFS(i, n, &res, &inDegrees, &groups)
		}
	}

	if len(res) != n {
		return nil
	}
	return res
}

func sortItemsDFS(i, n int, res, inDegrees *[]int, groups *[][]int) {
	if i < n {
		*res = append(*res, i)
	}

	(*inDegrees)[i] = -1
	for _, ch := range (*groups)[i] {
		if (*inDegrees)[ch]--; (*inDegrees)[ch] == 0 {
			sortItemsDFS(ch, n, res, inDegrees, groups)
		}
	}
}
