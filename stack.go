// The stack package implements a container for managing linear arrays of elements.
// This code is patterned after vectors and like vectors, stacks can change size dynamically.

package stack

import(
	"container";
)

const(
	MINIMUM_STACK_SIZE = 8;
)

const(
	OK = iota;
	EMPTY;
	UNDERFLOW;
	OVERFLOW;
)

// Stack is the container itself.
// The zero value for Stack is an empty stack ready to use.
type Stack struct {
	elements		container.Container;
	error_status	int;
}

func New(desired_elements int) *Stack {
	allocation_length := MINIMUM_STACK_SIZE;
	if desired_elements > allocation_length { allocation_length = desired_elements }
	return Stack{elements: container.New(allocation_length, 0)};
}

func (s *Stack) ClearErrors() {
	s.error_status = OK;
}

func (s *Stack) IsValid() bool {
	if error_status == OK { return true }
	else { return false }
}

func (s *Stack) HasDepth(minimum_size int) bool {
	if !s.IsValid() { return false }
	if s.Len() >= minimum_size {
		return true
	} else {
		s.error_status = UNDERFLOW;
		return false;
	}
}

func (s *Stack) Cap() int {
	return s.elements.Cap();
}

func (s *Stack) Len() int {
	return s.elements.Len();
}

func (s *Stack) At(position uint) Value {
	if s.HasDepth(position) { return s.elements[position] }
	s.error_status = OVERFLOW;
	return nil;
}

func (s *Stack) Pick(position uint) Value {
	if s.HasDepth(position) { return s.elements[s.TopIndex() - position] }
	s.error_status = OVERFLOW;
	return nil;
}

func (s *Stack) Set(position uint, x Value) bool {
	if s.HasDepth(position) {
		s.elements[position] = x;
		return true;
	}
	s.error_status = OVERFLOW;
	return false;
}

func (s *Stack) Top() Value {
	return s.elements.Last();
}

func (s *Stack) TopIndex() int {
	return s.Len() - 1;
}

func (s *Stack) SetTop(x Value) {
	s.Set(s.TopIndex(), x);
}

func (s *Stack) Insert(position uint, x Value) {
	s.elements = s.elements.Expand(position, 1);
	s.elements[position] = x;
}

func (s *Stack) Slice(start, finish uint) *Stack {
	new_stack := Stack{elements: s.elements[start:finish].Duplicate()};
	return new_stack;
}

func (s *Stack) Cut(start, finish uint) Stack {
	elements := s.elements;
	new_stack := s.Slice(start, finish);
	n := elements.Len();
	m := n - (finish - start);
	elements[finish:n].CopyTo(elements[start:m]);
	for k := m; k < n; k++ { elements[k] = nil }
	s.elements = elements[0:m];
	return new_stack;
}

func (s *Stack) Delete(position uint) {
	elements := s.elements;
	n := elements.Len();

	elements[position + 1 : n].CopyTo(elements[position : n - 1]);
	elements[n - 1] = nil;
	s.elements = elements[0 : n - 1];
}

// Convenience wrappers

func (s *Stack) Push(x Value) {
	s.Insert(s.Len(), x);
}

func (s *Stack) Pop() Value {
	i := s.Len() - 1;
	if i < 1 {
		s.error_status = EMPTY;
		return nil;
	} else {
		x := s.elements[i];
		s.elements[i] = nil;
		s.elements = s.elements[0:i];
		return x;
	}
}

func (s *Stack) Drop(elements uint) bool {
	if s.HasDepth(elements) {
		s.Delete(s.Len() - elements);
	} else {
		s.error_status = UNDERFLOW;
		return false;
	}
}

func (s *Stack) AppendStack(x *Stack) {
	s.InsertVector(s.Len(), x);
}

func (s *Stack) Swap(i, j int) {
	elements := s.elements;
	elements[i], elements[j] = elements[j], elements[i];
}

func (s *Stack) Dup() {
	s.Push(Top());
}

func (s *Stack) Rot() {
	e := s.elements;
	i := e.Len() - 1;
	j, k := i - 1, i - 2;
	e[k], e[j], e[i] = e[j], e[i], e[k]
}

func (s *Stack) Over() {
	s.Push(s.elements[s.Len() - 2]);
}