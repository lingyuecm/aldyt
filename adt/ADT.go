package adt

type Collection[T any] interface {
	ElementCount() int
}

type Stack[T any] interface {
	*Collection[T]
	Push(element T) error
	Pop() (T, error)
	Top() (T, error)
}
