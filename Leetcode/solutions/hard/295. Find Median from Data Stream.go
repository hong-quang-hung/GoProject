package hard

import "fmt"

// Reference: https://leetcode.com/problems/find-median-from-data-stream/
func init() {
	Solutions[295] = func() {
		fmt.Println("Input:")
		fmt.Println("['MedianFinder', 'addNum', 'addNum', 'findMedian', 'addNum', 'findMedian']")
		fmt.Println("[[], [1], [2], [], [3], []]")
		fmt.Println("Output:")

		medianFinder := MedianFinderConstructor()
		medianFinder.AddNum(1)
		medianFinder.AddNum(2)
		fmt.Println("medianFinder.findMedian();", "// return", medianFinder.FindMedian())
		medianFinder.AddNum(3)
		fmt.Println("medianFinder.findMedian();", "// return", medianFinder.FindMedian())
	}
}

type MedianFinder struct {
	sum    float64
	n      int
	stream map[int]bool
}

func MedianFinderConstructor() MedianFinder {
	return MedianFinder{}
}

func (m *MedianFinder) AddNum(num int) {
	m.sum += float64(num)
	m.n++
}

func (m *MedianFinder) FindMedian() float64 {
	return m.sum / float64(m.n)
}
