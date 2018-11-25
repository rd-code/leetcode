package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-11-25 19:42
 **/

func main() {
    /*header := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 4,
            Next: &ListNode{
                Val: 3,
                Next: &ListNode{
                    Val: 2,
                    Next: &ListNode{
                        Val: 5,
                        Next: &ListNode{
                            Val: 2,
                        },
                    },
                },
            },
        },
    }*/
    header := &ListNode{
        Val: 3,
        Next: &ListNode{
            Val: 1,
            Next: &ListNode{
                Val: 2,
            },
        },
    }

    header = partition(header, 3)
    for header != nil {
        fmt.Println(header.Val)
        header = header.Next
    }

}

type ListNode struct {
    Val  int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    l := &ListNode{
        Next: head,
    }
    head = l
    for l.Next != nil && l.Next.Val < x {
        l = l.Next
    }
    cursor := l
    for cursor.Next != nil {
        if cursor.Next.Val < x {
            tmp := cursor.Next
            cursor.Next = tmp.Next
            tmp.Next = l.Next
            l.Next = tmp
            l = tmp
        } else {
            cursor = cursor.Next
        }
    }
    return head.Next
}
