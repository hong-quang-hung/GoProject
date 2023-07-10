package easy

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[933] = Leetcode_Number_Of_Recent_Calls
}

// Reference: https://leetcode.com/problems/number-of-recent-calls/
func Leetcode_Number_Of_Recent_Calls() {
	fmt.Println("Input:")
	fmt.Println("['RecentCounter', 'ping', 'ping', 'ping', 'ping']")
	fmt.Println("[[], [1], [100], [3001], [3002]]")
	fmt.Println("Output:")

	recentCounter := RecentCounterConstructor()
	fmt.Println("recentCounter.ping(1)", "-->", recentCounter.Ping(1))
	fmt.Println("recentCounter.ping(100)", "-->", recentCounter.Ping(100))
	fmt.Println("recentCounter.ping(3001)", "-->", recentCounter.Ping(3001))
	fmt.Println("recentCounter.ping(3002)", "-->", recentCounter.Ping(3002))
}

type RecentCounter struct {
	arr []int
}

func RecentCounterConstructor() RecentCounter {
	return RecentCounter{arr: []int{}}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.arr = append(rc.arr, t)
	if t > 3000 {
		rc.arr = rc.arr[sort.SearchInts(rc.arr, t-3000):]
	}
	return len(rc.arr)
}
