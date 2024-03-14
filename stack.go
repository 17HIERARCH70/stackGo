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

func (s *Stack) Compare(i, k int) (int, error) {
	s.Lock()
	defer s.Unlock()

	if i < 0 || i >= len(s.s) || k < 0 || k >= len(s.s) {
		return 0, errors.New("index out of range")
	}

	elemI, okI := s.s[i].(int)
	elemK, okK := s.s[k].(int)
	if !okI || !okK {
		return 0, errors.New("elements are not integers")
	}

	if elemI == elemK {
		return 0, nil
	} else if elemI < elemK {
		return -1, nil
	} else {
		return 1, nil
	}
}

// swap func
func (s *Stack) Swap(i, j int) error {
	s.Lock()
	defer s.Unlock()

	if i < 0 || i >= len(s.s) || j < 0 || j >= len(s.s) {
		return errors.New("index out of range")
	}

	s.s[i], s.s[j] = s.s[j], s.s[i]
	return nil
}
