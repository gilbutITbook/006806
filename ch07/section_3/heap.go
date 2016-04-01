package main

import (
	"container/heap"
	"fmt"
)

type IntHeap struct {
	ints []int
}

func (h *IntHeap) Less(i, j int) bool {
	return h.ints[i] < h.ints[j]
}
func (h *IntHeap) Swap(i, j int) {
	h.ints[i], h.ints[j] = h.ints[j], h.ints[i]
}
func (h *IntHeap) Len() int {
	return len(h.ints)
}
func (h *IntHeap) Pop() interface{} {
	x := h.ints[h.Len()-1]
	h.ints = h.ints[:h.Len()-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	h.ints = append(h.ints, x.(int))
}
func main() {
	ints := &IntHeap{[]int{12, 34, 45, 1, 54, 89, 41, 92, 3}}
	heap.Init(ints)
	ints.Push(9)
	ints.Push(11)
	for ints.Len() > 0 {
		fmt.Printf("%v ", heap.Pop(ints))
	}
}
