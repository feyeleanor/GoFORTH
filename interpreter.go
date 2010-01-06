package goforth

import (
	"container/vector"
	"fmt";
	"os";
	"bufio";
	"strings";
)

const(
	BANNER = "GoFORTH v0.0a";
)

type Program struct {
	memory			vector.IntVector;
	instructions	[]byte;
	PC				uint;
}

func (p *Program) Goto(position uint) {
	p.PC = position;
}

const(
	INTERPRETING = iota;
	COMPILING;
)

type Interpreter struct {
	Program;
	stack			ForthStack;
	returnStack		ForthStack;
	primitives		map[string] Primitive {
						"!": Primitive{2, 0, 0},			"*": Primitive{2, 1, 1},		"+": Primitive{2, 1, 2},
						"-": Primitive{2, 1, 3},			".": Primitive{1, 0, 4},		"?": Primitive{0, 0, 5},
						"/": Primitive{2, 1, 6},			"/MOD": Primitive{2, 2, 7},		"0<": Primitive{1, 1, 8},
						"0=": Primitive{1, 0, 9},			"0>": Primitive{1, 0, 10},		"<": Primitive{2, 1, 11},
						"=": Primitive{2, 1, 12},			">": Primitive{2, 1, 13},		"?": Primitive{1, 0, 14},
						"@": Primitive{1, 1, 15},			"ABS": Primitive{1, 1, 16},		"AND": Primitive{2, 1, 17},
						"C@": Primitive{1, 1, 18},			"CR": Primitive{0, 0, 19},		"DROP": Primitive{1, 0, 20},
						"DUP": Primitive{1, 2, 21},			"EMIT": Primitive{1, 0, 22},	"KEY": Primitive{0, 1, 23},
						"MAX": Primitive{2, 1, 24},			"MIN": Primitive{2, 1, 25},		"MINUS": Primitive{1, 1, 26},
						"MOD": Primitive{2, 1, 27},			"OR": Primitive{2, 1, 28},		"OVER": Primitive{2, 3, 29},
						"SPACE": Primitive{0, 0, 30},		"SPACES": Primitive{1, 2, 31},	"SWAP": Primitive{2, 2, 32},
						"VARIABLE": Primitive{1, 0, 33},	"XOR": Primitive{2, 1, 34},		"BEGIN": Primitive{0, 0, 35},
						"UNTIL": Primitive{1, 0, 36},		"WHILE": Primitive{1, 0, 37},	"REPEAT": Primitive{0, 0, 38},
						"IF": Primitive{1, 0, 39},			"THEN": Primitive{0, 0, 40},	"ELSE": Primitive{0, 0, 41},
						"FORTH": Primitive{0, 0, 42},		"CLEAR": Primitive{0, 0, 43},	"ROT": Primitive{3, 3, 44},
						"DO": Primitive{2, 0, 45},			"LOOP": Primitive{0, 0, 46},	"+LOOP": Primitive{0, 0, 47},
						"I": Primitive{0, 1, 48},			"R@": Primitive{0, 1, 48},		"BYE": Primitive{0, 0, 49},
					}
	words			map[string] []byte;
	variables		map[string] int;
	state			int;
}


func NewInterpreter(program *[]byte) {
	i := new(Interpreter);
	i.stack = NewStack();
	i.returnStack = NewStack();
	i.memory = new(vector.IntVector);
	i.words = make(map[string] string);
	i.variables = make(map[string] int);
	i.state = INTERPRETING;
}

func (I *Interpreter) Read(in bufio.Reader) String {
	switch I.state {
	case COMPILING:		seeking := ';'
	default:			seeking : = '\n'
	}
	if str, e := in.ReadString(seeking); e {
		if e != os.EOF {
			fmt.Printf("Error reading input: %s\n", e);
			str = nil;
			// TODO Integrate this with the interpreter's normal error state mechanism
		}
	}
	return str;
}

func (I *Interpreter) Evaluate(statement *String) {
	// TODO
	// Split the string into tokens and then execute the tokens individually
	// Compile new words as required
}

func (I *Interpreter) Print() {
	// TODO
}

func (I *Interpreter) Run() {
	in := bufio.NewReader(os.Stdin);
	for {
		if statement := I.Read(in); statement {
			if result := I.Evaluate(statement); result {
				I.Print(result);
			}
		}
	}
}