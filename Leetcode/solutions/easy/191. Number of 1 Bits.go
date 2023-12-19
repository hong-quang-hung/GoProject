package easy

import (
	"fmt"
	"strconv"
)

// Reference: https://leetcode.com/problems/number-of-1-bits/
func init() {
	Solutions[191] = func() {
		u, _ := strconv.ParseUint(`00000000000000000000000000001011`, 2, 32)
		fmt.Println(`n = 00000000000000000000000000001011`)
		fmt.Println(`Output:`, hammingWeight(uint32(u)))
		u, _ = strconv.ParseUint(`00000000000000000000000010000000`, 2, 32)
		fmt.Println(`n = 00000000000000000000000010000000`)
		fmt.Println(`Output:`, hammingWeight(uint32(u)))
	}
}

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		res += int(num % 2)
		num /= 2
	}
	return res
}
