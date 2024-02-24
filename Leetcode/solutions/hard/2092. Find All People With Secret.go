package hard

import "fmt"

// Reference: https://leetcode.com/problems/find-all-people-with-secret/
func init() {
	Solutions[2092] = func() {
		fmt.Println(`Input: n = 6, meetings = [[1,2,5],[2,3,8],[1,5,10]], firstPerson = 1`)
		fmt.Println(`Output:`, findAllPeople(4, S2SoSliceInt("[[1,2,5],[2,3,8],[1,5,10]]"), 1))
		fmt.Println(`Input: n = 4, meetings = [[3,1,3],[1,2,2],[0,3,3]], firstPerson = 3`)
		fmt.Println(`Output:`, findAllPeople(4, S2SoSliceInt("[[3,1,3],[1,2,2],[0,3,3]]"), 3))
	}
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	return nil
}
