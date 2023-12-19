package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/seat-reservation-manager/
func init() {
	Solutions[1845] = func() {
		fmt.Println(`Input:`)
		fmt.Println(`["SeatManager", "reserve", "reserve", "unreserve", "reserve", "reserve", "reserve", "reserve", "unreserve"]`)
		fmt.Println(`[[5], [], [], [2], [], [], [], [], [5]]`)
		fmt.Println(`Output:`)

		seatManager := SeatManagerConstructor(5)
		fmt.Println(`seatManager.reserve();`, `//`, `return`, seatManager.Reserve())
		fmt.Println(`seatManager.reserve();`, `//`, `return`, seatManager.Reserve())
		fmt.Println(`seatManager.unreserve(2);`)
		seatManager.Unreserve(2)
		fmt.Println(`seatManager.reserve();`, `//`, `return`, seatManager.Reserve())
		fmt.Println(`seatManager.reserve();`, `//`, `return`, seatManager.Reserve())
		fmt.Println(`seatManager.reserve();`, `//`, `return`, seatManager.Reserve())
		fmt.Println(`seatManager.unreserve(5);`)
		seatManager.Unreserve(5)
	}
}

type SeatManager struct {
	available *MinHeap
	history   []int
}

func SeatManagerConstructor(n int) SeatManager {
	available := new(MinHeap)
	for i := 1; i <= n; i++ {
		heap.Push(available, i)
	}
	return SeatManager{
		available: available,
		history:   []int{},
	}
}

func (s *SeatManager) Reserve() int {
	seat := heap.Pop(s.available).(int)
	s.history = append(s.history, seat)
	return seat
}

func (s *SeatManager) Unreserve(seatNumber int) {
	heap.Push(s.available, seatNumber)
}
