package golightly

import (
	"container/vector"
	"fmt";
	"os";
	"bufio";
	"strings";
)

type OpCode struct {
	code		int;
	a, b, c		int;
	d			*interface{};
}
type Instruction func(o OpCode)

type RegisterBlock struct {
	PC				int;					//	program counter
	SP				int						//	stack pointer
	MP				int;					//	memory pointer
	RV				[REGISTER_COUNT]int;	//	16 general purpose registers
}
func (block *RegisterBlock) LD(r1, r2 int) { block.RV[r1] = block.RV[r2] }
func (block *RegisterBlock) LDC(r, v int) { block.RV[r] = v }
func (block *RegisterBlock) ADD(r1, r2 int) { block.RV[r1] += block.RV[r2] }
func (block *RegisterBlock) SUB(r1, r2 int) { block.RV[r1] -= block.RV[r2] }
func (block *RegisterBlock) MUL(r1, r2 int) { block.RV[r1] *= block.RV[r2] }
func (block *RegisterBlock) DIV(r1, r2 int) { block.RV[r1] /= block.RV[r2] }
func (block *RegisterBlock) AND(r1, r2 int) { block.RV[r1] &= block.RV[r2] }
func (block *RegisterBlock) OR(r1, r2 int) { block.RV[r1] |= block.RV[r2] }
func (block *RegisterBlock) XOR(r1, r2 int) { block.RV[r1] ^= block.RV[r2] }

func (block *RegisterBlock) ADDC(r, v int) { block.RV[r] += v }
func (block *RegisterBlock) SUBC(r, v int) { block.RV[r] -= v }
func (block *RegisterBlock) MULC(r, v int) { block.RV[r] *= v }
func (block *RegisterBlock) DIVC(r, v int) { block.RV[r] /= v }
func (block *RegisterBlock) ANDC(r, v int) { block.RV[r] /= v }
func (block *RegisterBlock) ORC(r, v int) { block.RV[r] /= v }
func (block *RegisterBlock) XORC(r, v int) { block.RV[r] /= v }

type VM struct {
	registers		RegisterBlock;
	call_stack		vector.IntVector;
	invalidPC		int;
	program			[]OpCode;
	running			bool;
	microcode		vector.Vector;
}
func NewVM(register_count uint) *VM {
	vm := VM{ registers: new(RegisterBlock), call_stack: vector.NewIntVector(16) }
	ops := [
		func (o OpCode) {},											//	NOOP
		func (o OpCode) { vm.registers.PC += o.a },					//	JUMP	n
		func (o OpCode) { vm.PushRegisters(); vm.jump(o.a) },		//	CALL	n
		func (o OpCode) { vm.PopRegisters() },						//	RET
		func (o OpCode) { vm.registers.LD(o.a, o.b) },				//	LD		r1, r2
		func (o OpCode) { vm.registers.LDC(o.a, o.b) },				//	LDC		r, v
		func (o OpCode) { vm.registers.ADD(o.a, o.b) },				//	ADD		r1, r2
		func (o OpCode) { vm.registers.ADDC(o.a, o.b) },			//	ADDC	r, v
		func (o OpCode) { vm.registers.SUB(o.a, o.b) },				//	SUB		r1, r2
		func (o OpCode) { vm.registers.SUBC(o.a, o.b) },			//	SUBC	r, v
		func (o OpCode) { vm.registers.MUL(o.a, o.b) },				//	MUL		r1, r2
		func (o OpCode) { vm.registers.MULC(o.a, o.b) },			//	MULC	r, v
		func (o OpCode) { vm.registers.DIV(o.a, o.b) },				//	DIV		r1, r2
		func (o OpCode) { vm.registers.DIVC(o.a, o.b) },			//	DIVC	r, v
		func (o OpCode) { vm.registers.AND(o.a, o.b) },				//	AND		r1, r2
		func (o OpCode) { vm.registers.ANDC(o.a, o.b) },			//	ANDC	r, v
		func (o OpCode) { vm.registers.OR(o.a, o.b) },				//	OR		r1, r2
		func (o OpCode) { vm.registers.ORC(o.a, o.b) },				//	ORC		r, v
		func (o OpCode) { vm.registers.XOR(o.a, o.b) },				//	XOR		r1, r2
		func (o OpCode) { vm.registers.XORC(o.a, o.b) }				//	XORC	r, v
	]
	vm.microcode = vector.Vector{ a: ops }
	return &VM;
}
func (vm *VM) PushRegisters() {
	vm.call_stack.Push(vm.registers);
	vm.registers.SP++;
}
func (vm *VM) PopRegisters() {
	if vm.registers.SP > -1 { vm.registers = vm.call_stack.Pop() }
	else { vm.running = false }
}
func (vm *VM) ValidPC() bool { return (vm.registers[PC] < vm.invalidPC) && vm.running }
func (vm *VM) Step() { vm.registers.PC++ }
func (vm *VM) Exec(o OpCode) { vm.microcode.At(o.action)(o) }
func (vm *VM) Load(program []Operands) {
	vm.program = program;
	vm.PC = 0;
	vm.invalidPC = Len(program);
	vm.running = false;
}
func (vm *VM) CurrentOp() OpCode {
	if vm.ValidPC { return program[vm.registers.PC] }
	else { return nil }
}
func (vm *VM) Run() {
	for i := vm.CurrentOp() {
		vm.Exec(i);
		vm.Step();
	}
}