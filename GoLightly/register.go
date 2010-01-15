package golightly

import (
	"container/vector"
	"fmt";
	"os";
	"bufio";
	"strings";
)

type RegisterBlock [32]int

func (block *RegisterBlock) Load(register uint, value int) {
	if register < Len(block) { block[register] = value }
}

func (block *RegisterBlock) Read(register uint) int {
	if register < Len(block) { return block[register] }
	return nil;
}

func (block *RegisterBlock) Copy(source, target uint) {
	register_count := Len(block);
	if source < register_count && target < register_count { block[target] = block[source] }
}

func (block *RegisterBlock) Add(register uint, value int) {
	if register < Len(block) { block[register] += value }
}

func (block *RegisterBlock) Subtract(register uint, value int) {
	if register < Len(block) { block[register] -= value }
}

func (block *RegisterBlock) Multiply(register uint, value int) {
	if register < Len(block) { block[register] *= value }
}

func (block *RegisterBlock) Divide(register uint, value int) {
	if register < Len(block) { block[register] /= value }
}

func (block *RegisterBlock) And(register uint, value int) {
	if register < Len(block) { block[register] &= value }
}

func (block *RegisterBlock) Or(register uint, value int) {
	if register < Len(block) { block[register] |= value }
}

func (block *RegisterBlock) Xor(register uint, value int) {
	if register < Len(block) { block[register] ^= value }
}