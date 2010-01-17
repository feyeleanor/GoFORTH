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
	SP				int;					//	stack pointer
	MP				*[]byte;				//	memory pointer
	RV				[REGISTER_COUNT]int;	//	16 general purpose registers
}
func (block *RegisterBlock) CSELECT(v int) { block.MP = &[]byte(v) }
func (block *RegisterBlock) SELECT(r int) { block.CSELECT(block.RV[r]) }

func (block *RegisterBlock) CLD(r, v int) { block.RV[r] = v }
func (block *RegisterBlock) LD(r1, r2 int) { block.CLD(r1, block.RV[r2]) }
func (block *RegisterBlock) ILD(r1, r2 int) { block.CLD(r1, block.RV[block.MP[r2]]) }

func (block *RegisterBlock) CADD(r, v int) { block.RV[r] += v }
func (block *RegisterBlock) ADD(r1, r2 int) { block.ADDC(r1, block.RV[r2]) }

func (block *RegisterBlock) CSUB(r, v int) { block.RV[r] -= v }
func (block *RegisterBlock) SUB(r1, r2 int) { block.SUBC(r1, block.RV[r2]) }

func (block *RegisterBlock) CMUL(r, v int) { block.RV[r] *= v }
func (block *RegisterBlock) MUL(r1, r2 int) { block.SUBC(r1, block.RV[r2]) }

func (block *RegisterBlock) CDIV(r, v int) { block.RV[r] /= v }
func (block *RegisterBlock) DIV(r1, r2 int) { block.DIVC(r1, block.RV[r2]) }

func (block *RegisterBlock) CAND(r, v int) { block.RV[r] /= v }
func (block *RegisterBlock) AND(r1, r2 int) { block.ANDC(r1, block.RV[r2]) }

func (block *RegisterBlock) COR(r, v int) { block.RV[r] /= v }
func (block *RegisterBlock) OR(r1, r2 int) { block.ORC(r1, block.RV[r2]) }

func (block *RegisterBlock) CXOR(r, v int) { block.RV[r] /= v }
func (block *RegisterBlock) XOR(r1, r2 int) { block.XORC(r1, block.RV[r2]) }

type VM struct {
	registers		RegisterBlock;
	call_stack		vector.Vector;
	invalidPC		int;
	program			[]OpCode;
	running			bool;
	microcode		vector.Vector;
}
func NewVM(register_count uint) *VM {
	vm := VM{ registers: new(RegisterBlock), call_stack: vector.New(16), memory_map: vector.New(128) }
	ops := [
		func (o OpCode) {},															//	NOOP
		func (o OpCode) { vm.registers.PC += o.a },									//	JUMP	n
		func (o OpCode) {															//	CALL	n
 			vm.PushRegisters();
			vm.jump(o.a);
		},
		func (o OpCode) { vm.PopRegisters() },										//	RET
		func (o OpCode) { vm.registers.CLD(o.a, uintptr(vm.AllocateMemory(o.b))) }	//	GRAB	r, n
		func (o OpCode) { vm.registers.SELECT(o.a) },								//	SELECT	r
		func (o OpCode) { vm.registers.CSELECT(o.a) },								//	CSELECT	n
		func (o OpCode) { vm.registers.LD(o.a, o.b) },								//	LD		r1, r2
		func (o OpCode) { vm.registers.CLD(o.a, o.b) },								//	CLD		r, v
		func (o OpCode) { vm.registers.ILD(o.a, o.b) },								//	ILD		r1, r2
		func (o OpCode) { vm.registers.ADD(o.a, o.b) },								//	ADD		r1, r2
		func (o OpCode) { vm.registers.CADD(o.a, o.b) },							//	CADD	r, v
		func (o OpCode) { vm.registers.SUB(o.a, o.b) },								//	SUB		r1, r2
		func (o OpCode) { vm.registers.CSUB(o.a, o.b) },							//	CSUB	r, v
		func (o OpCode) { vm.registers.MUL(o.a, o.b) },								//	MUL		r1, r2
		func (o OpCode) { vm.registers.CMUL(o.a, o.b) },							//	CMUL	r, v
		func (o OpCode) { vm.registers.DIV(o.a, o.b) },								//	DIV		r1, r2
		func (o OpCode) { vm.registers.CDIV(o.a, o.b) },							//	CDIV	r, v
		func (o OpCode) { vm.registers.AND(o.a, o.b) },								//	AND		r1, r2
		func (o OpCode) { vm.registers.CAND(o.a, o.b) },							//	CAND	r, v
		func (o OpCode) { vm.registers.OR(o.a, o.b) },								//	OR		r1, r2
		func (o OpCode) { vm.registers.COR(o.a, o.b) },								//	COR		r, v
		func (o OpCode) { vm.registers.XOR(o.a, o.b) },								//	XOR		r1, r2
		func (o OpCode) { vm.registers.CXOR(o.a, o.b) },							//	CXOR	r, v
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
func (vm *VM) AllocateMemory(len int) int {
	m := make([]int, len);
	vm.memory_map.Push(&m);	
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