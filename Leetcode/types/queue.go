package types

type Queue []int

func (q *Queue) EnQueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) DeQueue() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue) Top() int {
	return (*q)[0]
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}
