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
		for _, ch := range bank[0] {
			if ch == '1' {
				counter++
			}
		}
		return counter
	}

	count := oneCount(bank[0])
	res := count
	for i := 1; i < m; i++ {
		if count == 0 {
			return 0
		}

		count = oneCount(bank[i])
		res *= count
	}
	return res
}
