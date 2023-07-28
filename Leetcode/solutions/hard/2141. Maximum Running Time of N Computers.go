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
	sort.Ints(batteries)

	sum := int64(0)
	for _, batterie := range batteries {
		sum += int64(batterie)
	}

	i := len(batteries) - 1
	for {
		res := sum / int64(n)
		if res >= int64(batteries[i]) {
			return res
		}

		sum -= int64(batteries[i])
		n--
		i--
	}
}
