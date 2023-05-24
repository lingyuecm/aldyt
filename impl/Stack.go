package impl

import (
	"errors"
	"fmt"
)

type FixedStack[T any] struct {
	elements []T
	size     int
}

func CreateFixedStack[T any](length int) (*FixedStack[T], error) {
	if length <= 0 {
		return nil, errors.New(fmt.Sprintf("Invalid Length for Stack: %d", length))
	}

	s := new(FixedStack[T])

	s.elements = make([]T, length, length)
	s.size = 0

	return s, nil
}

func (s *FixedStack[T]) Push(element T) error {
	if s.size >= cap(s.elements) {
		return errors.New(fmt.Sprintf("Stack Overflow"))
	}
	s.size = s.size + 1
	s.elements[s.size-1] = element
	return nil
}

func (s *FixedStack[T]) Pop() (T, error) {
	if s.size <= 0 {
		return *new(T), errors.New(fmt.Sprintf("Empty Stack"))
	}
	s.size = s.size - 1
	return s.elements[s.size], nil
}

func (s *FixedStack[T]) Top() (T, error) {
	if s.size <= 0 {
		return *new(T), errors.New(fmt.Sprintf("Empty Stack"))
	}
	return s.elements[s.size-1], nil
}

func (s *FixedStack[T]) ElementCount() int {
	return s.size
}
