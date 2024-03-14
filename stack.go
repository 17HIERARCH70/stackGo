package stackGo

import (
	"errors"
	"sync"
)

type Stack struct {
	sync.Mutex // Embed the mutex for easier locking
	s          []int
}

func NewStack() *Stack {
	return &Stack{s: make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.Lock()
	defer s.Unlock()

	s.s = append(s.s, v)
}

func (s *Stack) Pop() (int, error) {
	s.Lock()
	defer s.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("empty stack")
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

	if s.s[i] == s.s[k] {
		return 0, nil
	} else if s.s[i] < s.s[k] {
		return -1, nil
	} else {
		return 1, nil
	}
}

func (s *Stack) Swap(i, k int) error {
	s.Lock()
	defer s.Unlock()

	if i < 0 || i >= len(s.s) || k < 0 || k >= len(s.s) {
		return errors.New("index out of range")
	}

	s.s[i], s.s[k] = s.s[k], s.s[i]
	return nil
}
