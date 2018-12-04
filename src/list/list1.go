package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{
		Val:val,
	}
}
// 获取链表的逆序
func ReverseList(head *ListNode)  *ListNode  {
	var newNode *ListNode
	for head != nil {
		// 备份head的下一个节点
		next := head.Next
		// 将head的指向，指向到新的头结点，也就是将先前head反向指向新节点
		head.Next = newNode
		// 将head赋值个新节点
		newNode = head
		// 将下一个节点赋值给head
		head = next
	}
	return newNode
}
// 获取指定位置链表的逆序
func ReverseListBetween(head *ListNode, m int, n int) *ListNode {
	// 计算需要逆距的节点个数
	size := n - m +1
	// 初始化开始逆距节点的前驱
	var preHead *ListNode
	// 结果节点的前驱
	resultNode := head
	// 寻找开始逆距的节点
	for head != nil && m > 0 {
		preHead = head
		head = head.Next
		m--
	}
	var modifyListTail  *ListNode
	// 记录开始逆距的节点
	modifyListTail = head
	var newNode *ListNode
	// 从m处的节点开始逆距到n处的节点
	for head != nil && size > 0 {
		next := head.Next
		head.Next = newNode
		newNode = head
		head = next
		size--
	}
	// 预留开始逆距的节点，拼接到head
	modifyListTail.Next = head
	// 如果逆距是从0处节点开始
	if preHead != nil {
		preHead.Next = newNode
	}else {
		resultNode =  newNode
	}
	return resultNode
}
func main()  {
	node1 := New(1)
	node2 := New(2)
	node3 := New(3)
	node4 := New(4)
	node5 := New(5)
	node6 := New(6)
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	// 头指针是一个空指针
	var head *ListNode
	head = node1
	//for head != nil {
	//	//fmt.Println(head.Val)
	//	head = head.Next
	//	//fmt.Println(head)
	//}
	//newNode := ReverseList(head)
	newNode := ReverseListBetween(head, 0,4)
	for newNode != nil {
		fmt.Println(newNode.Val)
		newNode = newNode.Next
	}
}
