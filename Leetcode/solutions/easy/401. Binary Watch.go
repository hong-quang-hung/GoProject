package easy

import (
	"fmt"
	"math/bits"
)

// Reference: https://leetcode.com/problems/binary-watch/
func init() {
	Solutions[401] = func() {
		fmt.Println("Input: turnedOn = 1")
		fmt.Println("Output:", readBinaryWatch(1))
	}
}

func readBinaryWatch(turnedOn int) []string {
	var res []string
	for h := 0; h < 12; h++ {
		i := bits.OnesCount(uint(h))
		for m := 0; m < 60; m++ {
			j := bits.OnesCount(uint(m))
			if turnedOn == i+j {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return res
}
