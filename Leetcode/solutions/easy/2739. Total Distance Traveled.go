package easy

import "fmt"

func init() {
	Solutions[2739] = Leetcode_Distance_Traveled
}

// Reference: https://leetcode.com/problems/total-distance-traveled/
func Leetcode_Distance_Traveled() {
	fmt.Println("Input: mainTank = 5, additionalTank = 10")
	fmt.Println("Output:", distanceTraveled(5, 10))
	fmt.Println("Input: mainTank = 1, additionalTank = 2")
	fmt.Println("Output:", distanceTraveled(1, 2))
}

func distanceTraveled(mainTank int, additionalTank int) int {
	return (mainTank + min((mainTank-1)/4, additionalTank)) * 10
}
