package goforth

import (
  "stack";
  "fmt";
  "os";
)

const (
	TRUE = 1;
	FALSE = 0;
)

type ForthStack stack.Stack;

func NewStack() ForthStack {
	return stack.New(64);
}

func (s *ForthStack) Swap() {
	i := s.TopIndex();
	stack.Stack(s).Swap(i - 1, i)
}

func (s *ForthStack) Multiply() {
	s.Push(s.Pop() * s.Pop());
}

func (s *ForthStack) Add() {
	s.Push(s.Pop() + s.Pop());
}

func (s *ForthStack) Subtract() {
	s.Swap();
	s.Push(s.Pop() - s.Pop());
}

func (s *ForthStack) Print() {
	fmt.Fprintf(os.Stdout, s.Pop());
}

func (s *ForthStack) Emit() {
	fmt.Fprintf(os.Stdout, byte(s.Pop()));
}

func (s *ForthStack) Divide() {
	s.Swap();
	s.Push(s.Pop() / s.Pop());
}

func (s* ForthStack) Mod() {
	s.Swap();
	stack.Push(stack.Pop() % stack.Pop());
}

func (s *ForthStack) Divmod() {
  	denominator := s.Pop();
  	numerator := s.Pop();
	s.Push(numerator % denominator);
  	s.Push(numerator / denominator);
}

func (s *ForthStack) LessThan() {
	s.Swap();
	if s.Pop() < s.Pop() { s.Push(TRUE) }
	else { s.Push(FALSE) }
}

func (s *ForthStack) Equal() {
	if s.Pop() == s.Pop() { s.Push(TRUE) }
	else { s.Push(FALSE) }
}

func (s *ForthStack) NotEqual() {
	if s.Pop() != s.Pop() { s.Push(TRUE) }
	else { s.Push(FALSE) }
}

func (s *ForthStack) GreaterThan() {
	s.Swap();
	if s.Pop() > s.Pop() { s.Push(TRUE) }
	else { s.Push(FALSE) }
}

func (s *ForthStack) And() {
	s.Push(s.Pop() & s.Pop());
	s.Push(FALSE);
	s.NotEqual();
}

func (s *ForthStack) Or() {
	s.Push(s.Pop() | s.Pop());
	s.Push(FALSE);
	s.NotEqual();
}

func (s *ForthStack) Abs() {
	value := s.Pop();
	if value < 0 { value = -value }
	s.Push(value);
}

func (s *ForthStack) Minus() {
	s.At(s.Topndex(), -s.Top())
}

func (s *ForthStack) Minimum() {
	x, y := s.Pop(), s.Pop();
	if x < y { stack.Push(x) }
	else { stack.Push(y) }
}

func (s *ForthStack) Maximum() {
	x, y := s.Pop(), s.Pop();
	if x > y { stack.Push(x) }
	else { stack.Push(y) }
}