package main

import "hashStudy/element"

func main() {
	var hashtable element.HashTable
	emp1 := &element.Node{
		Id:   5,
		Name: "xue",
	}
	emp2 := &element.Node{
		Id:   19,
		Name: "lei",
	}
	emp3 := &element.Node{
		Id:   12,
		Name: "leo",
	}
	hashtable.Insert(emp1)
	hashtable.Insert(emp2)
	hashtable.Insert(emp3)
	hashtable.List()
	emp := hashtable.Search(12)
	emp.ShowInfo()
}
