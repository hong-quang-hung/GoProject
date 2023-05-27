package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/buy-two-chocolates/
func Leetcode_Buy_Choco() {
	fmt.Println("Input: prices = [1,2,2], money = 3")
	fmt.Println("Output:", buyChoco([]int{1, 2, 2}, 3))
	fmt.Println("Input: prices = [3,2,3], money = 3")
	fmt.Println("Output:", buyChoco([]int{3, 2, 3}, 3))
}

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	fmt.Println(prices)
	if prices[0]+prices[1] > money {
		return money
	} else {
		return money - (prices[0] + prices[1])
	}
}
