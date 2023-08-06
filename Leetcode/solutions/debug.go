package solutions

import (
	"fmt"

	"leetcode.com/Leetcode/solutions/easy"
	"leetcode.com/Leetcode/solutions/hard"
	"leetcode.com/Leetcode/solutions/medium"
	"leetcode.com/Leetcode/types"
)

type Solution types.Solution

// problem th convert to problem + normalize
var (
	_SOLUTIONS_ Solution = make(Solution)
)

func init() {
	// Easy
	for index, function := range easy.Solutions {
		_SOLUTIONS_[index] = function
	}

	// Medium
	for index, function := range medium.Solutions {
		_SOLUTIONS_[index] = function
	}

	// Hard
	for index, function := range hard.Solutions {
		_SOLUTIONS_[index] = function
	}
}

func Len() int {
	return len(_SOLUTIONS_)
}

func Leetcode_Debug(problem int) {
	if Invoke, hasSolution := _SOLUTIONS_[problem]; hasSolution {
		Invoke()
	} else {
		Leetcode_Empty(problem)
	}
}

func Leetcode_Check_Golang_Solution(problem int) bool {
	_, hasSolution := _SOLUTIONS_[problem]
	return hasSolution
}

func Leetcode_Empty(problem int) {
	fmt.Printf("The problem %d hasn't been solved yet!\n", problem)
}
