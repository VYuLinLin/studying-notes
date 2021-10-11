package main

import "fmt"

// 将nil作为接收器
type IntNode struct {
	Value int
	Next  *IntNode
}

func (node *IntNode) Sum() int {
	if node == nil {
		return 0
	}
	return node.Value + node.Next.Sum()

}

func main() {
	node1 := IntNode{30, nil}
	node2 := IntNode{12, nil}
	node3 := IntNode{43, nil}

	node1.Next = &node2
	node2.Next = &node3

	fmt.Println(node1.Sum()) // 85
	fmt.Println(node2.Sum()) // 55

	node := &IntNode{3, nil}
	node = nil
	fmt.Println(node.Sum()) // 0
}
