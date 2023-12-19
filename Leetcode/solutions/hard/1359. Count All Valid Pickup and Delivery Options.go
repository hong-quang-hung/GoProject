package hard

import "fmt"

// Reference: https://leetcode.com/problems/count-all-valid-pickup-and-delivery-options/
func init() {
	Solutions[1359] = func() {
		fmt.Println(`Input: n = 1`)
		fmt.Println(`Output:`, countOrders(1))
		fmt.Println(`Input: n = 2`)
		fmt.Println(`Output:`, countOrders(2))
		fmt.Println(`Input: n = 3`)
		fmt.Println(`Output:`, countOrders(3))
	}
}

func countOrders(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res = (res * (2*i - 1) * i) % mod
	}
	return res
}
