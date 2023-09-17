package main

import "fmt"

type Node struct {
	prev  *Node
	value int
	next  *Node
}

func log(format string, a ...any) {
	fmt.Println(fmt.Sprintf(format, a...))
}

func insertAtTail(head *Node, value int) {
	// log("testing log %v \n", 10)
	log("pointer is %p", head)
	log("value is %v\n", *head)
	head = &Node{prev: &Node{prev: nil, value: 10, next: nil}, value: 10, next: nil}
}

func main() {
	head := &Node{prev: nil, value: 0, next: nil}
	log("header pointer is %p", head)
	insertAtTail(head, 10)
	log("header pointer is %p", head)
	log("value is %v\n", *&head.prev.value)
}
