package golightly

import (
	"container/vector"
	"fmt";
	"os";
	"bufio";
	"strings";
)

const(
	BANNER = "GoLightly VM v0.0a";
	NOOP = 0;
	ADD;
	ADD_VALUE;
	SUBTRACT;
	SUBTRACT_VALUE;
	MULTIPLY;
	MULTIPLY_VALUE;
	DIVIDE;
	DIVIDE_VALUE;
	AND;
	AND_VALUE;
	OR;
	OR_VALUE;
	XOR;
	XOR_VALUE;
	COPY;
	COPY_VALUE;
)

type Value interface{}
type Instruction uint

type VM struct {
	registers		RegisterBlock;
	memory			[]Value;
	code			[]Instruction;
	PC				uint;
}

func (vm *VM) Step() Instruction {
	if vm.PC < Len(vm.code) {
		instruction := code[vm.PC];
		vm.PC = vm.PC + 1;
		return instruction;
	} else {
		return NOOP;
	}
}

func (vm *VM) Skip(offset int) Instruction {
	if position < Len(vm.code) {
		vm.PC += offset;
		return code[vm.PC];
	} else {
		vm.PC = position;
		return NOOP;
	}
}

func (vm *VM) Goto(position uint) Instruction {
	if position < Len(vm.code) {
		vm.PC = position;
		return code[vm.PC];
	} else {
		vm.PC = position;
		return NOOP;
	}
}

func (vm *VM) Load(register, location uint) {
	if location < Len(vm.memory) { vm.registers.Load(register, vm.memory[location]) }
}

func (vm *VM) Store(register, location uint) {
	if location < Len(vm.memory) { vm.memory[location] = vm.registers.Read() }
}

func (vm *VM) Add(register uint, value Value) {
	if register < Len(vm.registers) { vm.register.Add(value) }
}