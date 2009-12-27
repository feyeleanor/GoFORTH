package goforth

import (
  "container/vector";
  "fmt";
  "os";
)

const (
	TRUE = 1;
	FALSE = 0;
)

type forthStack vector.IntVector;

func (s *forthStack)multiply() {
	s.Push(s.Pop() * s.Pop());
}

func (s *forthStack)add() {
	s.Push(s.Pop() + s.Pop());
}

func (s *forthStack)subtract() {
	x = s.Pop();
	s.Push(s.Pop() - x);
}

func (s *forthStack)print() {
	fmt.Fprintf(os.Stdout, s.Pop());
}

func (s *forthStack)emit() {
	fmt.Fprintf(os.Stdout, byte(s.Pop()));
}

func (s *forthStack)divide() {
	divisor := s.Pop();
	s.Push(s.Pop() / divisor);
}

func (s* forthStack)mod() {
	denominator := stack.Pop();
	stack.Push(stack.Pop() % denominator);
}

func (s *forthStack)divmod() {
  	denominator = s.Pop();
  	numerator := s.Pop();
		s.Push(numerator % denominator);
  	s.Push(numerator / denominator);
}

func (s *forthStack)less_than() {
	condition := s.Pop();
	if s.Pop() < condition { s.Push(TRUE) }
	else { s.Push(FALSE) }
}

func (s *forthStack)equal_to() {
		if s.Pop() == s.Pop() { s.Push(TRUE) }
		else { s.Push(FALSE) }
}

func (s *forthStack)not_equal_to() {
		if s.Pop() != s.Pop() { s.Push(TRUE) }
		else { s.Push(FALSE) }
}

func (s *forthStack)greater_than() {
	condition := s.Pop();
	if s.Pop() > condition { s.Push(TRUE) }
	else { s.Push(FALSE) }
}

func (s *forthStack)and() {
	s.Push(s.Pop() & s.Pop());
	s.Push(FALSE);
	s.not_equal_to();
}

func (s *forthStack)or() {
	s.Push(s.Pop() | s.Pop());
	s.Push(FALSE);
	s.not_equal_to();
}

func (s *forthStack)abs() {
	value := s.Pop();
	if value < 0 { value = -value }
	s.Push(value);
}

func (s *forthStack)over() {
	x, y := stack.Pop(), stack.Last();
	stack.Push(x);
	stack.Push(y);
}