package main

import "fmt"
type ListNode struct {
	Val int
	Next *ListNode
}
// 获取链表的长度
func GetListSize(head *ListNode) (size int) {
	size = 0
	//var newNode  *ListNode
	for head != nil {
		size ++
		head = head.Next
	}
	return
}
// 移动节点，使两个链表的从同等长度的节点，开始寻找它们交叉的节点
func MoveList(list *ListNode, num int) *ListNode {
	if list == nil {
		return list
	}
	//var head *ListNode
	for list != nil && num > 0 {
		list = list.Next
		num--
	}
	return list
}

// 获取两个链表的交叉的节点
func GetIntersectionNode(head1 *ListNode, head2 *ListNode) *ListNode {
	size := GetListSize(head1)
	size2 := GetListSize(head2)
	if size2 > size {
		head2 = MoveList(head2, size2 - size)
	} else {
		head1 = MoveList(head1, size - size2)
	}
	var node *ListNode
	for head1 != nil && head2 != nil  {

		if head1 == head2 {
			node = head1
			break
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	return node
}

func New(val int) *ListNode {
	return &ListNode{
		Val:val,
	}
}
func main()  {
	node1 := New(1)
	node2 := New(2)
	node3 := New(3)
	node4 := New(4)
	node5 := New(5)
	node6 := New(6)

	node7 := New(7)
	node8 := New(8)

	node7.Next = node8
	node8.Next = node5
	node5.Next = node6

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	// 头指针是一个空指针
	var head *ListNode
	head = node1
	var head2 *ListNode
	head2 = node7

	//size := GetListSize(head)
	//size2 := GetListSize(head2)
	node := GetIntersectionNode(head, head2)
	fmt.Println(node)
	//for head != nil {
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}
}

