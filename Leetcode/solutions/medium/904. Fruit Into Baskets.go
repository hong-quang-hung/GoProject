package medium

import "fmt"

// Reference: https://leetcode.com/problems/fruit-into-baskets/
func init() {
	Solutions[904] = func() {
		fmt.Println(`Input: fruits = [1,2,1]`)
		fmt.Println(`Output:`, totalFruit([]int{1, 2, 1}))
		fmt.Println(`Input: fruits = [0,1,2,2]`)
		fmt.Println(`Output:`, totalFruit([]int{0, 1, 2, 2}))
		fmt.Println(`Input: fruits = [1,2,3,2,2]`)
		fmt.Println(`Output:`, totalFruit([]int{1, 2, 3, 2, 2}))
	}
}

func totalFruit(fruits []int) int {
	m, left, res := make(map[int]int), 0, 0
	for right := 0; right < len(fruits); right++ {
		m[fruits[right]]++
		for len(m) > 2 {
			m[fruits[left]]--
			if m[fruits[left]] == 0 {
				delete(m, fruits[left])
			}
			left++
		}

		res = max(res, right-left+1)
	}
	return res
}
