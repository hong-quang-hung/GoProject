package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-garden-perimeter-to-collect-enough-apples/
func Leetcode_Minimum_Perimeter() {
	fmt.Println("Input: neededApples = 1000000000")
	fmt.Println("Output:", minimumPerimeter(1000000000))
}

func minimumPerimeter(neededApples int64) int64 {
	res, x := int64(0), int64(0)
	for res < neededApples {
		x++
		res += 12 * x * x
	}
	return x * 8
}
