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

type List[T any] interface {
	*Collection[T]
	Insert(index int, data T) error
	Remove(index int) error
	Update(index int, data T) error
	Retrieve(index int) (T, error)
	Append(data T) error
}
