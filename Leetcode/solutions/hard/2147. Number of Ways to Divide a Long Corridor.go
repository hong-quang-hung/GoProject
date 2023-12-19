package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/
func init() {
	Solutions[2147] = func() {
		fmt.Println(`Input: corridor = "SSPPSPS"`)
		fmt.Println(`Output:`, numberOfWays(`SSPPSPS`))
		fmt.Println(`Input: corridor = "PPSPSP"`)
		fmt.Println(`Output:`, numberOfWays(`PPSPSP`))
		fmt.Println(`Input: corridor = "SSS"`)
		fmt.Println(`Output:`, numberOfWays(`SSS`))
	}
}

func numberOfWays(corridor string) int {
	left := len(corridor) - 1
	for left >= 0 && corridor[left] != 'S' {
		left--
	}

	if left == 0 {
		return -1
	}

	res := 1
	counter := 0
	for i := 0; i <= left; i++ {
		if corridor[i] == 'S' {
			counter++
		}

		if counter == 2 {
			j := i + 1
			for j <= left && corridor[j] != 'S' {
				j++
			}

			res = (res * (j - i)) % mod
			counter = 0
		}
	}

	if counter != 0 {
		return 0
	}
	return res
}
