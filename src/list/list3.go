package main

import (
	"fmt"
	//"time"
)
type ListNode struct {
	Val int
	Next *ListNode
}
func New(val int) *ListNode {
	return &ListNode{
		Val:val,
	}
}
// 判断链表是否有环
// 原理：运用快慢指针，找到两个指针的相遇的节点。
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	var meet  *ListNode
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		//fmt.Println(fast.Val)
		//fmt.Println(slow.Val)
		fast = fast.Next
		if fast == slow {
			meet = fast
			break
		}
	}
	if meet == nil {
		return nil
	}
	for head != nil && meet != nil {
		if head == meet {
			return head
		}
		head = head.Next
		meet = meet.Next
	}
	return nil

}

func main() {
	node1 := New(1)
	node2 := New(2)
	node3 := New(3)
	node4 := New(4)
	node5 := New(5)
	node6 := New(6)
	node7 := New(7)
	//node8 := New(8)

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node3
	//node8.Next = node5

	// 头指针是一个空指针
	var head *ListNode
	head = node1
	node := detectCycle(head)
	fmt.Println(node)
	//size := GetListSize(head)
	//size2 := GetListSize(head2)
	//for head != nil {
	//	time.Sleep(time.Duration(1)*time.Second)
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}
}