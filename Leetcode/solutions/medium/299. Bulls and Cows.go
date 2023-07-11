package medium

import "fmt"

// Reference: https://leetcode.com/problems/bulls-and-cows/
func init() {
	Solutions[299] = func() {
		fmt.Println("Input: secret = '1234', guess = '0111'")
		fmt.Println("Output:", getHint("1234", "0111"))
		fmt.Println("Input: secret = '1122', guess = '2211'")
		fmt.Println("Output:", getHint("1122", "2211"))
	}
}

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	m := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			m[secret[i]]++
			m[guess[i]]--
			if m[secret[i]] <= 0 {
				cows++
			}
			if m[guess[i]] >= 0 {
				cows++
			}
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
