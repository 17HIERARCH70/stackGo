package stackGo_test

import (
	"testing"

	"github.com/17HIERARCH70/stackGo"
)

// base tests
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

func TestStack_Compare(t *testing.T) {
	s := stackGo.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test comparing elements at valid indices
	result, err := s.Compare(0, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}

	result, err = s.Compare(1, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}

	result, err = s.Compare(2, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	// Test comparing elements at invalid indices
	_, err = s.Compare(0, 3)
	if err == nil || err.Error() != "index out of range" {
		t.Errorf("Expected 'index out of range' error, got %v", err)
	}

	_, err = s.Compare(-1, 2)
	if err == nil || err.Error() != "index out of range" {
		t.Errorf("Expected 'index out of range' error, got %v", err)
	}

	// Test comparing elements of different types
	s.Push("test")
	_, err = s.Compare(0, 3)
	if err == nil || err.Error() != "elements are not integers" {
		t.Errorf("Expected 'elements are not integers' error, got %v", err)
	}
}

func TestStack_Swap(t *testing.T) {
	s := stackGo.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test valid swap
	err := s.Swap(0, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := s.Pop()
	if val != 1 {
		t.Errorf("Expected value 1 at the top after swap, got %v", val)
	}

	val, _ = s.Pop()
	if val != 2 {
		t.Errorf("Expected value 2 at the second position after swap, got %v", val)
	}

	val, _ = s.Pop()
	if val != 3 {
		t.Errorf("Expected value 3 at the bottom after swap, got %v", val)
	}

	// Test invalid swap (index out of range)
	err = s.Swap(0, 3)
	if err == nil || err.Error() != "index out of range" {
		t.Errorf("Expected 'index out of range' error, got %v", err)
	}

	// Test invalid swap (index out of range)
	err = s.Swap(-1, 2)
	if err == nil || err.Error() != "index out of range" {
		t.Errorf("Expected 'index out of range' error, got %v", err)
	}
}
