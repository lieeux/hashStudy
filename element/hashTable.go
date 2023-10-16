package element

type HashTable struct { //含有一个链表数组
	LinkArr [7]NodeLink
}

func (h *HashTable) Insert(node *Node) {
	LinkNum := h.LinkSelect(node.Id) //得到要存放链表的下标
	h.LinkArr[LinkNum].Insert(node)  //使用对应的链表添加
}

func (h *HashTable) LinkSelect(id int) int { //散列方法
	return id % 7 //模7，余数是几就放几链
}

func (h *HashTable) List() { //显示整个哈希表
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].List(i)
	}
}

func (h *HashTable) Search(id int) *Node {
	LinkNum := h.LinkSelect(id)          //得到要查询链表的下标
	return h.LinkArr[LinkNum].Search(id) //使用对应的链表查询
}
