/**
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

You may not alter the values in the nodes, only nodes itself may be changed.

Only constant memory is allowed.

For example,
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5
 */
package main

import (
    "fmt"
)

func main() {
    t := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 3,
                Next: &ListNode{
                    Val: 4,
                    Next: &ListNode{
                        Val: 5,
                    },
                },
            },
        },
    }
    t = reverseKGroup(t, 2)
    for ; t != nil; t = t.Next {
        fmt.Println(t.Val)
    }
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k < 2 {
        return head
    }
    var count int
    for tmp := head; tmp != nil; tmp = tmp.Next {
        count++
    }
    if count < k {
        return head
    }
    var res, next, tail *ListNode
    current := head
    //整理头部
    for i := 0; i < k-1; i++ {
        current = current.Next
    }
    next = current.Next
    current.Next = nil
    res = reverse(head, k)
    for tail = res; tail.Next != nil; tail = tail.Next {
    }

    for {
        current = next
        head = next
        for i := 0; i < k-1; i++ {
            if current == nil || current.Next == nil {
                goto Label
            }
            current = current.Next
        }
        next = current.Next
        current.Next = nil
        current = reverse(head, k)
        tail.Next = current
        for ; tail.Next != nil; tail = tail.Next {
        }
    }
Label:
    tail.Next = head
    return res
}

func reverse(head *ListNode, size int) *ListNode {
    res := head
    for i := 0; i < size-1; i++ {
        res = res.Next
    }
    tmp := head
    var last *ListNode

    for i := 0; i < size-1; i++ {
        last = tmp.Next
        tmp.Next = res.Next
        res.Next = tmp
        tmp = last
    }
    return res
}
