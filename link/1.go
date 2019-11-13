package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
	List  *List
}

type List struct {
	rootNode Node
	Len      int
}

//Init  初始化或者清空
func (this *List) Init() *List {
	this.Len = 0
	this.rootNode.Next = &this.rootNode
	this.rootNode.Prev = &this.rootNode
	return this
}

//NewList 新建
func NewList() *List {
	return new(List).Init()
}

//insert 在at 后插入e
func (this *List) insert(e, at *Node) *Node {
	tmp := at.Next
	at.Next = e
	e.Next = tmp
	e.Prev = at
	tmp.Prev = e
	e.List = this
	this.Len++
	return e
}

//remove 删除 e
func (this *List) remove(e *Node) *Node {
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	//避免内存泄漏
	e.Next = nil
	e.Prev = nil
	e.List = nil
	this.Len--
	return e
}

//DeleteAll 删除 所有值为 v
func (this *List) DeleteAll(v int) {
	cur_node := this.rootNode.Next
	for cur_node != nil && cur_node.Value >= v && cur_node != &this.rootNode {
		if cur_node.Value == v {
			this.remove(cur_node)
		}
		cur_node = cur_node.Next
	}
}

//Insert 插入 v 值, 从大到小的顺序
func (this *List) Insert(v int) *Node {
	node := &Node{
		Value: v,
		List:  this,
	}
	cur_node := this.rootNode.Next
	for cur_node != nil && cur_node.Value > v && cur_node != &this.rootNode {
		cur_node = cur_node.Next
	}
	return this.insert(node, cur_node.Prev)
}
func main() {
	fmt.Println("vim-go")
}
