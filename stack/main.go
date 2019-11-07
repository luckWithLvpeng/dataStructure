package main

import (
	"sync"
)

type Item interface{}
type Node struct {
	Value Item
	Next  *Node
}

type Stack struct {
	Top    *Node
	Length int
	Lock   *sync.RWMutex
}

// Init  初始化或者清空栈
func (this *Stack) Init() *Stack {
	this.Top = nil
	this.Length = 0
	this.Lock = &sync.RWMutex{}
	return this
}

// Len  长度
func (this *Stack) Len() int {
	return this.Length
}

// Peek  偷看
func (this *Stack) Peek() Item {
	if this.Length == 0 {
		return nil
	}
	return this.Top.Value
}

// Pop  弹出
func (this *Stack) Pop() Item {
	if this.Length == 0 {
		return nil
	}
	this.Lock.Lock()
	defer this.Lock.Unlock()
	cur_node := this.Top
	this.Top = cur_node.Next
	this.Length--
	return cur_node.Value
}

// Push  压入
func (this *Stack) Push(v Item) {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	node := &Node{Value: v}
	node.Next = this.Top
	this.Top = node
	this.Length++
}

// NewStack 创建Stack
func NewStack() *Stack {
	return new(Stack).Init()
}
func main() {

}
