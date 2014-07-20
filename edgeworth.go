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

const (
	RegisterCount           = 256
	MemorySize              = 65536
	InstructionPointerIndex = 255
	StackPointerIndex       = 254
	InstructionGroupMax     = 8
	OperationPerGroupMax    = 32
)

type Core struct {
	Registers          [RegisterCount]Word
	CodeMemory         [MemorySize]Instruction
	DataMemory         [MemorySize]Word
	StackMemory        [MemorySize]Word
	AdvanceIP          Bit
	TerminateExecution Bit
}
type InstructionRegisterRequestError struct {
	Write bool
	Index int
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
