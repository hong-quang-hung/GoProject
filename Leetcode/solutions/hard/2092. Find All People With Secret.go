package hard

import (
	"fmt"
	"math"
	"sort"
)

// Reference: https://leetcode.com/problems/find-all-people-with-secret/
func init() {
	Solutions[2092] = func() {
		fmt.Println(`Input: n = 6, meetings = [[1,2,5],[2,3,8],[1,5,10]], firstPerson = 1`)
		fmt.Println(`Output:`, findAllPeople(6, S2SoSliceInt("[[1,2,5],[2,3,8],[1,5,10]]"), 1))
		fmt.Println(`Input: n = 4, meetings = [[3,1,3],[1,2,2],[0,3,3]], firstPerson = 3`)
		fmt.Println(`Output:`, findAllPeople(4, S2SoSliceInt("[[3,1,3],[1,2,2],[0,3,3]]"), 3))
	}
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	prevMeeting := 0

	dsu := NewUnionFind(n)
	dsu.Rank[0] = math.MaxInt
	dsu.UnionSet(0, firstPerson)

	people := make(map[int]bool)
	for _, meeting := range meetings {
		if meeting[2] != prevMeeting {
			for k := range people {
				if dsu.Find(k) != 0 {
					dsu.Reset(k)
				}
			}
			prevMeeting = meeting[2]
			people = make(map[int]bool)
		}

		people[meeting[0]] = true
		people[meeting[1]] = true
		dsu.UnionSet(meeting[0], meeting[1])
	}

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if dsu.Find(i) == 0 {
			res = append(res, i)
		}
	}
	return res
}
