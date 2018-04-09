package main

type Queue struct {
    data interface{}
    next *Queue
}

func NewQueue() *Queue {
    return &Queue{}
}

func (q *Queue) Add(data interface{}) {
    t := q
    for t.next != nil {
        t = t.next
    }
    t.next = &Queue{
        data: data,
    }
}
func (q *Queue) Remove() interface{} {
    t := q.next
    q.next = q.next.next
    return t.data
}

func (q *Queue) IsEmpty() bool {
    return q.next == nil
}
