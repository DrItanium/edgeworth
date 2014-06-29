/*
   Package edgeworth is a second generation iris core.

   It provides instruction decoding, execution, unparsing, and mnemonic
   generation.

   Edgeworth extends the iris architecture to 32-bits and is a modified harvard
   architecture where data and code can be used somewhat interchangeably.

   It also provides a half instruction format for making it possible to have
   two instructions per wide instruction (most register based instructions do
   not need all of the encoding space so lets put several of them in there)

   The other important thing to note is that bottom five instruction groups are
   an exact duplicate of the iris ones and still operate on 16-bit values
   (except for the registers).
*/
package edgeworth

import "fmt"

const (
	RegisterCount                    = 256
	InstructionCount                 = 16777216
	DataCount                        = 16777216
	StackCount                       = 16777216
	InstructionControlBitSection     = 0x000000000000FFFF
	InstructionRegisterSection       = 0xFFFFFFFFFFFF0000
	InstructionMaxRegisterCount      = 6
	HalfInstructionControlBitSection = 0x000000FF
	HalfInstructionRegisterSection   = 0xFFFFFF00
	HalfInstructionMaxRegisterCount  = 3
	EndOfStackSpace                  = StackCount - 1
	EndOfCodeSpace                   = InstructionCount - 1
	EndOfDataSpace                   = DataCount - 1
	CodeSectionId                    = "code"
	DataSectionId                    = "data"
	StackSectionId                   = "stack"
	ProgramCounterId                 = "pc"
	StackPointerId                   = "sp"
)

/* reserved registers */
const (
	ProgramCounterRegisterIndex   = 255
	StackPointerRegisterIndex     = 254
	ImpliedPredicateRegisterIndex = 253
)

/* indirect registers */

type Instruction uint64
type ControlBits uint16

type HalfInstruction uint32
type HalfControlBits byte

type RegisterIndex byte
type Word uint32

type GetRegisterIndexError struct {
	IndexProvided int
}

type ProcessorCounterError struct {
	Type     string
	ValueSet Word
}

type MemoryError struct {
	Address Word
	IsLoad  bool
	Section string
}

func (e *GetRegisterIndexError) Error() string {
	return fmt.Sprintf("Register index %d is out of range!", e.IndexProvided)
}

func (e *ProcessorCounterError) Error() string {
	return fmt.Sprintf("Attempted to set %s to %d which is out of range!", e.Type, e.ValueSet)
}

func (e *MemoryError) Error() string {
	if e.IsLoad {
		return fmt.Sprintf("Attempted to load address %d which is out of %s memory range!", e.Address, e.Section)
	} else {
		return fmt.Sprintf("Attempted to write to an address %d which is out of %s memory range!", e.Address, e.Section)
	}
}

func (inst Instruction) GetControlBits() ControlBits {
	return ControlBits(inst & InstructionControlBitSection)
}

func (inst Instruction) GetRegisterIndex(index int) (RegisterIndex, error) {
	var a Instruction
	a = (inst & InstructionRegisterSection) >> 16
	if index >= 0 && index < InstructionMaxRegisterCount {
		return RegisterIndex(a >> (8 * Instruction(index))), nil
	} else {
		return 0, &GetRegisterIndexError{index}
	}
}

func (inst HalfInstruction) GetControlBits() HalfControlBits {
	return HalfControlBits(inst & HalfInstructionControlBitSection)
}

func (inst HalfInstruction) GetRegisterIndex(index int) (RegisterIndex, error) {
	var a HalfInstruction
	a = (inst & HalfInstructionRegisterSection) >> 8
	if index >= 0 && index < HalfInstructionMaxRegisterCount {
		return RegisterIndex(a >> (8 * HalfInstruction(index))), nil
	} else {
		return 0, &GetRegisterIndexError{index}
	}
}

type Core struct {
	registers [RegisterCount]Word
	code      [InstructionCount]Instruction
	data      [DataCount]Word
	stack     [StackCount]Word
}

func (core *Core) InitializeCore() {
	/* initialize all of the different pieces of the core */
	for i := 0; i < RegisterCount; i++ {
		core.registers[i] = 0
	}
	for i := 0; i < InstructionCount; i++ {
		core.code[i] = 0
	}
	for i := 0; i < DataCount; i++ {
		core.data[i] = 0
	}
	for i := 0; i < StackCount; i++ {
		core.stack[i] = 0
	}
	/* now set the corresponding registers correctly */
	core.SetStackPointer(EndOfStackSpace)
}

func (core *Core) IncrementProgramCounter() {
	pc := core.GetProgramCounter()
	if pc == EndOfCodeSpace {
		pc = 0
	} else {
		pc = pc + 1
	}
	core.SetProgramCounter(pc)
}

func (core *Core) DecrementProgramCounter() {
	pc := core.GetProgramCounter()
	if pc == 0 {
		pc = EndOfCodeSpace
	} else {
		pc = pc - 1
	}
	core.SetProgramCounter(pc)
}
func (core *Core) GetProgramCounter() Word {
	return core.registers[ProgramCounterRegisterIndex]
}
func (core *Core) SetProgramCounter(address Word) error {
	if address >= InstructionCount {
		return &ProcessorCounterError{ProgramCounterId, address}
	} else {
		core.registers[ProgramCounterRegisterIndex] = address
		return nil
	}
}

func (core *Core) LoadFromDataAddress(address Word) (Word, error) {
	if address >= DataCount {
		return 0, &MemoryError{address, true, DataSectionId}
	} else {
		return core.data[address], nil
	}
}

func (core *Core) WriteToDataAddress(address, value Word) error {
	if address >= DataCount {
		return &MemoryError{address, false, DataSectionId}
	} else {
		core.data[address] = value
		return nil
	}
}

func (core *Core) LoadFromCodeAddress(address Word) (Instruction, error) {
	if address >= InstructionCount {
		return 0, &MemoryError{address, true, CodeSectionId}
	} else {
		return core.code[address], nil
	}
}

func (core *Core) WriteToCodeAddress(address Word, value Instruction) error {
	if address >= InstructionCount {
		return &MemoryError{address, false, CodeSectionId}
	} else {
		core.code[address] = value
		return nil
	}
}
func (core *Core) GetStackPointer() Word {
	return core.registers[StackPointerRegisterIndex]
}

func (core *Core) SetStackPointer(value Word) error {
	if value >= StackCount {
		return &ProcessorCounterError{StackPointerId, value}
	} else {
		core.registers[StackPointerRegisterIndex] = value
		return nil
	}
}

func (core *Core) WriteToStackAddress(address, value Word) error {
	if address >= StackCount {
		return &MemoryError{address, false, StackSectionId}
	} else {
		core.stack[address] = value
		return nil
	}
}

func (core *Core) ReadFromStackAddress(address Word) (Word, error) {
	if address >= StackCount {
		return 0, &MemoryError{address, true, StackSectionId}
	} else {
		return core.stack[address], nil
	}
}

func (core *Core) PushOntoStack(value Word) {
	sp := core.GetStackPointer()
	if sp == EndOfStackSpace {
		sp = 0
	} else {
		sp = sp + 1
	}
	core.WriteToStackAddress(sp, value)
	core.SetStackPointer(sp)
}

func (core *Core) PopOffStack() Word {
	sp := core.GetStackPointer()
	value, _ := core.ReadFromStackAddress(sp)
	if sp == 0 {
		sp = EndOfStackSpace
	} else {
		sp = sp - 1
	}
	core.SetStackPointer(sp)
	return value
}
