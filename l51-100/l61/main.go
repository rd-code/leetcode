package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-17 18:19
 **/

func main() {
    //1->2->3->4->5->NULL, k = 2
    l := &ListNode{
        Val: 1,
        /*Next: &ListNode{
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
        },*/
    }
    display(l)

    l = rotateRight(l, 1)
    fmt.Println("=======")
    display(l)
}

func display(l *ListNode) {
    for l != nil {
        fmt.Println(l.Val)
        l = l.Next
    }
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil || head.Next == nil {
        return head
    }
    tmp := head
    size := 0
    for tmp != nil {
        size += 1
        tmp = tmp.Next
    }
    k = k % size
    if k == 0 {
        return head
    }
    tmp = head
    for i := 0; i < size-k-1; i++ {
        tmp = tmp.Next
    }

    tmp, tmp.Next = tmp.Next, nil
    tail := tmp
    for tail.Next != nil {
        tail = tail.Next
    }
    tail.Next = head
    return tmp
}
