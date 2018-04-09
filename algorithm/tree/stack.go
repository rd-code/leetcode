package main

import "github.com/pkg/errors"

type Stack struct {
    data []interface{}
}

func NewStack() *Stack {
    return &Stack{}
}

func (s *Stack) Push(data interface{}) {
    s.data = append(s.data, data)
}

func (s *Stack) Pop() (interface{}, error) {
    if len(s.data) == 0 {
        err := errors.New("empty statck cannot pop")
        return nil, err
    }
    data := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return data, nil
}

func (s *Stack) IsEmpty() bool {
    return len(s.data) == 0
}
