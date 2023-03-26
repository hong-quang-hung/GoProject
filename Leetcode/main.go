package main

import (
	"fmt"

	"leetcode.com/Leetcode/types"

	"leetcode.com/Leetcode/solutions"
)

func main() {
	fmt.Println("..............................................................................................................................................")
	fmt.Println("Golang Leetcode...")
	solutions.Leetcode_Longest_Palindrome()
	fmt.Println("..............................................................................................................................................")
	GetRandomProblem()
	// solutions.Test_Pattern()
}

func GetRandomProblem() {
	Leetcode := &types.Leetcode{}
	Leetcode.SetTotalProblem(2063)
	Leetcode.SetSolvedProblems(1, 2, 3, 4, 5, 9, 13, 14, 15, 20, 21, 23)

	fmt.Println("There are", Leetcode.Solved(), "problems has been solved in Leetcode.")
	fmt.Println("Today, Number of Leetcode Problem is: ", Leetcode.GetRandomProblem())
}
