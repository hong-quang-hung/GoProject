package easy

import "fmt"

// Reference: https://leetcode.com/problems/implement-queue-using-stacks/
func init() {
	Solutions[232] = func() {
		fmt.Println(`Input:`)
		fmt.Println(`["MyQueue", "push", "push", "peek", "pop", "empty"]`)
		fmt.Println(`[null, null, null, 1, 1, false]`)
		fmt.Println(`Output:`)
		myQueue := MyQueueConstructor()
		fmt.Println(`myQueue.push(1);`)
		myQueue.Push(1)
		fmt.Println(`myQueue.push(2);`)
		myQueue.Push(2)
		fmt.Println(`myQueue.peek();`, `// return`, myQueue.Peek())
		fmt.Println(`myQueue.pop();`, `// return`, myQueue.Pop())
		fmt.Println(`myQueue.empty();`, `// return`, myQueue.Empty())
	}
}

type MyQueue struct {
	a []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{a: make([]int, 0)}
}

func (q *MyQueue) Push(x int) {
	q.a = append(q.a, x)
}

func (q *MyQueue) Pop() int {
	res := q.a[0]
	q.a = q.a[1:]
	return res
}

func (q *MyQueue) Peek() int {
	return q.a[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.a) == 0
}
