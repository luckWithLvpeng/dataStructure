package heap

import (
	"container/heap"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := make(MaxHeap, 0)
	heap.Init(&h)
	heap.Push(&h, 1)
	heap.Push(&h, 8)
	heap.Push(&h, 10)
	heap.Push(&h, 5)
	heap.Push(&h, 3)
	max := heap.Pop(&h)
	if max.(int) != 10 {
		t.Fatalf("max is 10 but get %v", max)
	}

}
