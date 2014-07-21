//arithmetic operations
package edgeworth

type ArithmeticOperation func(Word, Word) Word

const (
	ArithmeticAdd = iota
	ArithmeticSubtract
	ArithmeticMultiply
	ArithmeticDivide
	ArithmeticRemainder
	ArithmeticShiftLeft
	ArithmeticShiftRight
	ArithmeticBitwiseAnd
	ArithmeticBitwiseOr
	ArithmeticBitwiseNot
	ArithmeticBitwiseXor
	ArithmeticBitwiseClear
	ArithmeticAddImmediate
	ArithmeticSubtractImmediate
	ArithmeticMultiplyImmediate
	ArithmeticDivideImmediate
	ArithmeticRemainderImmediate
	ArithmeticShiftLeftImmediate
	ArithmeticShiftRightImmediate
	ArithmeticBitwiseAndFullImmediate
	ArithmeticBitwiseOrFullImmediate
	ArithmeticBitwiseNotFullImmediate
	ArithmeticBitwiseXorFullImmediate
	ArithmeticBitwiseClearFullImmediate
)

var BasicArithmeticOperations = [OperationPerGroupMax]ArithmeticOperation{
	func(x, y Word) Word { return x + y },  // add
	func(x, y Word) Word { return x - y },  // sub
	func(x, y Word) Word { return x * y },  // mul
	func(x, y Word) Word { return x / y },  // div
	func(x, y Word) Word { return x % y },  // remainder
	func(x, y Word) Word { return x << y }, // shift left
	func(x, y Word) Word { return x >> y }, // shift right
	func(x, y Word) Word { return x & y },  // bitwise and
	func(x, y Word) Word { return x | y },  // bitwise or
	func(x, _ Word) Word { return ^x },     // bitwise not
	func(x, y Word) Word { return x ^ y },  // bitwise xor
	func(x, y Word) Word { return x &^ y }, //bitwise clear

}

func (op *OpBits) IsArithmeticRegisterForm() bool {
	switch *op {
	case ArithmeticAdd, ArithmeticSubtract, ArithmeticMultiply, ArithmeticDivide, ArithmeticRemainder, ArithmeticShiftLeft, ArithmeticShiftRight, ArithmeticBitwiseAnd, ArithmeticBitwiseOr, ArithmeticBitwiseNot, ArithmeticBitwiseXor, ArithmeticBitwiseClear:
		return true
	default:
		return false
	}
}

func (op *OpBits) IsArithmeticHalfImmediateForm() bool {
	switch *op {
	case ArithmeticAddImmediate, ArithmeticSubtractImmediate, ArithmeticMultiplyImmediate, ArithmeticDivideImmediate, ArithmeticRemainderImmediate, ArithmeticShiftLeftImmediate, ArithmeticShiftRightImmediate:
		return true
	default:
		return false
	}
}
func (op *OpBits) IsArithmeticFullImmediateForm() bool {
	switch *op {
	case ArithmeticBitwiseAndFullImmediate, ArithmeticBitwiseOrFullImmediate, ArithmeticBitwiseNotFullImmediate, ArithmeticBitwiseXorFullImmediate, ArithmeticBitwiseClearFullImmediate:
		return true
	default:
		return false
	}
}
func (c *Core) DecodeArithmeticOperation(i *Instruction) error {
	cbits := i.GetControlBits()
	op := cbits.GetOperation()
	if op.IsArithmeticRegisterForm() {
		idest, isrc0, isrc1 := i.ExtractRegisterFields()
		c.Registers[idest] = BasicArithmeticOperations[op](c.Registers[isrc0], c.Registers[isrc1])
		return nil
	} else if op.IsArithmeticHalfImmediateForm() {
		idest, isrc0, imm := i.ExtractRegisterFields()
		index := 0
		switch op {
		case ArithmeticAddImmediate:
			index = ArithmeticAdd
		case ArithmeticSubtractImmediate:
			index = ArithmeticSubtract
		case ArithmeticMultiplyImmediate:
			index = ArithmeticMultiply
		case ArithmeticDivideImmediate:
			index = ArithmeticDivide
		case ArithmeticRemainderImmediate:
			index = ArithmeticRemainder
		case ArithmeticShiftLeftImmediate:
			index = ArithmeticShiftLeft
		case ArithmeticShiftRightImmediate:
			index = ArithmeticShiftRight
		default:
			return &InvalidOperationError{Cbits: cbits}
		}
		c.Registers[idest] = BasicArithmeticOperations[index](c.Registers[isrc0], Word(imm))
		return nil
	} else if op.IsArithmeticFullImmediateForm() {
		idest, imm := i.ExtractRegisterImmediateFields()
		index := 0
		switch op {
		case ArithmeticBitwiseAndFullImmediate:
			index = ArithmeticBitwiseAnd
		case ArithmeticBitwiseOrFullImmediate:
			index = ArithmeticBitwiseOr
		case ArithmeticBitwiseXorFullImmediate:
			index = ArithmeticBitwiseXor
		case ArithmeticBitwiseClearFullImmediate:
			index = ArithmeticBitwiseClear
		case ArithmeticBitwiseNotFullImmediate:
			/* special case */
			c.Registers[idest] = BasicArithmeticOperations[ArithmeticBitwiseNot](imm, 0)
			return nil
		default:
			return &InvalidOperationError{Cbits: cbits}
		}
		/* This is a destructive modify */
		c.Registers[idest] = BasicArithmeticOperations[index](c.Registers[idest], imm)
		return nil
	} else {
		return &InvalidOperationError{Cbits: cbits}
	}
}
