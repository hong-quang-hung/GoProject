package hard

import "fmt"

// Reference: https://leetcode.com/problems/count-of-integers/
func init() {
	Solutions[2719] = func() {
		fmt.Println(`Input: num1 = '1', num2 = '12', min_num = 1, max_num = 8`)
		fmt.Println(`Output:`, count(`1`, `12`, 1, 8))
		fmt.Println(`Input: num1 = '1', num2 = '5', min_num = 1, max_num = 5`)
		fmt.Println(`Output:`, count(`1`, `5`, 1, 5))
	}
}

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	dp := make([][][]int, 23)
	for i := 0; i < 23; i++ {
		dp[i] = make([][]int, 401)
		for j := 0; j < 401; j++ {
			dp[i][j] = []int{-1, -1}
		}
	}

	upper := countSolved(dp, 0, 0, 1, num2, min_sum, max_sum)
	for i := 0; i < 23; i++ {
		dp[i] = make([][]int, 401)
		for j := 0; j < 401; j++ {
			dp[i][j] = []int{-1, -1}
		}
	}

	lower := countSolved(dp, 0, 0, 1, num1, min_sum, max_sum)
	sum := 0
	for _, c := range num1 {
		sum += int(c - '0')
	}

	if sum >= min_sum && sum <= max_sum {
		lower--
	}

	res := (upper - lower + mod) % mod
	return res
}

func countSolved(dp [][][]int, idx int, sum int, tight int, s string, min_sum int, max_sum int) int {
	if idx == len(s) {
		if sum >= min_sum && sum <= max_sum {
			return 1
		} else {
			return 0
		}
	}

	if dp[idx][sum][tight] != -1 {
		return dp[idx][sum][tight]
	}

	res := 0
	if tight == 1 {
		for i := 0; i <= int(s[idx]-'0'); i++ {
			if int(s[idx]-'0') == i {
				res = (res + countSolved(dp, idx+1, sum+i, 1, s, min_sum, max_sum)) % mod
			} else {
				res = (res + countSolved(dp, idx+1, sum+i, 0, s, min_sum, max_sum)) % mod
			}
		}
	} else {
		for i := 0; i <= 9; i++ {
			res = (res + countSolved(dp, idx+1, sum+i, 1, s, min_sum, max_sum)) % mod
		}
	}

	dp[idx][sum][tight] = res % mod
	return dp[idx][sum][tight]
}
