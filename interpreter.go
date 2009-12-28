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

var (
	stack			ForthStack;
	returnStack		ForthStack;
	memory			vector.IntVector;
	primitives		map[string] Primitive;
	words			map[string] string;
	variables		map[string] int;
)

func (p *Primitive) Execute() (result int, error int) {
	if stack.Len() < parameters { return 0, 1 }
	switch p.opcode {
	case 0:		// !
		value = int(stack.Pop());
		memory.Set(int(stack.Pop()), value);

	case 1:		// *
		stack.Multiply();

	case 2:		// +
		stack.Add();

	case 3:		// -
		stack.Subtract();

	case 4:		// .
		stack.Print();

	case 5:		// ."
		// TODO:

	case 6:		// /
		stack.Divide();

	case 7:		// /MOD
		stack.Divmod();

	case 8:		// 0<
		stack.Push(0);
		stack.LessThan();

	case 9:		// 0=
		stack.Push(0);
		stack.Equal();

	case 9.1:	// 0>
		stack.Push(0);
		stack.GreaterThan();

	case 10:	// <
		stack.LessThan();

	case 11:	// =
		stack.Equal();

	case 12:	// >
		stack.GreaterThan();

	case 13:	// ?
		fmt.Fprintf(os.Stdout, memory.At(stack.Pop()));

	case 14:	// @
		stack.Push(memory.At(stack.Pop()));

	case 15:	// ABS
		stack.Abs();

	case 16:	// AND
		stack.And();

	case 17:	// C@
		stack.Push(memory.At(stack.Pop()));

	case 18:	// CR
		fmt.Fprintln(os.Stdout);

	case 19:	// DROP
		stack.Pop();

	case 20:	// DUP
		stack.Push(stack.Last());

	case 21:	// EMIT
		stack.Emit();

	case 22:	// KEY
		code := [1]byte;
		os.Stdin.Read(code);
		stack.Push(code);

	case 23:	// MAX
		stack.Maximum();

	case 24:	// MIN
		stack.Minimum();

	case 25:	// MINUS
		stack.Minus();

	case 26:	// MOD
		stack.Mod();

	case 27:	// OR
		stack.Or();

	case 28:	// OVER
		stack.Over();

	case 29:	// SPACE
		fmt.Fprintf(os.Stdout, " ");

	case 30:	// SPACES
		fmt.Fprintf(os.Stdout, strings.Repeat(" ", stack.Pop()));

	case 31:	// SWAP
		stack.Swap();

	case 32:	// VARIABLE
		// TODO:

	case 33:	// XOR
		if (stack.Pop() ^ stack.Pop()) != FALSE { stack.Push(TRUE) }
		else { stack.Push(FALSE) }

	case 34:	// BEGIN
		// TODO:

	case 35:	// UNTIL
		// TODO:

	case 36:	// WHILE
		// TODO:

	case 37:	// REPEAT
		// TODO:

	case 38:	// IF
		// TODO:

	case 39:	// THEN
		// TODO:

	case 40:	// ELSE
		// TODO:

	case 41:	// FORTH
		// TODO:

	case 42:	// CLEAR
		// TODO:

	case 43:	// ROT
		stack.Rot();

	case 44:	// DO
		// TODO:

	case 45:	// LOOP
		// TODO:

	case 46:	// I
		// TODO:

	case 47:	// BYE
		os.Exit(0);

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