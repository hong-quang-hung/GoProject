package medium

import "fmt"

// Reference: https://leetcode.com/problems/online-stock-span/
func init() {
	Solutions[901] = func() {
		fmt.Println("Input:")
		fmt.Println("['StockSpanner', 'next', 'next', 'next', 'next', 'next', 'next', 'next']")
		fmt.Println("[[], [100], [80], [60], [70], [60], [75], [85]]")
		fmt.Println("Output:")

		stockSpanner := StockSpannerConstructor()
		fmt.Println("stockSpanner.next(100);", "return", stockSpanner.Next(100))
		fmt.Println("stockSpanner.next(80); ", "return", stockSpanner.Next(80))
		fmt.Println("stockSpanner.next(60); ", "return", stockSpanner.Next(60))
		fmt.Println("stockSpanner.next(70); ", "return", stockSpanner.Next(70))
		fmt.Println("stockSpanner.next(60); ", "return", stockSpanner.Next(60))
		fmt.Println("stockSpanner.next(75); ", "return", stockSpanner.Next(75))
		fmt.Println("stockSpanner.next(85); ", "return", stockSpanner.Next(85))
	}
}

type StockSpanner struct {
	stack [][]int
}

func StockSpannerConstructor() StockSpanner {
	stock := new(StockSpanner)
	return *stock
}

func (s *StockSpanner) Next(price int) int {
	res := 1
	for len(s.stack) > 0 && s.stack[len(s.stack)-1][0] <= price {
		res += s.stack[len(s.stack)-1][1]
		s.stack = s.stack[:len(s.stack)-1]
	}

	s.stack = append(s.stack, []int{price, res})
	return res
}
