package structures

import (
	"slices"
)

type Stack[T comparable] struct {

	_backed []T;
}

func NewStack[T comparable]() *Stack[T] {

	return &Stack[T]{
		_backed: make([]T, 0),
	}
}

func (stack *Stack[T]) Pop() T {
	if stack.Size() == 0 {
		panic("Cant Pop Empty Stack")
	}
	val := stack._backed[0]
	stack._backed = stack._backed[1:]
	return val
}
func (stack *Stack[T]) PushFront(val T) {
	stack._backed = append([]T{ val, }, stack._backed...)
}

func (stack *Stack[T]) PushFrontAll(vals []T) {
	stack._backed = append(vals, stack._backed...)
}

func (stack *Stack[T]) Size() int {
	return len(stack._backed)
}

func (stack *Stack[T]) Push(val T) {
	stack._backed = append(stack._backed, val)
}

func (stack *Stack[T]) PushAll(vals []T) {
	stack._backed = append(stack._backed, vals...)
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack._backed) == 0
}


func (stack *Stack[T]) SortFunction(scorer func(a T) float64) {
	
	slices.SortFunc(stack._backed, func(a T, b T) int {

		aScore := scorer(a)
		bScore := scorer(b)

		if aScore < bScore {
			return -1
		} else if aScore > bScore {
			return 1
		}
		return 0
	});
}

func (stack *Stack[T]) Items() []T {
	return stack._backed
}