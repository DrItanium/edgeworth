// Arithmetic operations, an arithmetic operation is a half instruction
package edgeworth

import "fmt"

type SimpleArithmeticOperation func(Word, Word) Word

/* List of operations */
const (
	ArithmeticOperationAdd = iota
	ArithmeticOperationSubtract
	ArithmeticOperationMultiply
	ArithmeticOperationDivide
	ArithmeticOperationRemainder
	ArithmeticOperationShiftLeft
	ArithmeticOperationShiftRight
	ArithmeticOperationBitwiseAnd
	ArithmeticOperationBitwiseOr
	ArithmeticOperationBitwiseNot
	ArithmeticOperationBitwiseXor
)

/* Some helpful constants */
const (
	// If the opcode is greater than this value then we are in immediate mode
	ImmediateFormOffset = 16
)

var simpleArithmeticOperations = []SimpleArithmeticOperation{
	func(x, y Word) Word { return x + y },
	func(x, y Word) Word { return x - y },
	func(x, y Word) Word { return x * y },
	func(x, y Word) Word { return x / y },
	func(x, y Word) Word { return x % y },
	func(x, y Word) Word { return x << y },
	func(x, y Word) Word { return x >> y },
	func(x, y Word) Word { return x & y },
	func(x, y Word) Word { return x | y },
	func(x, _ Word) Word { return ^x },
	func(x, y Word) Word { return x ^ y },
	nil,
	nil,
	nil,
	nil,
	nil,
	/* No more after this point please :D */
}

type IllegalOperation struct {
	Opcode byte
	Group  string
}

type DivideByZeroError struct{}

func (e *IllegalOperation) Error() string {
	return fmt.Sprintf("Invalid opcode %d in group %s", e.Opcode, e.Group)
}

func (e *DivideByZeroError) Error() string {
	return "Attempted to divide by zero"
}

func (core *Core) performArithmeticOperation(op byte, dest, src0, src1 RegisterIndex) error {
	var s0, s1, o Word
	var fn SimpleArithmeticOperation
	o = Word(op)
	s0 = core.registers[src0]
	if o < ImmediateFormOffset {
		s1 = core.registers[src1]
	} else {
		s1 = Word(src1)
		o = o % ImmediateFormOffset
	}
	/* Check to make sure we aren't doing anything dumb */
	if (o == ArithmeticOperationDivide || o == ArithmeticOperationRemainder) && s1 == 0 {
		return &DivideByZeroError{}
	} else {
		fn = simpleArithmeticOperations[o]
		if fn == nil {
			return &IllegalOperation{op, "Arithmetic"}
		} else {
			core.registers[dest] = fn(s0, s1)
			return nil
		}
	}
}
