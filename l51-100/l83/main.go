package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-26 20:06
 **/

func main() {
    head := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 1,
            Next: &ListNode{
                Val: 2,
            },
        },
    }
    head = deleteDuplicates(head)
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    value := head.Val
    last := head
    for last.Next != nil {
        if last.Next.Val == value {
            tmp := last.Next
            for tmp != nil && tmp.Val == value {
                tmp = tmp.Next
            }
            last.Next = tmp
            if tmp == nil {
                break
            }
            value = tmp.Val
            last = last.Next
        } else {
            last = last.Next
            value = last.Val
        }
    }
    return head
}
