package medium

import "fmt"

// Reference: https://leetcode.com/problems/gas-station/
func init() {
	Solutions[134] = func() {
		fmt.Println("Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]")
		fmt.Println("Output:", canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
		fmt.Println("Input: gas = [2,3,4], cost = [3,4,3]")
		fmt.Println("Output:", canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	}
}

func canCompleteCircuit(gas []int, cost []int) int {
	res, total, tank := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			res = i + 1
			total += tank
			tank = 0
		}
	}
	if total+tank < 0 {
		return -1
	}
	return res
}
