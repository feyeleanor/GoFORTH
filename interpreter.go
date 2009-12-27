package goforth

import (
  "container/vector";
  "fmt";
  "os";
  "strings";
)

type Primitive struct {
  parameters	int;
  results			int;
  opcode			int;
  memory			[]byte;
}

var (
  stack				forthStack;
  returnStack	forthStack;
  memory			vector.IntVector;
  primitives	map[string] Primitive;
  words				map[string] string;
  variables		map[string] int;
)

func Primitive.execute() (result int, error int) {
	if stack.Len() < parameters { return 0, 1 }
  switch primitive.opcode {
  case 0:		// !
  	value = int(stack.Pop());
		memory.Set(int(stack.Pop()), value);

  case 1:		// *
		stack.multiply();

  case 2:		// +
		stack.add();

  case 3:		// -
		stack.subtract();

  case 4:		// .
  	stack.print();

  case 5:		// ."
  	// TODO:

  case 6:		// /
  	stack.divide();

  case 7:		// /MOD
  	stack.divmod();

  case 8:		// 0<
  	stack.Push(0);
  	stack.less_than();

  case 9:		// 0=
  	stack.Push(0);
  	stack.equal_to();

	case 9.1:	// 0>
		stack.Push(0);
		stack.greater_than();

  case 10:	// <
  	stack.less_than();

  case 11:	// =
  	stack.equal_to();

  case 12:	// >
  	stack.greater_than();

	case 13:	// ?
		fmt.Fprintf(os.Stdout, memory.At(stack.Pop()));

	case 14:	// @
		stack.Push(memory.At(stack.Pop()));

	case 15:	// ABS
		stack.abs();

	case 16:	// AND
		stack.and();

	case 17:	// C@
		stack.Push(memory.At(stack.Pop()));

	case 18:	// CR
		fmt.Fprintln(os.Stdout);

	case 19:	// DROP
		stack.Pop();

	case 20:	// DUP
		stack.Push(stack.Last());

	case 21:	// EMIT
		stack.emit();

	case 22:	// KEY
		code := [1]byte;
		os.Stdin.Read(code);
		stack.Push(code);

	case 23:	// MAX
		x, y := stack.Pop(), stack.Pop();
		if x > y { stack.Push(x) }
		else { stack.Push(y) }

	case 24:	// MIN
		x, y := stack.Pop(), stack.Pop();
		if x < y { stack.Push(x) }
		else { stack.Push(y) }

	case 25:	// MINUS
		stack.Push(-stack.Pop());

	case 26:	// MOD
		stack.mod();

	case 27:	// OR
		stack.or();

	case 28:	// OVER
		stack.over();

	case 29:	// SPACE
		fmt.Fprintf(os.Stdout, " ");

	case 30:	// SPACES
		fmt.Fprintf(os.Stdout, strings.Repeat(" ", stack.Pop()));

	case 31:	// SWAP
		stack.Swap(stack.Len() - 1, stack.Len() - 2);

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
		top_of_stack := stack.Len() - 1;
		stack.Push(stack.At(top_of_stack - 1));

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
	stack = new(vector.IntVector);
	returnStack = new(vector.IntVector);
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
	primitives["<"] = Primitive{2, 1, 10};
	primitives["="] = Primitive{2, 1, 11};
	primitives[">"] = Primitive{2, 1, 12};
	primitives["?"] = Primitive{1, 0, 13};
	primitives["@"] = Primitive{1, 1, 14};
	primitives["ABS"] = Primitive{1, 1, 15};
	primitives["AND"] = Primitive{2, 1, 16};
	primitives["C@"] = Primitive{1, 1, 17};
	primitives["CR"] = Primitive{0, 0, 18};
	primitives["DROP"] = Primitive{1, 0, 19};
	primitives["DUP"] = Primitive{1, 2, 20};
	primitives["EMIT"] = Primitive{1, 0, 21};
	primitives["KEY"] = Primitive{0, 1, 22};
	primitives["MAX"] = Primitive{2, 1, 23};
	primitives["MIN"] = Primitive{2, 1, 24};
	primitives["MINUS"] = Primitive{1, 1, 25};
	primitives["MOD"] = Primitive{2, 1, 26};
	primitives["OR"] = Primitive{2, 1, 27};
	primitives["OVER"] = Primitive{2, 3, 28};
	primitives["SPACE"] = Primitive{0, 0, 29};
	primitives["SPACES"] = Primitive{1, 2, 30};
	primitives["SWAP"] = Primitive{2, 2, 31};
	primitives["VARIABLE"] = Primitive{1, 0, 32};
	primitives["XOR"] = Primitive{2, 1, 33};
	primitives["BEGIN"] = Primitive{0, 0, 34};
	primitives["UNTIL"] = Primitive{1, 0, 35};
	primitives["WHILE"] = Primitive{1, 0, 36};
	primitives["REPEAT"] = Primitive{0, 0, 37};
	primitives["IF"] = Primitive{1, 0, 38};
	primitives["THEN"] = Primitive{0, 0, 39};
	primitives["ELSE"] = Primitive{0, 0, 40};
	primitives["FORTH"] = Primitive{0, 0, 41};
	primitives["CLEAR"] = Primitive{0, 0, 42};
	primitives["ROT"] = Primitive{3, 3, 43};
	primitives["DO"] = Primitive{2, 0, 44};
	primitives["LOOP"] = Primitive{0, 0, 45};
	primitives["I"] = Primitive{0, 1, 46};
	primitives["BYE"] = Primitive{0, 0, 47};
}