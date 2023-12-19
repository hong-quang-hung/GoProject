package easy

import "fmt"

// Reference: https://leetcode.com/problems/k-items-with-the-maximum-sum/
func init() {
	Solutions[2600] = func() {
		fmt.Println(`Input: numOnes = 6, numZeros = 6, numNegOnes = 6, k = 13`)
		fmt.Println(`Output:`, kItemsWithMaximumSum(6, 6, 6, 13))
	}
}

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if numOnes >= k {
		return k
	}
	if numZeros >= k-numOnes {
		return numOnes
	}
	return 2*numOnes - k + numZeros
}
