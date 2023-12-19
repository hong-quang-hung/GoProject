package easy

import "fmt"

// Reference: https://leetcode.com/problems/distribute-money-to-maximum-children/
func init() {
	Solutions[2591] = func() {
		fmt.Println(`Input: money = 20, children = 3`)
		fmt.Println(`Output:`, distMoney(20, 3))
		fmt.Println(`Input: money = 17, children = 2`)
		fmt.Println(`Output:`, distMoney(17, 2))
	}
}

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}

	if money > children*8 {
		return children - 1
	}

	if money == 4 && children == 1 {
		return -1
	}

	d := money - children
	k, r := d/7, d%7
	if r == 3 && children-k == 1 {
		k--
	}
	return k
}
