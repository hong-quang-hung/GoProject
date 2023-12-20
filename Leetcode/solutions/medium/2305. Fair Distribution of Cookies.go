package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/fair-distribution-of-cookies/
func init() {
	Solutions[2305] = func() {
		fmt.Println(`Input: cookies = [8,15,10,20,8], k = 2`)
		fmt.Println(`Output:`, distributeCookies([]int{8, 15, 10, 20, 8}, 2))
		fmt.Println(`Input: cookies = [6,1,3,2,2,4,1,2], k = 3`)
		fmt.Println(`Output:`, distributeCookies([]int{6, 1, 3, 2, 2, 4, 1, 2}, 3))
	}
}

func distributeCookies(cookies []int, k int) int {
	distribute := make([]int, k)
	return distributeCookiesDFS(0, k, k, distribute, cookies)
}

func distributeCookiesDFS(idx, k, zeroCount int, distribute, cookies []int) int {
	if zeroCount > len(cookies)-idx {
		return math.MaxInt32
	}

	if idx == len(cookies) {
		res := 0
		for _, num := range distribute {
			if num > res {
				res = num
			}
		}
		return res
	}

	res := math.MaxInt32
	for i := 0; i < k; i++ {
		if distribute[i] == 0 {
			zeroCount--
		}

		distribute[i] += cookies[idx]
		res = min(res, distributeCookiesDFS(idx+1, k, zeroCount, distribute, cookies))
		distribute[i] -= cookies[idx]
		if distribute[i] == 0 {
			zeroCount++
		}
	}
	return res
}
