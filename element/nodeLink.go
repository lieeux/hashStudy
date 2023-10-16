package element

import "fmt"

type NodeLink struct { //不带表头，第一个节点就存放雇员
	Head *Node
}

func (n *NodeLink) Insert(node *Node) { //按id顺序由小到大添加
	if n.Head == nil { //如果是空链，直接把新节点作为表头
		n.Head = node
		return
	}
	temp := n.Head
	for { //找插入位置
		if temp.Next == nil {
			break
		} else if temp.Next.Id > node.Id { //如果新加节点的id比头结点小，会出问题
			break
		}
		temp = temp.Next
	}
	node.Next = temp.Next //先把新节点指向下一个节点
	temp.Next = node      //再把当前节点指向新节点，不然链就断了
}

func (n *NodeLink) List(num int) {
	if n.Head == nil {
		fmt.Printf("链表%d为空\n", num)
		return
	}
	temp := n.Head
	for {
		fmt.Printf("链表%d-[%d %s]-->", num, temp.Id, temp.Name)
		temp = temp.Next //先跳转
		if temp == nil { //再判断
			break
		}
	}
	fmt.Println()
}

func (n *NodeLink) Search(id int) *Node {
	if n.Head == nil {
		return nil
	}
	temp := n.Head
	for {
		if temp.Id == id {
			return temp
		}
		temp = temp.Next
		if temp == nil {
			return nil
		}
	}
}
