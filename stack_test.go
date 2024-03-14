package stackGo_test

import (
	"testing"

	"github.com/17HIERARCH70/stackGo"
)

func TestStack_PushPop(t *testing.T) {
	s := stackGo.NewStack()

	s.Push(1)
	s.Push(2)

	val, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected value 2, got %v", val)
	}

	val, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected value 1, got %v", val)
	}

	// Test popping from an empty stack
	_, err = s.Pop()
	if err == nil || err.Error() != "empty stack" {
		t.Errorf("Expected 'empty stack' error, got %v", err)
	}
}

func TestStack_EmptyStack(t *testing.T) {
	s := stackGo.NewStack()

	_, err := s.Pop()
	if err == nil || err.Error() != "empty stack" {
		t.Errorf("Expected 'empty stack' error, got %v", err)
	}
}

func TestStack_PushPopMixedTypes(t *testing.T) {
	s := stackGo.NewStack()

	s.Push(1)
	s.Push("test")
	s.Push(3.14)

	val, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 3.14 {
		t.Errorf("Expected value 3.14, got %v", val)
	}

	val, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != "test" {
		t.Errorf("Expected value 'test', got %v", val)
	}

	val, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected value 1, got %v", val)
	}
}
