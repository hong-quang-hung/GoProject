package hard

import (
	"fmt"
	"math/bits"
)

// Reference: https://leetcode.com/problems/maximum-number-of-achievable-transfer-requests/
func init() {
	Solutions[1601] = func() {
		fmt.Println(`Input: n = 5, requests = [[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]]`)
		fmt.Println(`Output:`, maximumRequests(5, S2SoSliceInt(`[[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]]`)))
		fmt.Println(`Input: n = 3, requests = [[0,0],[1,2],[2,1]]`)
		fmt.Println(`Output:`, maximumRequests(3, S2SoSliceInt(`[[0,0],[1,2],[2,1]]`)))
		fmt.Println(`Input: n = 4, requests = [[0,3],[3,1],[1,2],[2,0]]`)
		fmt.Println(`Output:`, maximumRequests(4, S2SoSliceInt(`[[0,3],[3,1],[1,2],[2,0]]`)))
	}
}

func maximumRequests(n int, requests [][]int) int {
	res := 0
	for i := 0; i < (1 << len(requests)); i++ {
		arr := make([]int, n)
		pos := len(requests) - 1
		bitCount := bits.OnesCount(uint(i))
		if bitCount <= res {
			continue
		}

		for j := i; j > 0; j >>= 1 {
			if (j & 1) == 1 {
				arr[requests[pos][0]]--
				arr[requests[pos][1]]++
			}
			pos--
		}

		flag := true
		for i := 0; i < n; i++ {
			if arr[i] != 0 {
				flag = false
				break
			}
		}

		if flag {
			res = bitCount
		}
	}
	return res
}
