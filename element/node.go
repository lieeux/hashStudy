package element

import "fmt"

type Node struct {
	Id   int
	Name string
	Next *Node
}

func (n *Node) ShowInfo() {
	fmt.Printf("在链表%d找到id为%d的雇员\n", n.Id%7, n.Id)
}
