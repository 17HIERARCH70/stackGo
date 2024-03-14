package stackGo

import (
	"errors"
	"sync"
)

type Stack struct {
	sync.Mutex // Embed the mutex for easier locking
	s          []interface{}
}

func NewStack() *Stack {
	return &Stack{s: make([]interface{}, 0)}
}

func (s *Stack) Push(v interface{}) {
	s.Lock()
	defer s.Unlock()

	s.s = append(s.s, v)
}

func (s *Stack) Pop() (interface{}, error) {
	s.Lock()
	defer s.Unlock()

	l := len(s.s)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *Stack) Len() int {
	s.Lock()
	defer s.Unlock()

	return len(s.s)
}
