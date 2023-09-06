package easy

import "fmt"

// Reference: https://leetcode.com/problems/implement-stack-using-queues/
func init() {
	Solutions[225] = func() {
		fmt.Println("Input:")
		fmt.Println("['MyStack', 'push', 'push', 'top', 'pop', 'empty']")
		fmt.Println("[[], [1], [2], [], [], []]")
		fmt.Println("Output:")
		myStack := MyStackConstructor()
		myStack.Push(1)
		myStack.Push(2)
		fmt.Println("myStack.top()", "// return", myStack.Top())
		fmt.Println("myStack.pop()", "// return", myStack.Pop())
		fmt.Println("myStack.empty()", "// return", myStack.Empty())
	}
}

type MyStack struct {
	out Queue
}

func MyStackConstructor() MyStack {
	return MyStack{
		out: Queue{},
	}
}

func (s *MyStack) Push(x int) {
	tmp := Queue{x}
	for !s.out.Empty() {
		tmp.EnQueue(s.out.DeQueue())
	}
	s.out = tmp
}

func (s *MyStack) Pop() int {
	return s.out.DeQueue()
}

func (s *MyStack) Top() int {
	return s.out.Top()
}

func (s *MyStack) Empty() bool {
	return s.out.Empty()
}
