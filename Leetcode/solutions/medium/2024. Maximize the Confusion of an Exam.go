package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximize-the-confusion-of-an-exam/
func init() {
	Solutions[2024] = func() {
		fmt.Println(`Input: answerKey = "TTFF", k = 2`)
		fmt.Println(`Output:`, maxConsecutiveAnswers(`TTFF`, 2))
		fmt.Println(`Input: answerKey = "TTFTTFTT", k = 1`)
		fmt.Println(`Output:`, maxConsecutiveAnswers(`TTFTTFTT`, 1))
	}
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	rF, lF, kF := 0, 0, k
	rT, lT, kT := 0, 0, k
	for rF < len(answerKey) {
		if answerKey[rF] == 'F' {
			kF--
		}

		if answerKey[rT] == 'T' {
			kT--
		}

		if kF < 0 {
			if answerKey[lF] == 'F' {
				kF++
			}
			lF++
		}

		if kT < 0 {
			if answerKey[lT] == 'T' {
				kT++
			}
			lT++
		}

		rF++
		rT++
	}
	return max(rT-lT, rF-lF)
}
