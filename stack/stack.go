package stack

import "fmt"

type Stack struct {
	st  []interface{}
	top int
}

func (s *Stack) Pop() interface{} {
	if s.top == 0 {
		fmt.Println("Stack is empty!")
		return nil
	}
	s.top--
	topItem := s.st[s.top]
	s.st = s.st[:s.top]
	return topItem
}

func (s *Stack) Top() interface{} {
	if s.top == 0 {
		fmt.Println("Stack is empty!")
		return nil
	}
	return s.st[s.top-1]
}

func (s *Stack) Push(t interface{}) {
	s.top++
	s.st = append(s.st, t)
}

func NewStack() *Stack {
	return &Stack{
		st:  make([]interface{}, 0),
		top: 0,
	}
}
