package queue

import (
	"container/list"
	"testing"
)

func Benchmark_ListQueue(b *testing.B) {
	q := list.New()
	for i := 0; i < b.N; i++ {
		q.PushBack([1024]byte{})
		q.PushBack([1024]byte{})
		q.Remove(q.Front())
	}
}
func Benchmark_SliceQueue(b *testing.B) {
	var q [][1024]byte
	for i := 0; i < b.N; i++ {
		q = append(q, [1024]byte{})
		q = append(q, [1024]byte{})
		q = q[1:]
	}
}
