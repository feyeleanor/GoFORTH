package goforth

// The Dictionary is the core of any FORTH system, allowing words to be defined and then
// recalled as needed. It comprises a lookup table indexed by name and a pointer to the
// ByteCode to be executed when the word is invoked.

import(
	"container/vector";
)

const(
	CodeLibrary = vector.NewVector();
)

type Identifier string
type ByteCode []byte
type Dictionary struct {
	Index			map[string] uint;
	CodeLibrary		vector.Vector;
}

func New() *Dictionary {
	return &Dictionary{Index: create(map[string] int), CodeLibrary: vector.NewVector(100)};
}

func (d* Dictionary) Here() uint {
	return d.CodeLibrary.Len() - 1;
}

func (d* Dictionary) Create(name Identifier) uint {
	d.CodeLibrary.Push(nil);
	d[name] = d.Here();
	return d[name];
}

func (d *Dictionary) Does(index uint, code ByteCode) {
	d.CodeLibrary.SetAt(index, code);
}