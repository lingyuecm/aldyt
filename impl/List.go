package impl

import (
	"errors"
	"fmt"
)

type LinkedList[T any] struct {
	header *LinkedNode[T]
	size   int
}

func CreateLinkedList[T any]() *LinkedList[T] {
	l := new(LinkedList[T])

	l.header = new(LinkedNode[T])
	l.size = 0

	return l
}

func (l *LinkedList[T]) Insert(index int, data T) error {
	if index < 0 || index > l.size {
		return errors.New(fmt.Sprintf("Array Index out of Bound"))
	}
	node := l.header
	for m := 0; m < index; m++ {
		node = node.nextNode
	}
	nn := node.nextNode
	node.AppendData(data)
	node.nextNode.nextNode = nn
	l.size = l.size + 1
	return nil
}

func (l *LinkedList[T]) Remove(index int) error {
	if index < 0 || index >= l.size {
		return errors.New(fmt.Sprintf("Array Index out of Bound"))
	}
	node := l.header
	for m := 0; m < index; m++ {
		node = node.nextNode
	}
	nn := node.nextNode
	node.nextNode = nn.nextNode
	nn.nextNode = nil
	l.size = l.size - 1
	return nil
}

func (l *LinkedList[T]) Update(index int, data T) error {
	if index < 0 || index >= l.size {
		return errors.New(fmt.Sprintf("Array Index out of Bound"))
	}
	node := l.header
	for m := 0; m < index; m++ {
		node = node.nextNode
	}
	nn := node.nextNode.nextNode
	node.nextNode.nextNode = nil
	node.AppendData(data)
	node.nextNode.nextNode = nn
	return nil
}

func (l *LinkedList[T]) Retrieve(index int) (T, error) {
	if index < 0 || index >= l.size {
		return *new(T), errors.New(fmt.Sprintf("Array Index out of Bound"))
	}
	node := l.header
	for m := 0; m < index; m++ {
		node = node.nextNode
	}
	return node.nextNode.data, nil
}

func (l *LinkedList[T]) Append(data T) error {
	return l.Insert(l.size, data)
}

func (l *LinkedList[T]) ElementCount() int {
	return l.size
}
