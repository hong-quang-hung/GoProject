package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital
func init() {
	Solutions[2477] = func() {
		fmt.Println(`Input: roads = [[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]], seats = 2`)
		fmt.Println(`Output:`, minimumFuelCost(S2SoSliceInt(`[[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]]`), 2))
	}
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	m := make(map[int][]int)
	fuel := int64(0)
	for _, road := range roads {
		if m[road[0]] == nil {
			m[road[0]] = []int{}
		}
		m[road[0]] = append(m[road[0]], road[1])

		if m[road[1]] == nil {
			m[road[1]] = []int{}
		}
		m[road[1]] = append(m[road[1]], road[0])
	}
	minimumFuelCostDFS(m, 0, -1, seats, &fuel)
	return fuel
}

func minimumFuelCostDFS(m map[int][]int, from, to, seats int, fuel *int64) int64 {
	if len(m) == 0 {
		return int64(0)
	}

	representatives := int64(1)
	for _, node := range m[from] {
		if node != to {
			representatives += minimumFuelCostDFS(m, node, from, seats, fuel)
		}
	}

	if from != 0 {
		*fuel += int64(math.Ceil(float64(representatives) / float64(seats)))
	}
	return representatives
}
