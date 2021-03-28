package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Sum the elements of two linked lists
func addLists(node1 *node, node2 *node, carry int) *node {
	if node1 == nil && node2 == nil && carry == 0 {
		return nil
	}

	var result *node
	value := carry

	if node1 != nil {
		value += node1.data
	}

	if node2 != nil {
		value += node2.data
	}

	node := &node{data: value % 10}
	result = node

	if node1 != nil || node2 != nil {
		var param3 int

		if node1 == nil {
			node1 = nil
		} else {
			node1 = node1.next
		}

		if node2 == nil {
			node2 = nil
		} else {
			node2 = node2.next
		}

		if value >= 10 {
			param3 = 1
		} else {
			param3 = 0
		}

		more := addLists(node1, node2, param3)
		result.next = more
	}

	return result
}

// Returns the length of linked list
func (list linkedList) getLength() int {
	return list.length
}

// Prints the elements of linked list
func (list linkedList) printList() {
	for list.head != nil {
		fmt.Printf("%v -> ", list.head.data)
		list.head = list.head.next
	}

	fmt.Println()
}

// Adds new node in the linked list
func (list *linkedList) pushBack(node *node) {
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++

		return
	}

	list.tail.next = node
	list.tail = node
	list.length++
}

// Sum the elements of current linked list with another linked list
func (list1 *linkedList) sumToAnother(list2 *linkedList) {
	result := addLists(list1.head, list2.head, 0)
	list1.head = result
}

func main() {
	node1 := &node{data: 7}
	node2 := &node{data: 1}
	node3 := &node{data: 6}
	linkedList1 := linkedList{}

	linkedList1.pushBack(node1)
	linkedList1.pushBack(node2)
	linkedList1.pushBack(node3)

	node4 := &node{data: 5}
	node5 := &node{data: 9}
	node6 := &node{data: 2}
	linkedList2 := linkedList{}

	linkedList2.pushBack(node4)
	linkedList2.pushBack(node5)
	linkedList2.pushBack(node6)

	linkedList1.sumToAnother(&linkedList2)
	linkedList1.printList()
}
