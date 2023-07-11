package easy

import (
	"fmt"
	"sort"
	"strconv"
)

// Reference: https://leetcode.com/problems/split-with-minimum-sum/
func init() {
	Solutions[2578] = func() {
		fmt.Println("Input: num = 4325")
		fmt.Println("Output:", splitNum(4325))
	}
}

func splitNum(num int) int {
	arr := make([]int, len(strconv.Itoa(num)))
	i, numA, numB := 0, 0, 0
	for num > 0 {
		arr[i] = num % 10
		num = num / 10
		i++
	}

	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			numA = numA*10 + arr[i]
		} else {
			numB = numB*10 + arr[i]
		}
	}
	return numA + numB
}
