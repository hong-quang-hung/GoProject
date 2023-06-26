package types

// Min Heap for integer type
type MinHeap []int

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Empty() bool        { return len(h) == 0 }
func (h *MinHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *MinHeap) Peek() int {
	r := (*h)[0]
	return r
}
func (h *MinHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

// Max Heap for integer type
type MaxHeap []int

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Empty() bool        { return len(h) == 0 }
func (h *MaxHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *MaxHeap) Peek() int {
	r := (*h)[0]
	return r
}
func (h *MaxHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}
