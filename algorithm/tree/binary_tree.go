package main

import (
    "strings"
    "fmt"
)

type Node struct {
    data  string
    left  *Node
    right *Node
}

func main() {
    stack := NewStack()
    stack.Push("a")
    stack.Push("b")
    stack.Push("c")
    for !stack.IsEmpty() {
        fmt.Println(stack.Pop())
    }
}

func GetNode(node *Node, data string) *Node {
    if node == nil {
        return nil
    }
    if node.data == data {
        return node
    }
    res := GetNode(node.left, data)
    if res != nil && res.data == data {
        return res
    }
    res = GetNode(node.right, data)
    if res != nil && res.data == data {
        return res
    }
    return nil

}

func display(node *Node) {
    if node == nil {
        return
    }
    fmt.Println(node.data)
    display(node.left)
    display(node.right)
}

func levelDisplay(node *Node) {
    if node == nil {
        return
    }
    q := NewQueue()
    q.Add(node)
    for !q.IsEmpty() {
        t := q.Remove().(*Node)
        fmt.Println(t.data)
        if t.left != nil {
            q.Add(t.left)
        }
        if t.right != nil {
            q.Add(t.right)
        }
    }
}

func create(data string) *Node {
    items := strings.Split(data, ", ")

    var res *Node

    for _, item := range items {
        tmps := strings.Split(item, "-")
        if res == nil {
            res = &Node{
                data: tmps[0],
            }
        }
        node := &Node{
            data: tmps[1],
        }
        t := GetNode(res, tmps[0])
        if t.left == nil {
            t.left = node
        } else if t.right == nil {
            t.right = node
        } else {
            panic("invalid children")
        }
    }
    return res
}
