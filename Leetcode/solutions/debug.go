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
	solutions Solution = make(Solution)

	EASY   int
	MEDIUM int
	HARD   int
)

func init() {
	defer clear(easy.Solutions)
	defer clear(medium.Solutions)
	defer clear(hard.Solutions)

	EASY, MEDIUM, HARD = len(easy.Solutions), len(medium.Solutions), len(hard.Solutions)

	// Easy
	for index, function := range easy.Solutions {
		solutions[index] = function
	}

	// Medium
	for index, function := range medium.Solutions {
		solutions[index] = function
	}

	// Hard
	for index, function := range hard.Solutions {
		solutions[index] = function
	}
}

func Len() int {
	return len(solutions)
}

func Leetcode_Debug(problem int) {
	if Invoke, hasSolution := solutions[problem]; hasSolution {
		Invoke()
	} else {
		Leetcode_Empty(problem)
	}
}

func Leetcode_Check_Golang_Solution(problem int) bool {
	_, hasSolution := solutions[problem]
	return hasSolution
}

func Leetcode_Empty(problem int) {
	fmt.Printf("The problem %d hasn't been solved yet!\n", problem)
}
