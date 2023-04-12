package easy

import "fmt"

// Reference: https://leetcode.com/problems/can-place-flowers/
func Leetcode_Can_Place_Flowers() {
	fmt.Println("Input: flowerbed = [1,0,0,0,1], n = 1")
	fmt.Println("Output:", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	for i := 0; i < size; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == size-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			i++
		}
	}
	return n <= 0
}
