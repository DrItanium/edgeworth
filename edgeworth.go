// An implementation of iris in go
package edgeworth

import "fmt"

type Bit bool
type RegisterIndex byte
type GroupBits byte
type OpBits byte
type ControlBits byte
type Instruction uint32
type Word uint16
type DecodedOperation func()

const (
	RegisterCount        = 256
	MemorySize           = 65536
	InstructionGroupMax  = 8
	OperationPerGroupMax = 32
	// Registers implicitly used by the microarchitecture
	InstructionPointerIndex = 255
	StackPointerIndex       = 254
	ZeroRegisterIndex       = 253
	Temporary0Register      = 252
	Temporary1Register      = 251
	Temporary2Register      = 250
	Temporary3Register      = 249
	BankSelectRegister      = 248
)

const (
	ArithmeticGroup = iota
	ConditionalGroup
	BranchGroup
	MemoryOperationsGroup
	MiscGroup
)

type Core struct {
	Registers                    [RegisterCount]Word
	CodeMemory                   [MemorySize]Instruction
	DataMemory                   [MemorySize]Word
	StackMemory                  [MemorySize]Word
	AdvanceIP                    Bit
	TerminateExecution           Bit
	ExecutionStream              chan DecodedOperation
	TemporaryInstructionRegister Instruction
}

func (c *Core) InitializeCore() {
	c.ExecutionStream = make(chan DecodedOperation, 16)
	c.Registers[ZeroRegisterIndex] = 0
}

type InstructionRegisterRequestError struct {
	Write bool
	Index int
}

type InvalidOperationError struct {
	Cbits ControlBits
}

type DivideByZeroError struct{}

func (e *DivideByZeroError) Error() string {
	return "Attempted to divide by zero"
}

func (e *InvalidOperationError) Error() string {
	return fmt.Sprintf("Invalid operation: %s", e.Cbits)
}

func (e *InstructionRegisterRequestError) Error() string {
	if e.Write {
		return fmt.Sprintf("Attempted to write to register field %d which is out of range", e.Index)
	} else {
		return fmt.Sprintf("Attempted to request register field %d which is out of range", e.Index)
	}
}

func (c *Core) GetInstructionPointer() Word {
	return c.Registers[InstructionPointerIndex]
}

func (c *Core) GetStackPointer() Word {
	return c.Registers[StackPointerIndex]
}

func (c *Core) SetStackPointer(value Word) {
	c.Registers[StackPointerIndex] = value
}

func (c *Core) SetInstructionPointer(value Word) {
	c.Registers[InstructionPointerIndex] = value
}

func (i *Instruction) GetControlBits() ControlBits {
	return ControlBits(*i & 0x000000FF)
}

func (i *Instruction) GetRegisterField(index int) (RegisterIndex, error) {
	switch index {
	case 0:
		return RegisterIndex(*i & 0x0000FF00 >> 8), nil
	case 1:
		return RegisterIndex(*i & 0x00FF0000 >> 16), nil
	case 2:
		return RegisterIndex(*i & 0xFF000000 >> 24), nil
	default:
		return 0, &InstructionRegisterRequestError{Index: index, Write: false}
	}
}

func (i *ControlBits) GetGroup() GroupBits {
	return GroupBits(*i & 0x7)
}

func (i *ControlBits) GetOperation() OpBits {
	return OpBits(*i & 0xF8 >> 3)
}

func (i *Instruction) GetImmediate() Word {
	return Word(*i & 0xFFFF0000 >> 16)
}

func (i *Instruction) SetImmediate(value Word) {
	*i = *i&^0xFFFF0000 | Instruction(value)<<16
}

func (i *Instruction) SetRegisterField(index int, value RegisterIndex) error {
	switch index {
	case 0:
		*i = *i&^0x0000FF00 | Instruction(value)<<8
		return nil
	case 1:
		*i = *i&^0x00FF0000 | Instruction(value)<<16
		return nil
	case 2:
		*i = *i&^0xFF000000 | Instruction(value)<<24
		return nil
	default:
		return &InstructionRegisterRequestError{Index: index, Write: true}
	}
}

func (i *Instruction) ExtractRegisterFields() (RegisterIndex, RegisterIndex, RegisterIndex) {
	a, _ := i.GetRegisterField(0)
	b, _ := i.GetRegisterField(1)
	c, _ := i.GetRegisterField(2)
	return a, b, c
}

func (i *Instruction) ExtractRegisterImmediateFields() (RegisterIndex, Word) {
	a, _ := i.GetRegisterField(0)
	return a, i.GetImmediate()
}
func (c ControlBits) String() string {
	return fmt.Sprintf("(%s, %d)", c.GetGroup(), c.GetOperation())
}

func (g GroupBits) String() string {
	switch g {
	case ArithmeticGroup:
		return "arithmetic"
	case ConditionalGroup:
		return "conditional"
	case BranchGroup:
		return "branch"
	case MemoryOperationsGroup:
		return "memory-ops"
	case MiscGroup:
		return "misc"
	default:
		return "undefined"
	}
}
