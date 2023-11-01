package medium

import (
	"fmt"
	"strconv"
)

// Reference: https://leetcode.com/problems/evaluate-reverse-polish-notation/
func init() {
	Solutions[150] = func() {
		fmt.Println("Input: tokens = ['2','1','+','3','*']")
		fmt.Println("Output:", evalRPN([]string{"2", "1", "+", "3", "*"}))
		fmt.Println("Input: tokens = ['4','13','5','/','+']")
		fmt.Println("Output:", evalRPN([]string{"4", "13", "5", "/", "+"}))
		fmt.Println("Input: tokens = ['10','6','9','3','+','-11','*','/','*','17','+','5','+']")
		fmt.Println("Output:", evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	}
}

func evalRPN(tokens []string) int {
	stack := []int{}
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, token := range tokens {
		if calculate, exist := operators[token]; exist {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], calculate(a, b))
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
