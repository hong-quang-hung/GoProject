package solutions

import (
	"fmt"

	"leetcode.com/Leetcode/solutions/easy"
	"leetcode.com/Leetcode/solutions/hard"
	"leetcode.com/Leetcode/solutions/medium"
	"leetcode.com/Leetcode/types"
)

type Solution types.Solution

var (
	Solutions Solution = make(Solution)

	EASY   int
	MEDIUM int
	HARD   int

	Duplicate []int
)

func init() {
	defer clear(easy.Solutions)
	defer clear(medium.Solutions)
	defer clear(hard.Solutions)

	EASY, MEDIUM, HARD = len(easy.Solutions), len(medium.Solutions), len(hard.Solutions)

	// Easy
	Add(easy.Solutions)
	// Medium
	Add(medium.Solutions)
	// Hard
	Add(hard.Solutions)
}

func Add(solutions types.Solution) {
	for index, function := range solutions {
		Solutions[index] = function
	}
}

func Len() int {
	return len(Solutions)
}

func Leetcode_Debug(problem int) {
	if Invoke, hasSolution := Solutions[problem]; hasSolution {
		Invoke()
	} else {
		Leetcode_Empty(problem)
	}
}

func Leetcode_Check_Golang_Solution(problem int) bool {
	_, hasSolution := Solutions[problem]
	return hasSolution
}

func Leetcode_Empty(problem int) {
	fmt.Printf("The problem %d hasn't been solved yet!\n", problem)
}

func Leetcode_Solutions_Loop(function func(i int)) {
	for i := range Solutions {
		function(i)
	}
}
