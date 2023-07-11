package medium

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/count-and-say/
func init() {
	Solutions[38] = func() {
		fmt.Println("Input: n = 1")
		fmt.Println("Output:", countAndSay(1))
		fmt.Println("Input: n = 4")
		fmt.Println("Output:", countAndSay(4))
	}
}

func countAndSay(n int) string {
	res := new(strings.Builder)
	res.WriteString("1")
	if n < 2 {
		return res.String()
	}

	res.WriteString("1")
	for i := 2; i < n; i++ {
		ctr := 1
		temp := new(strings.Builder)
		for j := 1; j < len(res.String()); j++ {
			if res.String()[j] != res.String()[j-1] {
				temp.WriteByte(byte(ctr) + 48)
				temp.WriteByte(res.String()[j-1])
				ctr = 0
			}

			ctr++
			if j == len(res.String())-1 {
				temp.WriteByte(byte(ctr) + 48)
				temp.WriteByte(res.String()[j])
			}
		}
		res = temp
	}
	return res.String()
}
