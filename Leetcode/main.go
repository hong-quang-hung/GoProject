package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"leetcode.com/Leetcode/solutions"
	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

const (
	solvedPath string = "data/solved.txt"
)

var (
	problemDebug int
	problemTotal int
	problemGroup int
)

func init() {
	problemDebug = 2366
	problemDebug = 92

	problemGroup = 10
	problemTotal = 2851
}

func main() {
	PrintLine()
	LeetcodeDebug()
	PrintLine()
	LeetcodeInformation()
	PrintLine()
}

func PrintLine() {
	fmt.Println("..............................................................................................................................................")
}

func LeetcodeDebug() {
	fmt.Printf("Leetcode Debug Start...\n")
	solutions.Leetcode_Debug(problemDebug)
	fmt.Printf("Leetcode Debug End. The question %d belongs to group %dth.\n", problemDebug, (problemDebug)%problemGroup)
}

func LeetcodeInformation() {
	Leetcode := new(types.Leetcode)
	Leetcode.SetTotal(problemTotal)

	rf, err := os.Open(solvedPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	wg := new(sync.WaitGroup)
	fsc := bufio.NewScanner(rf)
	fsc.Split(bufio.ScanLines)
	for fsc.Scan() {
		wg.Add(1)
		go Leetcode.SetSolved(wg, utils.SSplitInt(fsc.Text())...)
	}
	rf.Close()
	wg.Wait()

	// defer ShowHasNotSubmitted(Leetcode)

	fmt.Println("There are", Leetcode.Solved(), "/", Leetcode.Total(), "problem(s) has been solved in Leetcode.")
	// fmt.Println("Today, Number of Leetcode Problem is:", Leetcode.PickProblem())
}

func ShowHasNotSubmitted(L *types.Leetcode) {
	for i := 0; i < problemTotal; i++ {
		if L.IsSolved(i) && !solutions.Leetcode_Check_Golang_Solution(i+1) {
			fmt.Println(i+1, "hasn't submit solution with Golang language.")
		}
	}
}
