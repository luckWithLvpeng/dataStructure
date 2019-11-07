package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	RootNode *Node
}

//Add 添加节点
//有序添加节点,左小右大
func (this *Tree) Add(v int) {
	node := &Node{Value: v}
	if this.RootNode == nil {
		this.RootNode = node
		return
	}
	current := this.RootNode
	for true {
		if current.Value == v {
			return
		} else if v < current.Value {
			if current.Left == nil {
				current.Left = node
				return
			} else {
				current = current.Left
			}

		} else {
			if current.Right == nil {
				current.Right = node
				return
			} else {
				current = current.Right
			}

		}

	}
}

//Append 添加节点
//广度优先，无规则添加
func (this *Tree) Append(v int) {
	node := &Node{Value: v}
	if this.RootNode == nil {
		this.RootNode = node
		return
	}
	queue := []*Node{this.RootNode}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left == nil {
			current.Left = node
			return
		} else {
			queue = append(queue, current.Left)
		}
		if current.Right == nil {
			current.Right = node
			return
		} else {
			queue = append(queue, current.Right)
		}
	}
}

//PreOrder 前序遍历
func (this *Tree) PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	this.PreOrder(node.Left)
	this.PreOrder(node.Right)
}

//MidOrder 中序遍历
func (this *Tree) MidOrder(node *Node) {
	if node == nil {
		return
	}
	this.MidOrder(node.Left)
	fmt.Println(node.Value)
	this.MidOrder(node.Right)
}

//PostOrder 后序遍历
func (this *Tree) PostOrder(node *Node) {
	if node == nil {
		return
	}
	this.PostOrder(node.Left)
	this.PostOrder(node.Right)
	fmt.Println(node.Value)
}

//BreadthTravel 广度遍历
func (this *Tree) BreadthTravel() {
	if this.RootNode == nil {
		return
	}
	queue := []*Node{this.RootNode}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.Value)
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

}
func main() {
	tree := Tree{}

	tree.Add(0)
	tree.Add(5)
	tree.Add(2)
	tree.Add(10)
	tree.Add(7)
	tree.Add(50)
	tree.Add(100)
	tree.Add(33)
	tree.Add(66)

	// tree.Append(0)
	// tree.Append(5)
	// tree.Append(2)
	// tree.Append(10)
	// tree.Append(7)
	// tree.Append(50)
	// tree.Append(100)
	// tree.Append(33)
	// tree.Append(66)

	tree.PreOrder(tree.RootNode)
	fmt.Println("-----------------")
	tree.MidOrder(tree.RootNode)
	fmt.Println("-----------------")
	tree.PostOrder(tree.RootNode)
	fmt.Println("-----------------")
	tree.BreadthTravel()

}
