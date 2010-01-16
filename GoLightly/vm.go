package golightly

import (
	"container/vector"
	"fmt";
	"os";
	"bufio";
	"strings";
)

const(
	NOOP = iota;
	JUMP;				//	JMP n
	CALL;				//	CAL n
	RETURN;				//	RET
	LOAD;				//	LOD r, v			loads v into r
	COPY;				//	CPY r1, r2			copies contents of r2 to r1
	ADD;				//	ADD r1, r2			add contents of r2 to r1
	ADD_VALUE;			//	ADDC r1, v			add constant of v to r1
	SUBTRACT;			//	SUB r1, r2			subtract contents of r2 from r1
	SUBTRACT_VALUE;		//	SUBC r1, v			subtract constant v from r1
	MULTIPLY;			//	MUL r1, r2			multiply contents of r1 by r2
	MULTIPLY_VALUE;		//	MULC r1, v			multiply contents of r1 by constant v
	DIVIDE;				//	DIV r1, r2			divide contents of r1 by r2
	DIVIDE_VALUE;		//	DIVC r1, v			divide contents of r1 by constant v
	AND;				//	AND r1, r2			logical and contents of r1 and r2
	AND_VALUE;			//	ANDC r1, v			logical and contents of r1 and constant v
	OR;					//	OR  r1, r2			logical or contents of r1 and r2
	OR_VALUE;			//	ORC r1, v			logical or contents of r1 and constant v
	XOR;				//	XOR r1, r2			logical xor contents of r1 and r2
	XOR_VALUE;			//	XORC r1, v			logical xor contents of r1 and constant v
)

type OpCode struct {
	action		int;
	a, b, c		int;
	d			*interface{};
}

type MemoryBlock	vector.IntVector;
func (block *MemoryBlock) Exec(o OpCode) {
	switch o.action {
	case ADD:					block[o.a] += block[o.b]
	case ADD_VALUE:				block[o.a] += o.b
	case SUBTRACT:				block[o.a] -= block[o.b]
	case SUBTRACT_VALUE:		block[o.a] -= o.b
	case MULTIPLY:				block[o.a] *= block[o.b]
	case MULTIPLY_VALUE:		block[o.a] *= o.b
	case DIVIDE:				block[o.a] /= block[o.b]
	case DIVIDE_VALUE:			block[o.a] /= o.b
	case AND:					block[o.a] &= block[o.b]
	case AND_VALUE:				block[o.a] &= o.b
	case OR:					block[o.a] |= block[o.b]
	case OR_VALUE:				block[o.a] |= o.b
	case XOR:					block[o.a] ^= block[o.b]
	case XOR_VALUE:				block[o.a] ^= o.b
	case COPY:					block[o.a] = block[o.b]
	case COPY_VALUE:			block[o.a] = o.b
	default:
	}
}

type RegisterBlock struct {
	PC				int;					//	program counter
	SP				int						//	stack pointer
	MP				int;					//	memory pointer
	RV				[REGISTER_COUNT]int;	//	16 general purpose registers
}
func NewRegisterBlock(count int) RegisterBlock {
	return RegisterBlock{RV: vector.NewIntVector(count) }
}
func (block *RegisterBlock) Load(r, v int) bool {
	if r < Len(block.RV) { block.RV[r] = v }
	else { return false }
	return true;
}
func (block *RegisterBlock) Copy(r1, r2 int) bool {
	if r1 < Len(block.RV) && r2 < Len(block.RV) { block.RV[r1] = block.RV[r2] }
	else { return false }
	return true;
}
func (block *RegisterBlock) Add(r, v int) bool {
	if r < Len(block.RV) { block.RV[r] += v }
	else { return false }
	return true;
}
func (block *RegisterBlock) Subtract(r, v int) bool {
	if r < Len(block.RV) { block.RV[r] -= v }
	else { return false }
	return true;
}
func (block *RegisterBlock) Multiply(r, v int) bool {
	if r < Len(block.RV) { block.RV[r] *= v }
	else { return false }
	return true;
}
func (block *RegisterBlock) Divide(r, v int) bool {
	if r < Len(block.RV) { block.RV[r] /= v }
	else { return false }
	return true;
}
func (block *RegisterBlock) And(r, v int) bool {
	if r < Len(block.RV) { block.RV[r] &= v }
	else { return false }
	return true;
}
func (block *RegisterBlock) Or(r, v int) bool {
	if r < Len(block.RV) { block.RV[r] |= v }
	else { return false }
	return true;
}
func (block *RegisterBlock) Xor(r, v int) bool {
	if r < Len(block.RV) { block.RV[r] ^= v }
	else { return false }
	return true;
}

type VM struct {
	registers		RegisterBlock;
	call_stack		vector.IntVector;
	invalidPC		int;
	program			[]OpCode;
	running			bool;
}
func NewVM(register_count uint) VM {
	return VM{	registers:		NewRegisterBlock(register_count),
				call_stack:		vector.NewIntVector(16)	}
}
func (vm *VM) ValidPC() bool { return (vm.registers[PC] < vm.invalidPC) && vm.running }
func (vm *VM) Exec(o OpCode) {
	switch o.action {
	case JUMP:
		vm.registers.PC += o.a;
	case CALL:
		vm.call_stack.Push(vm.registers);
		vm.registers.SP++;
		vm.registers.PC = o.a;
	case RETURN:
		if vm.registers.SP > -1 { vm.registers = vm.call_stack.Pop() }
		else { vm.running = false }
	case LOAD:
		if !vm.registers.Load(o.a, o.b) { vm.running = false }
	case COPY:
		if !vm.registers.Copy(o.a, o.b) { vm.running = false }
	case ADD_VALUE:
		if !vm.registers.Add(o.a, o.b) { vm.running = false }
	case SUBTRACT_VALUE:
		if !vm.registers.Subtract(o.a, o.b) { vm.running = false }
	case MULTIPLY_VALUE:
		if !vm.registers.Multiply(o.a, o.b) { vm.running = false }
	case DIVIDE_VALUE:
		if !vm.registers.Divide(o.a, o.b) { vm.running = false }
	case AND_VALUE:
		if !vm.registers.And(o.a, o.b) { vm.running = false }
	case OR_VALUE:
		if !vm.registers.Or(o.a, o.b) { vm.running = false }
	case XOR_VALUE:
		if !vm.registers.Xor(o.a, o.b) { vm.running = false }
	case NOOP:
	}
}
func (vm *VM) Load(program []Operands) {
	vm.program = program;
	vm.PC = 0;
	vm.invalidPC = Len(program);
	vm.running = false;
}
func (vm *VM) Step() OpCode {
	if vm.ValidPC { return program[vm.registers.PC] }
	else { return nil }
}
func (vm *VM) Run() {
	for i := vm.Step(); i { vm.Exec(i) }
}