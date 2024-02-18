package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
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
)

func init() {
	problemDebug = 1458
	problemDebug = 2050
	problemDebug = 1425
	problemDebug = 1743
	problemDebug = 2642
	problemDebug = 815
	problemDebug = 210
	problemDebug = 909
	problemDebug = 907
	problemDebug = 1239
	problemDebug = 1074
	problemDebug = 1043
	problemDebug = 474
	problemDebug = 368
	problemDebug = 1247
	problemDebug = 2402

	problemTotal = 3045
}

func main() {
	defer LeetcodeHandleError()

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
	fmt.Printf("Leetcode Debug End....\n")
}

func LeetcodeHandleError() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
		debug.PrintStack()
	}
}

func LeetcodeInformation() {
	Leetcode := new(types.Leetcode)
	Leetcode.SetTotal(problemTotal)

	rf, err := os.Open(solvedPath)
	if err != nil {
		panic(err)
	}

	wg := new(sync.WaitGroup)

	fs := bufio.NewScanner(rf)
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		wg.Add(1)
		go Leetcode.SetSolved(wg, utils.SSplitInt(fs.Text())...)
	}
	rf.Close()

	wg.Wait()

	fmt.Println("There are", Leetcode.Solved(), "/", Leetcode.Total(), "problem(s) has been solved in Leetcode.")

	// defer ShowHasNotSubmitted(Leetcode)
	// defer ShowSolutionDuplicate(Leetcode)
	// defer ShowSolutionHasNotNote(Leetcode)
}

func ShowHasNotSubmitted(L *types.Leetcode) {
	for i := 0; i < problemTotal; i++ {
		if L.IsSolved(i) && !solutions.Leetcode_Check_Golang_Solution(i+1) {
			fmt.Printf("%4d hasn't been submitted solution with Golang language.\n", i+1)
		}
	}
}

func ShowSolutionDuplicate(L *types.Leetcode) {
	for i := 0; i < problemTotal; i++ {
		times := L.Count(i)
		if L.Count(i) > 1 {
			fmt.Printf("%4d is duplicate %d time(s) in \"solved.txt\".\n", i+1, times)
		}
	}
}

func ShowSolutionHasNotNote(L *types.Leetcode) {
	j := 0
	solutions.Leetcode_Solutions_Loop(func(i int) {
		if !L.IsSolved(i - 1) {
			fmt.Printf("%4d hasn't been noted in \"solved.txt\".\n", i)
		}
		j++
	})
}
