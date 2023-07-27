package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/maximum-running-time-of-n-computers/
func init() {
	Solutions[2141] = func() {
		fmt.Println("Input: n = 2, batteries = [3,3,3]")
		fmt.Println("Output:", maxRunTime(2, []int{3, 3, 3}))
		fmt.Println("Input: n = 2, batteries = [1,1,1,1]")
		fmt.Println("Output:", maxRunTime(2, []int{1, 1, 1, 1}))
		fmt.Println("Input: n = 3, batteries = [10,10,3,5]")
		fmt.Println("Output:", maxRunTime(3, []int{10, 10, 3, 5}))
	}
}

func maxRunTime(n int, batteries []int) int64 {
	sort.Slice(batteries, func(i, j int) bool { return batteries[i] > batteries[j] })
	res := int64(0)
	for {
		head, tail := batteries[:n], batteries[n:]
		for i := 0; i < n; i++ {
			if head[i] == 0 {
				return res
			} else {
				head[i]--
			}
		}

		idx1 := sort.Search(len(head), func(i int) bool { return head[i] < tail[0] })
		idx2 := sort.Search(len(tail), func(i int) bool { return tail[i] < head[n-1] })
		if idx1 == 0 && idx2 == len(tail) {
			batteries = append(tail, head...)
		} else if idx2 > 0 {
			if idx1 < n {
				mid := batteries[idx1 : n+idx2]
				sort.Slice(mid, func(i, j int) bool { return mid[i] > mid[j] })
			}
		}
		res++
	}
}
