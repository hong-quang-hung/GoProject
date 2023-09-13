package hard

import "fmt"

// Reference: https://leetcode.com/problems/candy/
func init() {
	Solutions[135] = func() {
		fmt.Println("Input: ratings = [1,0,2]")
		fmt.Println("Output:", candy([]int{1, 0, 2}))
		fmt.Println("Input: ratings = [1,2,2]")
		fmt.Println("Output:", candy([]int{1, 2, 2}))
	}
}

func candy(ratings []int) int {
	res := 1
	up, down, peak := 0, 0, 0
	for i := 1; i < len(ratings); i++ {
		prev, curr := ratings[i-1], ratings[i]
		if prev < curr {
			up++
			down = 0
			peak = up
			res += up
		} else if prev == curr {
			up, down, peak = 0, 0, 0
		} else {
			up = 0
			down++
			res += down
			if peak >= down {
				res--
			}
		}
		res++
	}
	return res
}
