package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-the-original-array-of-prefix-xor/
func init() {
	Solutions[2433] = func() {
		fmt.Println(`Input: pref = [5,2,0,3,1]`)
		fmt.Println(`Output:`, findArray([]int{5, 2, 0, 3, 1}))
		fmt.Println(`Input: pref = [13]`)
		fmt.Println(`Output:`, findArray([]int{13}))
	}
}

func findArray(pref []int) []int {
	prefix := pref[0]
	for i := 1; i < len(pref); i++ {
		temp := pref[i]
		pref[i] = prefix ^ pref[i]
		prefix = temp
	}
	return pref
}
