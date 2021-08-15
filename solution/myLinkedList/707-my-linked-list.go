package myLinkedList

import "fmt"

type Item struct {
	val  int
	next *Item
}

type MyLinkedList struct {
	len  int
	head *Item
	tail *Item
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index > this.len {
		return -1
	}
	node := this.head
	for index > 0 {
		if node == nil {
			break
		}
		node = node.next
		index--
	}
	if node == nil {
		return -1
	}
	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Item{val, this.head}
	if this.len == 0 {
		this.tail = this.head
	}
	this.len++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.tail = &Item{val, nil}
	if this.len == 0 {
		this.head = this.tail
	}
	this.len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
	}
	prev := this.head
	index--
	for index > 0 {
		if prev.next == nil {
			return
		}
		prev = prev.next
		index--
	}
	prev.next = &Item{val, prev.next}
	this.len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.len {
		return
	}
	prev := this.head
	for index > 0 {
		if prev.next == nil {
			return
		}
		prev = prev.next
		index--
	}
	if prev.next != nil {
		prev.next = prev.next.next
	} else {
		prev.next = nil
	}
	this.len--
}

func (this *MyLinkedList) Print() {
	node := this.head
	for node != nil {
		fmt.Print(*node, " ")
		node = node.next
	}
	fmt.Print("\n")
}
