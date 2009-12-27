// The stack package implements a container for managing linear arrays of elements.
// This code is patterned after vectors and like vectors, stacks can change size dynamically.

package stack

type container []interface{}

// Stack is the container itself.
// The zero value for Stack is an empty stack ready to use.
type Stack struct {
	elements container;
}

func (v *container)copy_to(destination *container) {
	for i, x := range v { destination[i] = x }
}

func (v *container)cap() int { return cap(v) }
func (v *container)len() int { return len(v) }

// Init initializes a new or resized stack.  The initial_len may be <= 0 to
// request a default length.  If initial_len is shorter than the current
// length of the Stack, trailing elements of the Stack will be cleared.
func (s *Stack) Init(initial_len int) *Stack {
	if s.elements.cap() == 0 || s.elements.cap() < initial_len {
		n := 8;
		if initial_len > n { n = initial_len }
		s.elements = make(container, n);
	} else {
		s.elements = s.elements[0:initial_len];
	}
	for i := range s.elements { stack.elements[i] = nil }
	return s;
}

// New returns an initialized new Vector with length at least len.
func New(len int) *Stack	{ return new(Stack).Init(len) }

// Len returns the number of elements in the vector.
// Len is 0 if p == nil.
func (s *Stack) Len() int {
	if s == nil { return 0 }
	else { return len(s.elements) }
}

// At returns the i'th element of the vector.
func (s *Stack) At(i int) interface{} { return s.elements[i] }

// Set sets the i'th element of the vector to value x.
func (s *Stack) Set(i int, x interface{})	{ s.elements[i] = x }

// Last returns the element in the vector of highest index.
func (s *Stack) Last() interface{}	{ return s.elements[s.elements.len() - 1] }

// Data returns all the elements as a slice.
func (s *Stack) Data() array {
	arr := make(container, s.Len());
	for i, v := range s.elements { arr[i] = v }
	return arr;
}


// Insert n elements at position i.
func expand(a []interface{}, i, n int) []interface{} {
	// make sure we have enough space
	len0 := len(a);
	len1 := len0 + n;
	if len1 < cap(a) {
		// enough space - just expand
		a = a[0:len1]
	} else {
		// not enough space - double capacity
		capb := cap(a) * 2;
		if capb < len1 {
			// still not enough - use required length
			capb = len1
		}
		// capb >= len1
		b := make([]interface{}, len1, capb);
		copy(b, a);
		a = b;
	}

	// make a hole
	for j := len0 - 1; j >= i; j-- {
		a[j+n] = a[j]
	}
	return a;
}


// Insert inserts into the vector an element of value x before
// the current element at index i.
func (s *Stack) Insert(i int, x interface{}) {
	s.elements = expand(s.elements, i, 1);
	s.elements[i] = x;
}

// Delete deletes the i'th element of the vector.  The gap is closed so the old
// element at index i+1 has index i afterwards.
func (p *Vector) Delete(i int) {
	a := p.a;
	n := len(a);

	a[i + 1:n].copy_to(a[i:n - 1])
	a[n-1] = nil;	// support GC, nil out entry
	p.a = a[0 : n-1];
}


// InsertVector inserts into the vector the contents of the Vector
// x such that the 0th element of x appears at index i after insertion.
func (p *Vector) InsertVector(i int, x *Vector) {
	p.a = expand(p.a, i, len(x.a));
	x.a.copy_to(p.a[i:i + len(x.a)]);
}


// Cut deletes elements i through j-1, inclusive.
func (p *Vector) Cut(i, j int) {
	a := p.a;
	n := len(a);
	m := n - (j - i);

	a[j:n].copy_to(a[i:m]);
	for k := m; k < n; k++ {
		a[k] = nil	// support GC, nil out entries
	}

	p.a = a[0:m];
}


// Slice returns a new Vector by slicing the old one to extract slice [i:j].
// The elements are copied. The original vector is unchanged.
func (p *Vector) Slice(i, j int) *Vector {
	s := New(j - i);	// will fail in Init() if j < j
	p.a[i:j].copy_to(s.a);
	return s;
}


// Do calls function f for each element of the vector, in order.
// The function should not change the indexing of the vector underfoot.
func (p *Vector) Do(f func(elem interface{})) {
	for i := 0; i < len(p.a); i++ {
		f(p.a[i])	// not too safe if f changes the Vector
	}
}


// Convenience wrappers

// Push appends x to the end of the vector.
func (p *Vector) Push(x interface{})	{ p.Insert(len(p.a), x) }


// Pop deletes the last element of the vector.
func (p *Vector) Pop() interface{} {
	i := len(p.a) - 1;
	x := p.a[i];
	p.a[i] = nil;	// support GC, nil out entry
	p.a = p.a[0:i];
	return x;
}


// AppendVector appends the entire Vector x to the end of this vector.
func (p *Vector) AppendVector(x *Vector)	{ p.InsertVector(len(p.a), x) }


// Partial sort.Interface support

// LessInterface provides partial support of the sort.Interface.
type LessInterface interface {
	Less(y interface{}) bool;
}


// Less returns a boolean denoting whether the i'th element is less than the j'th element.
func (p *Vector) Less(i, j int) bool	{ return p.a[i].(LessInterface).Less(p.a[j]) }


// Swap exchanges the elements at indexes i and j.
func (p *Vector) Swap(i, j int) {
	a := p.a;
	a[i], a[j] = a[j], a[i];
}


// Iterate over all elements; driver for range
func (p *Vector) iterate(c chan<- interface{}) {
	for _, v := range p.a {
		c <- v
	}
	close(c);
}


// Channel iterator for range.
func (p *Vector) Iter() <-chan interface{} {
	c := make(chan interface{});
	go p.iterate(c);
	return c;
}
