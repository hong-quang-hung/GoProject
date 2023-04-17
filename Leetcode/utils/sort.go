package utils

func CountingSort(nums []int) {
	counts := make(map[int]int)
	var minVal int = nums[0]
	var maxVal int = nums[0]

	for _, num := range nums {
		if minVal > num {
			minVal = num
		}
		if maxVal < num {
			maxVal = num
		}
		if d, v := counts[num]; v {
			counts[num] = d + 1
		} else {
			counts[num] = 1
		}
	}

	index := 0
	for val := minVal; val <= maxVal; val++ {
		d, v := counts[val]
		for v && d > 0 {
			nums[index] = val
			index++
			counts[val]--
			d, v = counts[val]
		}
	}
}

func SortedFind(sorted []int, value int) int {
	low, high := 0, len(sorted)-1
	for low <= high {
		median := (low + high) / 2
		if sorted[median] < value {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	return low
}
