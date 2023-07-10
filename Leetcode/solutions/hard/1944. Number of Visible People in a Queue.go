package hard

import "fmt"

func init() {
	Solutions[1944] = Leetcode_Can_See_Persons_Count
}

// Reference: https://leetcode.com/problems/number-of-visible-people-in-a-queue/
func Leetcode_Can_See_Persons_Count() {
	fmt.Println("Input: heights = [10,6,8,5,11,9]")
	fmt.Println("Output:", canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
	fmt.Println("Input: heights = [5,1,2,3,10]")
	fmt.Println("Output:", canSeePersonsCount([]int{5, 1, 2, 3, 10}))
}

func canSeePersonsCount(heights []int) []int {
	monotonic := []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		count, height := 0, heights[i]
		for len(monotonic) > 0 && monotonic[len(monotonic)-1] < height {
			monotonic = monotonic[:len(monotonic)-1]
			count++
		}
		if len(monotonic) > 0 {
			count++
		}
		heights[i] = count
		monotonic = append(monotonic, height)
	}
	return heights
}
