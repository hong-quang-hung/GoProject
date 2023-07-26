package easy

import "fmt"

// Reference: https://leetcode.com/problems/min-cost-climbing-stairs/
func init() {
	Solutions[746] = func() {
		fmt.Println("Input: cost = [10,15,20]")
		fmt.Println("Output:", minCostClimbingStairs([]int{10, 15, 20}))
		fmt.Println("Input: cost = [1,100,1,1,1,100,1,1,100,1]")
		fmt.Println("Output:", minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
		fmt.Println("Input: cost = [0,0,1,1]")
		fmt.Println("Output:", minCostClimbingStairs([]int{0, 0, 1, 1}))
	}
}

func minCostClimbingStairs(cost []int) int {
	first := cost[0]
	second := cost[1]
	for i := 2; i < len(cost); i++ {
		if first > second {
			first, second = second, second+cost[i]
		} else {
			first, second = second, first+cost[i]
		}
	}

	if first < second {
		return first
	}
	return second
}
