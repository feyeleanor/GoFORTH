package goforth

type Primitive struct {
	parameters		int;
	results			int;
	opcode			int;
	memory			[]byte;
}

func (p *Primitive) Execute(I Interpreter) bool {
	if stack.Len() < parameters { return 0, stack.UNDERFLOW }
	switch p.opcode {
	// !
	case 0:			if I.stack.Swap() { I.memory.Set(I.stack.Pop(), I.stack.Pop()); }
					else { goto ExecutionError }

	// *
	case 1:			I.stack.Multiply();

	// +
	case 2:			I.stack.Add();

	// -
	case 3:			I.stack.Subtract();

	// .
	case 4:			I.stack.Print();

	// ."
	case 5:			// TODO

	// /
	case 6:			I.stack.Divide();

	// /MOD
	case 7:			I.stack.Divmod();

	// 0<
	case 8:			I.stack.Push(0);
					I.stack.LessThan();

	// 0=
	case 9:			I.stack.Push(0);
					I.stack.Equal();

	// 0>
	case 10:		I.stack.Push(0);
					I.stack.GreaterThan();

	// <
	case 11:		I.stack.LessThan();

	// =
	case 12:		I.stack.Equal();

	// >
	case 13:		I.stack.GreaterThan();

	// ?
	case 14:		fmt.Fprintf(os.Stdout, I.memory.At(I.stack.Pop()));

	// @
	case 15:		I.stack.Push(I.memory.At(I.stack.Pop()));

	// ABS
	case 16:		I.stack.Abs();

	// AND
	case 17:		I.stack.And();

	// C@
	case 18:		I.stack.Push(I.memory.At(I.stack.Pop()));

	// CR
	case 19:		fmt.Fprintln(os.Stdout);

	// DROP
	case 20:		I.stack.Pop();

	// DUP
	case 21:		I.stack.Push(I.stack.Last());

	// EMIT
	case 22:		I.stack.Emit();

	// KEY
	case 23:		code := [1]byte;
					os.Stdin.Read(code);
					I.stack.Push(code);

	// MAX
	case 24:		I.stack.Maximum();

	// MIN
	case 25:		I.stack.Minimum();

	// MINUS
	case 26:		I.stack.Minus();

	// MOD
	case 27:		I.stack.Mod();

	// OR
	case 28:		I.stack.Or();

	// OVER
	case 29:		I.stack.Over();

	// SPACE
	case 30:		fmt.Fprintf(os.Stdout, " ");

	// SPACES
	case 31:		fmt.Fprintf(os.Stdout, strings.Repeat(" ", I.stack.Pop()));

	// SWAP
	case 32:		I.stack.Swap();

	// VARIABLE
	case 33:		// TODO

	// XOR
	case 34:		I.stack.Xor();

	// BEGIN
	case 35:		I.returnStack.Push(I.programCounter);

	// UNTIL
	case 36:		if I.stack.IsFalse() { I.Goto(returnStack.Pop()) }

	// WHILE
	case 37:		if I.stack.IsFalse() { I.returnStack.Push(stack.TRUE) }
					else { I.returnStack.Push(stack.FALSE) }

	// REPEAT
	case 38:		if I.returnStack.IsFalse() { I.Goto(I.returnStack.Pop()) }
					else { I.returnStack.Pop() }

	// IF
	case 39:		// TODO
					// Skip to ELSE if condition false
					if I.stack.IsFalse() { I.Goto(ADDRESS_OF_ELSE_STATEMENT) }

	// THEN
	case 40:		// TODO
					// Clean up return Stack

	// ELSE
	case 41:		// TODO
					// Skip to THEN

	// FORTH
	case 42:		// TODO

	// CLEAR
	case 43:		// TODO

	// ROT
	case 44:		I.stack.Rot();

	// DO
	case 45:		I.returnStack.Push(I.PC);
					I.stack.Swap();
					I.returnStack.Push(I.stack.Pop());
					I.returnStack.Push(I.stack.Pop());

	// LOOP
	case 46:		I.returnStack.Push(1);
					I.returnStack.Add();
					index := I.returnStack.Top();
					I.returnStack.LessThan();
					if I.returnStack.IsTrue() { I.Goto(I.returnStack.Pick(2)) }
					else { I.returnStack.Drop() }

	// +LOOP
	case 47:		I.returnStack.Push(I.stack.Pop());
					I.returnStack.Add();
					index := I.returnStack.Top();
					I.returnStack.LessThan();
					if I.returnStack.IsTrue() { I.Goto(I.returnStack.Pick(2)) }
					else { I.returnStack.Drop() }

	// I
	case 48:		I.stack.Push(I.returnStack.Top());

	// BYE
	case 49:		os.Exit(0);

	default:
	}
	return true;

ExecutionError:
	return false;
}