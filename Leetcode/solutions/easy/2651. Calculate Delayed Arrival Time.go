package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/calculate-delayed-arrival-time/
func Leetcode_Find_Delayed_Arrival_Time() {
	fmt.Println("arrivalTime = 15, delayedTime = 5 ]")
	fmt.Println("Output:", findDelayedArrivalTime(15, 5))
	fmt.Println("Input: arrivalTime = 13, delayedTime = 11")
	fmt.Println("Output:", findDelayedArrivalTime(13, 11))
}

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
