package easy

import "fmt"

// Reference: https://leetcode.com/problems/calculate-delayed-arrival-time/
func init() {
	Solutions[2651] = func() {
		fmt.Println(`Input: arrivalTime = 15, delayedTime = 5 `)
		fmt.Println(`Output:`, findDelayedArrivalTime(15, 5))
		fmt.Println(`Input: arrivalTime = 13, delayedTime = 11`)
		fmt.Println(`Output:`, findDelayedArrivalTime(13, 11))
	}
}

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
