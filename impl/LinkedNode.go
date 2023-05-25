package impl

type LinkedNode[T any] struct {
	data     T
	nextNode *LinkedNode[T]
}

func CreateLinkedNode[T any]() *LinkedNode[T] {
	return new(LinkedNode[T])
}

func CreateLinkedNodeWithData[T any](data T) *LinkedNode[T] {
	node := CreateLinkedNode[T]()

	node.data = data

	return node
}

func (node *LinkedNode[T]) AppendData(data T) {
	node.nextNode = CreateLinkedNodeWithData[T](data)
}

type DoublyLinkedNode[T any] struct {
	data         T
	nextNode     *DoublyLinkedNode[T]
	previousNode *DoublyLinkedNode[T]
}

func CreateDoublyLinkedNode[T any]() *DoublyLinkedNode[T] {
	return new(DoublyLinkedNode[T])
}

func CreateDoublyLinkedNodeWithData[T any](data T) *DoublyLinkedNode[T] {
	node := CreateDoublyLinkedNode[T]()

	node.data = data

	return node
}

func (node *DoublyLinkedNode[T]) AppendData(data T) {
	node.nextNode = CreateDoublyLinkedNodeWithData[T](data)
	node.nextNode.previousNode = node
}
