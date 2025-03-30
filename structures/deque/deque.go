// Package deque provides an implementation for a Generic Double Ended Queue.
package deque


// node represents an element in a deque.
// it has a reference to the nodes in front and behind it.
type node[T comparable] struct {
	val T // The Value associated with this node.
	prev *node[T] // The Node prior to this one.
	next *node[T] // The Node that follows this one.
}

func newNode[T comparable](value T) *node[T] {
	return &node[T]{
		value,
		nil, nil,
	}
}

// A Deque (Double-Ended Queue) Struct.
// Implemented as a linked-list.
type Deque[T comparable] struct {
	size int // The current size of the deque.

	head *node[T] // The head of this deque, can be nil when empty.
	tail *node[T] // The tail of this deque, can be nil when empty.
}

func New[T comparable]() *Deque[T] {
	return &Deque[T] {
		0,
		nil, nil,
	}
}

// Size returns the current element count of this deque.
func (deque *Deque[T]) Size() int {
	return deque.size
}
func (deque *Deque[T]) IsEmpty() bool {
	return deque.Size() == 0
}

// PeekFront returns the first element in the Deque
//  without removing it.
//
// Returns:
//  - T: The value at the front of the deque when found, otherwise the default.
//  - bool: Whether or not an element was found at the front.
func (deque *Deque[T]) PeekFront() (T, bool) {

	if deque.head == nil {

		var empty T
		return empty, false
	}

	return deque.head.val, true
}
// PopFront returns the first element in the Deque
//  and removes it.
//
// Returns:
//  - T: The value at the front of the deque when found, otherwise the default.
//  - bool: Whether or not an element was found at the front.
func (deque *Deque[T]) PopFront() (T, bool) {
	if deque.head == nil {
		var empty T
		return empty, false
	}

	value := deque.head.val

	deque.size--

	deque.head = deque.head.next
	// Only if we are not empty now
	if deque.head != nil {
		deque.head.prev = nil
	} else {
		deque.tail = nil
	}

	return value, true
}

// PeekLast returns the last element in the Deque
//  without removing it.
//
// Returns:
//  - T: The value at the end of the deque when found, otherwise the default.
//  - bool: Whether or not an element was found at the end.
func (deque *Deque[T]) PeekLast() (T, bool) {

	if deque.tail == nil {
		var empty T
		return empty, false
	}

	return deque.tail.val, true
}
// PopLast returns the last element in the Deque
// and removes it.
//
// Returns:
//  - T: The value at the end of the deque when found, otherwise the default.
//  - bool: Whether or not an element was found at the end.
func (deque *Deque[T]) PopLast() (T, bool) {
	if deque.tail == nil {
		var empty T
		return empty, false
	}

	value := deque.tail.val

	deque.tail = deque.tail.prev
	
	// Only if we are not empty now
	if deque.tail != nil {
		deque.tail.next = nil
	} else {
		deque.head = nil
	}

	deque.size--

	return value, true
}


// Contains checks if the value presented is contained within this deque.
//
// Returns:
//  - bool: True when the element is found, otherwise false.
func (deque *Deque[T]) Contains(val T) bool {
	curr := deque.head
	for curr != nil {
		if curr.val == val {
			return true
		}
		curr = curr.next
	}
	return false
}

// PustLast appends the value to the end of the deque.
func (deque *Deque[T]) PushLast(value T) {

	node := newNode(value)

	deque.size += 1;

	if deque.IsEmpty() {
		deque.head = node
		deque.tail = node
		return
	}

	// Shift Tail
	node.prev = deque.tail
	deque.tail.next = node
	deque.tail = node
}
// PushAllLast appends the given values to the end of the deque.
func (deque *Deque[T]) PushAllLast(values ...T) {

	for _, val := range values {
		deque.PushLast(val)
	}
}
// PushFront inserts the value at the front of the deque.
func (deque *Deque[T]) PushFront(value T) {

	node := newNode(value)

	deque.size += 1;

	if deque.IsEmpty() {
		deque.head = node
		deque.tail = node
		return
	}

	// Shift Head
	node.next = deque.head
	deque.head.prev = node
	deque.head = node
}

// PushAllFront inserts the given values at the front of the deque.
// This will insert each element in the front, no longer
// preserving the order of the values passed in.
//
// Example:
//
//  deq := deque.New[int]()
//  deq.PushAllFront([]int{1, 2, 3}...)
//  // Deq now is {3, 2, 1}
func (deque *Deque[T]) PushAllFront(values ...T) {

	for _, val := range values {
		deque.PushFront(val)
	}
}

// At attempts to retrive the value at the specified index.
//
// Returns:
//  - T: The value at the given index if found, otherwise the default for T.
//  - bool: Whether or not a value was found at the given index.
func (deque *Deque[T]) At(index int) (T, bool) {
	var empty T
	if deque.size <= index {
		return empty, false
	}

	curr := deque.head

	for range index {
		// Failsafe
		if curr == nil {
			return empty, false
		}
		curr = curr.next
	}

	return curr.val, true
}

// Clear restores the deque to empty.
func (deque *Deque[T]) Clear() {
	deque.size = 0

	deque.head = nil
	deque.tail = nil
}

// Iterator returns a channel for iteration over the elements.
func (deque *Deque[T]) Iterator() <-chan T {
	ch := make(chan T)

	go func() {
		curr := deque.head

		for curr != nil {
			ch <- curr.val
			curr = curr.next
		}

		close(ch)
	}()

	return ch
}