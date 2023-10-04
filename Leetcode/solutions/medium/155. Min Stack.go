package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/min-stack/
func init() {
	Solutions[155] = func() {
		fmt.Println("Input:")
		fmt.Println("['MinStack','push','push','push','getMin','pop','top','getMin']")
		fmt.Println("[[],[-2],[0],[-3],[],[],[],[]]")
		fmt.Println("Output:")
		minStack := MinStackConstructor()
		fmt.Println("minStack.push(-2)")
		minStack.Push(-2)
		fmt.Println("minStack.push(0)")
		minStack.Push(0)
		fmt.Println("minStack.push(-3)")
		minStack.Push(-3)
		fmt.Println("minStack.getMin()", "// return", minStack.GetMin())
		fmt.Println("minStack.pop()")
		minStack.Pop()
		fmt.Println("minStack.pop()")
		minStack.Pop()
		fmt.Println("minStack.top()", "// return", minStack.Top())
		fmt.Println("minStack.getMin()", "// return", minStack.GetMin())
	}
}

type MinStack struct {
	stack [][2]int
}

func MinStackConstructor() MinStack {
	return MinStack{stack: [][2]int{}}
}

func (m *MinStack) Push(val int) {
	if len(m.stack) == 0 {
		m.stack = append(m.stack, [2]int{val, val})
	} else {
		top := m.stack[len(m.stack)-1]
		m.stack = append(m.stack, [2]int{val, min(top[1], val)})
	}
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1][0]
}

func (m *MinStack) GetMin() int {
	return m.stack[len(m.stack)-1][1]
}
