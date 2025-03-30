package deque_test

import (
	"testing"

	"github.com/Bennm23/golibs/structures/deque"
)

func TestNewDeque(t *testing.T) {
	deq := deque.New[int]()

	if deq == nil {
		t.Error("Deque Init Failed")
	}
}

func TestPushLast(t *testing.T) {
	deq := deque.New[int]()

	deq.PushLast(5)
	deq.PushLast(10)

	if deq.Size() != 2 {
		t.Error("Push Last Size Incorrect")
	}

	if !deq.Contains(5) {
		t.Error("Push Last Contains Failed")
	}
	if !deq.Contains(10) {
		t.Error("Push Last Contains Failed")
	}

	val, found := deq.PeekFront()
	if !found {
		t.Error("Head Not Found")
	}
	if val != 5 {
		t.Error("Head expected = 5, was ", val)
	}

	val, found = deq.PeekLast()
	if !found {
		t.Error("Tail Not Found")
	}
	if val != 10 {
		t.Error("Tail expected = 10, was ", val)
	}
}
func TestPushFront(t *testing.T) {
	deq := deque.New[int]()

	deq.PushFront(5)
	deq.PushFront(10)

	if deq.Size() != 2 {
		t.Error("Push Front Size Incorrect")
	}

	if !deq.Contains(5) {
		t.Error("Push Front Contains Failed")
	}
	if !deq.Contains(10) {
		t.Error("Push Front Contains Failed")
	}

	val, found := deq.PeekFront()
	if !found {
		t.Error("Head Not Found")
	}
	if val != 10 {
		t.Error("Head expected = 10, was ", val)
	}

	val, found = deq.PeekLast()
	if !found {
		t.Error("Tail Not Found")
	}
	if val != 5 {
		t.Error("Tail expected = 5, was ", val)
	}
}
func TestPeekLast(t *testing.T) {
	deq := deque.New[int]()

	val, found := deq.PeekLast()
	if found != false {
		t.Error("Peek Last Should Have Been False")
	}
	if val != 0 {
		t.Error("Peek Last Value Should be 0")
	}
	deq.PushLast(5)
	val, found = deq.PeekLast()
	if found != true {
		t.Error("Peek Last Should Have Been True")
	}
	if val != 5 {
		t.Error("Peek Last Value Should be 5, was ", val)
	}
}

func TestPeekFront(t *testing.T) {
	deq := deque.New[int]()

	val, found := deq.PeekFront()
	if found != false {
		t.Error("Peek Front Should Have Been False")
	}
	if val != 0 {
		t.Error("Peek Front Value Should be 0")
	}
	deq.PushLast(5)
	val, found = deq.PeekFront()
	if found != true {
		t.Error("Peek Front Should Have Been True")
	}
	if val != 5 {
		t.Error("Peek Front Value Should be 5, was ", val)
	}
}

func TestPopFront(t *testing.T) {
	deq := deque.New[int]()

	val, found := deq.PopFront()
	if found != false {
		t.Error("Pop Front Should Have Been False")
	}
	if val != 0 {
		t.Error("Pop Front Value Should be 0")
	}

	deq.PushLast(10)
	val, found = deq.PopFront()
	if found != true {
		t.Error("Pop Front Should Have Been True")
	}
	if val != 10 {
		t.Error("Pop Front Value Should be 10, was ", val)
	}

	if !deq.IsEmpty() {
		t.Error("Deq should be empty after pop")
	}

	deq.PushFront(10)
	deq.PushLast(5)
	val, found = deq.PopFront()
	if found != true {
		t.Error("Pop Front Should Have Been True")
	}
	if val != 10 {
		t.Error("Pop Front Value Should be 10, was ", val)
	}

	if deq.Size() != 1 {
		t.Error("Deq should have size 1")
	}

	val, found = deq.PopFront()
	if found != true {
		t.Error("Pop Front Should Have Been True")
	}
	if val != 5 {
		t.Error("Pop Front Value Should be 5, was ", val)
	}

	if !deq.IsEmpty() {
		t.Error("Deq should be empty")
	}
}

func TestPopLast(t *testing.T) {
	deq := deque.New[int]()

	val, found := deq.PopLast()
	if found != false {
		t.Error("Pop Last Should Have Been False")
	}
	if val != 0 {
		t.Error("Pop Last Value Should be 0")
	}

	deq.PushFront(10)
	val, found = deq.PopLast()
	if found != true {
		t.Error("Pop Last Should Have Been True")
	}
	if val != 10 {
		t.Error("Pop Last Value Should be 10, was ", val)
	}

	if !deq.IsEmpty() {
		t.Error("Deq should be empty after pop")
	}

	deq.PushLast(10)
	deq.PushFront(5)
	val, found = deq.PopLast()
	if found != true {
		t.Error("Pop Last Should Have Been True")
	}
	if val != 10 {
		t.Error("Pop Last Value Should be 10, was ", val)
	}

	if deq.Size() != 1 {
		t.Error("Deq should have size 1")
	}

	val, found = deq.PopLast()
	if found != true {
		t.Error("Pop Last Should Have Been True")
	}
	if val != 5 {
		t.Error("Pop Last Value Should be 5, was ", val)
	}

	if !deq.IsEmpty() {
		t.Error("Deq should be empty")
	}
}

func TestPushAllFront(t *testing.T) {
	deq := deque.New[int]()

	values := []int {5, 6, 7, 8, 9}

	deq.PushAllFront(values...)

	if deq.Size() != len(values) {
		t.Errorf("Mismatch Size, expected = %d, was %d", len(values), deq.Size())
	}

	for _, expectedVal := range values {

		val, found := deq.PopLast()

		if !found {
			t.Error("Value Should Be Present")
		}
		if val != expectedVal {
			t.Error("Mismatch, Expected = ", expectedVal, " Was ", val)
		}
	}

	if !deq.IsEmpty() {
		t.Errorf("Deq should have been empty, had size = %d", deq.Size())
	}
}

func TestPushAllLast(t *testing.T) {
	deq := deque.New[int]()

	values := []int {5, 6, 7, 8, 9}

	deq.PushAllLast(values...)

	if deq.Size() != len(values) {
		t.Errorf("Mismatch Size, expected = %d, was %d", len(values), deq.Size())
	}

	for i := len(values) - 1; i >= 0; i-- {
		expectedVal := values[i]
		val, found := deq.PopLast()

		if !found {
			t.Error("Value Should Be Present")
		}
		if val != expectedVal {
			t.Error("Mismatch, Expected = ", expectedVal, " Was ", val)
		}
	}

	if !deq.IsEmpty() {
		t.Errorf("Deq should have been empty, had size = %d", deq.Size())
	}
}

func TestAt(t *testing.T) {
	deq := deque.New[int]()

	val, found := deq.At(0)
	if found != false {
		t.Error("Value for At should not be found")
	}
	if val != 0 {
		t.Errorf("Value for At should be 0, was %d", val)
	}

	values := []int{1, 2, 3}
	//3, 2, 1
	deq.PushAllFront(values...)

	for i := range values {

		val, found = deq.At(i)
		if found != true {
			t.Error("At found should be true")
		}
		expectedVal := values[len(values) - 1 - i]
		if val != expectedVal {
			t.Errorf("Value for At should be %d, was %d", expectedVal, val)
		}
	}

	if deq.Size() != len(values) {
		t.Errorf("Deque Size should be constant, was %d", deq.Size())
	}
}