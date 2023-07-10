package medium

import "fmt"

func init() {
	Solutions[2145] = Leetcode_Number_Of_Arrays
}

// Reference: https://leetcode.com/problems/count-the-hidden-sequences/
func Leetcode_Number_Of_Arrays() {
	fmt.Println("Input: differences = [1,-3,4], lower = 1, upper = 6")
	fmt.Println("Output:", numberOfArrays([]int{1, -3, 4}, 1, 6))
	fmt.Println("Input: differences = [3,-4,5,1,-2], lower = -4, upper = 5")
	fmt.Println("Output:", numberOfArrays([]int{3, -4, 5, 1, -2}, -4, 5))
}

func numberOfArrays(differences []int, lower int, upper int) int {
	prefix := make([]int, len(differences))
	prefix[0] = differences[0]
	maxP, minP := prefix[0], prefix[0]
	for i := 1; i < len(differences); i++ {
		prefix[i] = prefix[i-1] + differences[i]
		maxP = max(maxP, prefix[i])
		minP = min(minP, prefix[i])
	}

	res := 0
	for i := lower; i <= upper; i++ {
		if maxP+i >= lower && maxP+i <= upper && minP+i >= lower && minP+i <= upper {
			res++
		}
	}
	return res
}
