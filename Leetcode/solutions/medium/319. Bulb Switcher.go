package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/bulb-switcher/
func init() {
	Solutions[319] = func() {
		fmt.Println(`Input: n = 4`)
		fmt.Println(`Output:`, bulbSwitch(4))
		fmt.Println(`Input: n = 6`)
		fmt.Println(`Output:`, bulbSwitch(6))
	}
}

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
