package structures

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func CreateListNode (arr []int) ListNode {
	list := ListNode{}
	ptr := &list
	for _, v := range arr {
		ptr.Val = v
		ptr.Next = new(ListNode)
		ptr = ptr.Next
	}

	return list
}

func DoubleList (l ListNode) {
	if l.Next ==  nil {
		return
	}
	fmt.Println(l.Val)
	DoubleList(*l.Next)
}

func DeleteDublicates(l *ListNode, mapList map[int]*ListNode) {
	if l.Next ==  nil {
		return
	}
	parentPtr, ok := mapList[l.Val]
	if !ok {
		mapList[l.Val] = l
		parentPtr = l
	} else {
		parentPtr.Next = l.Next
	}
	DeleteDublicates(l.Next, mapList)
}

