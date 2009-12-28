package goforth

import (
	"container/vector"
	"fmt";
	"os";
	"strings";
)

type Primitive struct {
	parameters		int;
	results			int;
	opcode			int;
	memory			[]byte;
}

type Interpreter struct {
	stack			ForthStack;
	returnStack		ForthStack;
	memory			vector.IntVector;
	primitives		map[string] Primitive;
	words			map[string] []byte;
	variables		map[string] int;
}

func (p *Primitive) Execute(runtime Interpreter) (result int, error int) {
	if stack.Len() < parameters { return 0, stack.UNDERFLOW }
	switch p.opcode {
	case 0:			runtime.memory.Set(runtime.stack.SwapPop());						// !
	case 1:			runtime.stack.Multiply();											// *
	case 2:			runtime.stack.Add();												// +
	case 3:			runtime.stack.Subtract();											// -
	case 4:			runtime.stack.Print();												// .
	case 5:			// TODO																// ."
	case 6:			runtime.stack.Divide();												// /
	case 7:			runtime.stack.Divmod();												// /MOD
	case 8:			runtime.stack.Push(0);												// 0<
					runtime.stack.LessThan();
	case 9:			runtime.stack.Push(0);												// 0=
					runtime.stack.Equal();
	case 10:		runtime.stack.Push(0);												// 0>
					runtime.stack.GreaterThan();
	case 11:		runtime.stack.LessThan();											// <
	case 12:		runtime.stack.Equal();												// =
	case 13:		runtime.stack.GreaterThan();										// >
	case 14:		fmt.Fprintf(os.Stdout, runtime.memory.At(runtime.stack.Pop()));		// ?
	case 15:		runtime.stack.Push(runtime.memory.At(runtime.stack.Pop()));			// @
	case 16:		runtime.stack.Abs();												// ABS
	case 17:		runtime.stack.And();												// AND
	case 18:		runtime.stack.Push(runtime.memory.At(runtime.stack.Pop()));			// C@
	case 19:		fmt.Fprintln(os.Stdout);											// CR
	case 20:		runtime.stack.Pop();												// DROP
	case 21:		runtime.stack.Push(runtime.stack.Last());							// DUP
	case 22:		runtime.stack.Emit();												// EMIT
	case 23:		code := [1]byte;													// KEY
					os.Stdin.Read(code);
					runtime.stack.Push(code);
	case 24:		runtime.stack.Maximum();											// MAX
	case 25:		runtime.stack.Minimum();											// MIN
	case 26:		runtime.stack.Minus();												// MINUS
	case 27:		runtime.stack.Mod();												// MOD
	case 28:		runtime.stack.Or();													// OR
	case 29:		runtime.stack.Over();												// OVER
	case 30:		fmt.Fprintf(os.Stdout, " ");										// SPACE
	case 31:		fmt.Fprintf(os.Stdout, strings.Repeat(" ", runtime.stack.Pop()));	// SPACES
	case 32:		runtime.stack.Swap();												// SWAP
	case 33:		// TODO																// VARIABLE
	case 34:		runtime.stack.Xor();												// XOR
	case 35:		// TODO																// BEGIN
	case 36:		// TODO																// UNTIL
	case 37:		// TODO																// WHILE
	case 38:		// TODO																// REPEAT
	case 39:		// TODO																// IF
	case 40:		// TODO																// THEN
	case 41:		// TODO																// ELSE
	case 42:		// TODO																// FORTH
	case 43:		// TODO																// CLEAR
	case 44:		runtime.stack.Rot();												// ROT
	case 45:		// TODO																// DO
	case 46:		// TODO																// LOOP
	case 47:		// TODO																// I
	case 48:		os.Exit(0);															// BYE
	default:
  }
}

func init() {
	stack = NewStack();
	returnStack = NewStack();
	memory = new(vector.IntVector);
	words = make(map[string] string);
	variables = make(map[string] int);
	primitives = make(map[string] Primitive);
	primitives["!"] = Primitive{2, 0, 0};
	primitives["*"] = Primitive{2, 1, 1};
	primitives["+"] = Primitive{2, 1, 2};
	primitives["-"] = Primitive{2, 1, 3};
	primitives["."] = Primitive{1, 0, 4};
	primitives["?"] = Primitive{0, 0, 5};
	primitives["/"] = Primitive{2, 1, 6};
	primitives["/MOD"] = Primitive{2, 2, 7};
	primitives["0<"] = Primitive{1, 1, 8};
	primitives["0="] = Primitive{1, 0, 9};
	primitives["0>"] = Primitive{1, 0, 10};
	primitives["<"] = Primitive{2, 1, 11};
	primitives["="] = Primitive{2, 1, 12};
	primitives[">"] = Primitive{2, 1, 13};
	primitives["?"] = Primitive{1, 0, 14};
	primitives["@"] = Primitive{1, 1, 15};
	primitives["ABS"] = Primitive{1, 1, 16};
	primitives["AND"] = Primitive{2, 1, 17};
	primitives["C@"] = Primitive{1, 1, 18};
	primitives["CR"] = Primitive{0, 0, 19};
	primitives["DROP"] = Primitive{1, 0, 20};
	primitives["DUP"] = Primitive{1, 2, 21};
	primitives["EMIT"] = Primitive{1, 0, 22};
	primitives["KEY"] = Primitive{0, 1, 23};
	primitives["MAX"] = Primitive{2, 1, 24};
	primitives["MIN"] = Primitive{2, 1, 25};
	primitives["MINUS"] = Primitive{1, 1, 26};
	primitives["MOD"] = Primitive{2, 1, 27};
	primitives["OR"] = Primitive{2, 1, 28};
	primitives["OVER"] = Primitive{2, 3, 29};
	primitives["SPACE"] = Primitive{0, 0, 30};
	primitives["SPACES"] = Primitive{1, 2, 31};
	primitives["SWAP"] = Primitive{2, 2, 32};
	primitives["VARIABLE"] = Primitive{1, 0, 33};
	primitives["XOR"] = Primitive{2, 1, 34};
	primitives["BEGIN"] = Primitive{0, 0, 35};
	primitives["UNTIL"] = Primitive{1, 0, 36};
	primitives["WHILE"] = Primitive{1, 0, 37};
	primitives["REPEAT"] = Primitive{0, 0, 38};
	primitives["IF"] = Primitive{1, 0, 39};
	primitives["THEN"] = Primitive{0, 0, 40};
	primitives["ELSE"] = Primitive{0, 0, 41};
	primitives["FORTH"] = Primitive{0, 0, 42};
	primitives["CLEAR"] = Primitive{0, 0, 43};
	primitives["ROT"] = Primitive{3, 3, 44};
	primitives["DO"] = Primitive{2, 0, 45};
	primitives["LOOP"] = Primitive{0, 0, 46};
	primitives["I"] = Primitive{0, 1, 47};
	primitives["BYE"] = Primitive{0, 0, 48};
}