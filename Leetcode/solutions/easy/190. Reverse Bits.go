package easy

import (
	"fmt"
	"strconv"
)

// Reference: https://leetcode.com/problems/reverse-bits/
func init() {
	Solutions[190] = func() {
		u, _ := strconv.ParseUint("00000010100101000001111010011100", 2, 32)
		fmt.Println("n = 00000010100101000001111010011100")
		fmt.Println("Output:", reverseBits(uint32(u)))
		u, _ = strconv.ParseUint("11111111111111111111111111111101", 2, 32)
		fmt.Println("n = 11111111111111111111111111111101")
		fmt.Println("Output:", reverseBits(uint32(u)))
	}
}

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res = (res << 1) + num%2
		num /= 2
	}
	return res
}
