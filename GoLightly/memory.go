package golightly

type MemoryBlock []int
type Operands struct {
	a, b, c, d		int;
}

func (block *MemoryBlock) Operation(action uint, operands Operands) {
	switch action {
	case ADD:					block[operands.a] += block[operands.b]
	case ADD_VALUE:				block[operands.a] += operands.b
	case SUBTRACT:				block[operands.a] -= block[operands.b]
	case SUBTRACT_VALUE:		block[operands.a] -= operands.b
	case MULTIPLY:				block[operands.a] *= block[operands.b]
	case MULTIPLY_VALUE:		block[operands.a] *= operands.b
	case DIVIDE:				block[operands.a] /= block[operands.b]
	case DIVIDE_VALUE:			block[operands.a] /= operands.b
	case AND:					block[operands.a] &= block[operands.b]
	case AND_VALUE:				block[operands.a] &= operands.b
	case OR:					block[operands.a] |= block[operands.b]
	case OR_VALUE:				block[operands.a] |= operands.b
	case XOR:					block[operands.a] ^= block[operands.b]
	case XOR_VALUE:				block[operands.a] ^= operands.b
	case COPY:					block[operands.a] = block[operands.b]
	case COPY_VALUE:			block[operands.a] = operands.b
	default:
	}
}

func (block *MemoryBlock) Copy(source, target uint) {
	block_size := Len(MemoryBlock);
	if source < block_size && target < block_size { block[target] = block[source] }
}

func (block *MemoryBlock) Add(source, target uint) {
	block_size := Len(MemoryBlock);
	if source < block_size && target < block_size { block[target] += block[source] }
}