package hard

import (
	"fmt"
	"slices"
)

// Reference: https://leetcode.com/problems/meeting-rooms-iii/
func init() {
	Solutions[2402] = func() {
		fmt.Println(`Input: n = 2, meetings = [[0,10],[1,5],[2,7],[3,4]]`)
		fmt.Println(`Output:`, mostBooked(2, S2SoSliceInt("[[0,10],[1,5],[2,7],[3,4]]")))
		fmt.Println(`Input: n = 3, meetings = [[1,20],[2,10],[3,5],[4,9],[6,8]]`)
		fmt.Println(`Output:`, mostBooked(3, S2SoSliceInt("[[1,20],[2,10],[3,5],[4,9],[6,8]]")))
	}
}

func mostBooked(n int, meetings [][]int) int {
	m := len(meetings)
	rooms := make([]int, n)
	meetingsCount := make([]int, n)

	slices.SortFunc(meetings, func(i, j []int) int {
		return i[0] - j[0]
	})

	for i := 0; i < m; i++ {
		minEnd := 10000000000
		roomNum := -1
		start := meetings[i][0]
		for j := 0; j < n; j++ {
			rooms[j] = max(rooms[j], start)
			if rooms[j] < minEnd {
				minEnd = rooms[j]
				roomNum = j
			}
		}

		rooms[roomNum] += meetings[i][1] - meetings[i][0]
		meetingsCount[roomNum]++
	}

	maxMeetings := -1
	res := -1
	for i := 0; i < n; i++ {
		if meetingsCount[i] > maxMeetings {
			maxMeetings = meetingsCount[i]
			res = i
		}
	}
	return res
}
