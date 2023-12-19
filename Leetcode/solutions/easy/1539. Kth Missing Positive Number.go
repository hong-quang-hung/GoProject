package easy

import "fmt"

// Reference: https://leetcode.com/problems/kth-missing-positive-number/
func init() {
	Solutions[1539] = func() {
		fmt.Println(`Input: arr = [2,3,4,7,11], k = 5`)
		fmt.Println(`Output:`, findKthPositive([]int{2, 3, 4, 7, 11}, 5))
		fmt.Println(`Input: arr = [1,2,3,4], k = 2`)
		fmt.Println(`Output:`, findKthPositive([]int{1, 2, 3, 4}, 2))
	}
}

func findKthPositive(arr []int, k int) int {
	cur, res, n := 0, 0, len(arr)
	for cur < n && k > 0 {
		jumb := arr[cur] - res - 1
		if jumb >= k {
			return res + k
		}
		res = arr[cur]
		cur++
		k = k - jumb
	}
	return res + k
}
