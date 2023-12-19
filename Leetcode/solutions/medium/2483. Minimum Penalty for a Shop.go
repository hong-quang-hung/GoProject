package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-penalty-for-a-shop/
func init() {
	Solutions[2483] = func() {
		fmt.Println(`Input: customers = "YYNY"`)
		fmt.Println(`Output:`, bestClosingTime(`YYNY`))
		fmt.Println(`Input: customers = "NNNNN"`)
		fmt.Println(`Output:`, bestClosingTime(`NNNNN`))
		fmt.Println(`Input: customers = "YYYY"`)
		fmt.Println(`Output:`, bestClosingTime(`YYYY`))
	}
}

func bestClosingTime(customers string) int {
	maxScore, score, bestHour := 0, 0, -1
	for i := range customers {
		if customers[i] == 'Y' {
			score++
		} else {
			score--
		}

		if score > maxScore {
			maxScore = score
			bestHour = i
		}
	}
	return bestHour + 1
}
