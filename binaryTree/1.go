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
func (this *Tree) Add(v int) {
	node := &Node{Value: v}
	//左小右大
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

	tree.PreOrder(tree.RootNode)
	fmt.Println("-----------------")
	tree.MidOrder(tree.RootNode)
	fmt.Println("-----------------")
	tree.PostOrder(tree.RootNode)

}
