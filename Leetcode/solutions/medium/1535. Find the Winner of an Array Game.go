package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-the-winner-of-an-array-game/
func init() {
	Solutions[1535] = func() {
		fmt.Println(`Input: arr = [2,1,3,5,4,6,7], k = 2`)
		fmt.Println(`Output:`, getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
		fmt.Println(`Input: arr = [3,2,1], k = 10`)
		fmt.Println(`Output:`, getWinner([]int{3, 2, 1}, 10))
	}
}

func getWinner(arr []int, k int) int {
	res := 0
	if arr[res] < arr[1] {
		arr[res] = arr[1]
	}

	arr[1] = 1
	for i := 2; i < len(arr); i++ {
		if arr[1] == k {
			break
		}

		if arr[res] > arr[i] {
			arr[1]++
		} else {
			res = i
			arr[1] = 1
		}
	}
	return arr[res]
}
