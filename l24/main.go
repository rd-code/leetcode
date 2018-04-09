package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func main() {
    /*
    [2,5,3,4,6,2,2]
     */
    items := &ListNode{
        Val: 2,
        Next: &ListNode{
            Val: 5,
            Next: &ListNode{
                Val: 3,
                Next: &ListNode{
                    Val: 4,
                    Next: &ListNode{
                        Val: 6,
                        Next: &ListNode{
                            Val: 2,
                            Next: &ListNode{
                                Val: 2,
                            },
                        },
                    },
                },
            },
        },
    }
    t := swapPairs(items)
    for t != nil {
        fmt.Println(t.Val)
        t = t.Next
    }
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    tmp := head
    fmt.Println(tmp.Val)
    result := head.Next
    head.Next = result.Next
    result.Next = tmp
    var current *ListNode
    tail := result.Next
    current = result.Next
    last := result.Next
    for current.Next != nil && current.Next.Next != nil {
        last = current.Next
        current = current.Next.Next
        last.Next = current.Next
        current.Next = last
        tail.Next = current
        current = last
        tail = current
    }

    return result
}
