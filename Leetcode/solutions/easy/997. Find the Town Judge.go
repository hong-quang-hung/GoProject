package easy

import "fmt"

// Reference: https://leetcode.com/problems/find-the-town-judge/
func init() {
	Solutions[997] = func() {
		fmt.Println(`Input: n = 2, trust = [[1,2]]`)
		fmt.Println(`Output:`, findJudge(2, S2SoSliceInt("[[1,2]]")))
		fmt.Println(`Input: n = 3, trust = [[1,3],[2,3]]`)
		fmt.Println(`Output:`, findJudge(3, S2SoSliceInt("[[1,3],[2,3]]")))
		fmt.Println(`Input: n = 3, trust = [[1,3],[2,3],[3,1]]`)
		fmt.Println(`Output:`, findJudge(3, S2SoSliceInt("[[1,3],[2,3],[3,1]]")))
	}
}

func findJudge(n int, trust [][]int) int {
	trusted := make([]int, n+1)
	trusts := make([]int, n+1)
	for i := range trust {
		trusted[trust[i][1]]++
		trusts[trust[i][0]]++
	}

	for i := 1; i <= n; i++ {
		if trusted[i] == n-1 && trusts[i] == 0 {
			return i
		}
	}
	return -1
}
