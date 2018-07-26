package main

import "fmt"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-26 19:47
 **/

type ListNode struct {
    Val  int
    Next *ListNode
}

func main() {
    /* head := &ListNode{
         Val: 1,
         Next: &ListNode{
             Val: 2,
             Next: &ListNode{
                 Val: 3,
                 Next: &ListNode{
                     Val: 3,
                     Next: &ListNode{
                         Val: 4,
                         Next: &ListNode{
                             Val: 4,
                             Next: &ListNode{
                                 Val: 5,
                             },
                         },
                     },
                 },
             },
         },
     }
     head = deleteDuplicates(head)*/
    head := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 1,
        },
    }
    head = deleteDuplicates(head)
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
    //1->2->3->3->4->4->5
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    res := &ListNode{ //头指针
        Next: head,
    }
    value := head.Val
    last := res
    for last.Next.Next != nil {
        if last.Next.Next.Val == value {
            tmp := last.Next
            for tmp.Next != nil && tmp.Next.Val == value {
                tmp = tmp.Next
            }
            last.Next = tmp.Next
            if tmp.Next == nil {
                break
            }
            value = tmp.Next.Val
        } else {
            value = last.Next.Next.Val
            last = last.Next
        }

    }

    return res.Next
}
