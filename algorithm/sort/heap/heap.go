package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Heap struct {
    data []int
    size int
}

func (h *Heap) Parent(i int) int {
    return i / 2
}
func (h *Heap) Left(i int) int {
    return i * 2
}

func (h *Heap) Right(i int) int {
    return i*2 + 1
}

func (h *Heap) Hold(i int) {
    if i >= h.size {
        return
    }
    l := h.Left(i)
    r := h.Right(i)
    max := i
    if l < h.size && h.data[l] > h.data[i] {
        max = l
    }
    if r < h.size && h.data[r] > h.data[max] {
        max = r
    }
    if max != i {
        h.data[i], h.data[max] = h.data[max], h.data[i]
        h.Hold(max)
    }
}

func (h *Heap) Sort() {
    for h.size > 0 {
        h.data[h.size-1], h.data[0] = h.data[0], h.data[h.size-1]
        h.size -= 1
        h.Hold(0)
    }
}

func Build(data []int) *Heap {
    h := &Heap{
        data: data,
        size: len(data),
    }
    for i := h.size / 2; i >= 0; i-- {
        h.Hold(i)
    }
    return h
}

func check(data []int) bool {
    for i := 0; i < len(data)-1; i++ {
        if data[i] > data[i+1] {
            return false
        }
    }
    return true
}

func main() {
    rand := rand.New(rand.NewSource(time.Now().Unix()))
    items := make([]int, 100)
    for i := 0; i < 100; i++ {
        items[i] = rand.Intn(1000)
    }
    t := Build(items)
    t.Sort()

    fmt.Println(check(t.data))

}
