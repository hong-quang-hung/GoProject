package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-laser-beams-in-a-bank/
func init() {
	Solutions[2125] = func() {
		fmt.Println(`Input: bank = ["011001","000000","010100","001000"]`)
		fmt.Println(`Output:`, numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
		fmt.Println(`Input: bank = ["000","111","000"]`)
		fmt.Println(`Output:`, numberOfBeams([]string{"000", "111", "000"}))
	}
}

func numberOfBeams(bank []string) int {
	m := len(bank)
	oneCount := func(s string) int {
		counter := 0
		for _, ch := range s {
			if ch == '1' {
				counter++
			}
		}
		return counter
	}

	left, leftVal := 0, 0
	for left < m {
		leftVal = oneCount(bank[left])
		if leftVal != 0 {
			break
		}
		left++
	}

	res := 0
	right := left + 1
	for right < m {
		rightVal := oneCount(bank[right])
		if rightVal != 0 {
			res += rightVal * leftVal
			leftVal = rightVal
		}
		right++
	}
	return res
}
