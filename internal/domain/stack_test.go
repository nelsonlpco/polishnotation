package domain

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[Element]()

	if stack.Length() != 0 {
		t.Error("want stack length is 0")
	}
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Length() != 3 {
		t.Errorf("expected the stack length are equals to 3 but is %v", stack.Length())
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for i := 3; i > 0; i-- {
		el := stack.Pop()

		if el != i {
			t.Errorf("want el(%v) equals  %v ", el, i)
		}

		if stack.Length() != i-1 {
			t.Errorf("want stack length (%v) are equal to %v ", stack.Length(), i-1)
		}
	}
}

func TestStack_Value(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	el := stack.Value()
	if el != 3 {
		t.Errorf("want el(%v) equals  3 ", el)
	}
}
