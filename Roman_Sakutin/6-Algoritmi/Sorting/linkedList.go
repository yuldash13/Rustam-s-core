package main

//func main() {
//	list1 := NewList(1, NewList(2, NewList(4, nil)))
//	list2 := NewList(1, NewList(3, NewList(4, nil)))
//
//	list := mergeTwoLists(list1, list2)
//
//	for cur := list; cur != nil; cur = cur.Next {
//		fmt.Println(cur.Val)
//	}
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(val int, List *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: List,
	}
}

func isPalindrome(head *ListNode) bool {
	cur := head
	if head.Val == head.Next.Val {
		return true
	} else {
		head = head.Next
	}
	if isPalindrome(head) == true {
		if cur.Next.Next.Next != nil {
			cur.Next = cur.Next.Next.Next
			return isPalindrome(cur)
		} else {
			return isPalindrome(cur)
		}
	}
	return false
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := &ListNode{}
	ll := list1
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			ll.Next = list1
			list1 = list1.Next
		} else {
			ll.Next = list2
			list2 = list2.Next
		}
		ll = ll.Next
	}
	if list1 != nil {
		ll.Next = list1
	}
	if list2 != nil {
		ll.Next = list2
	}
	return list.Next
}
