/*
   Package edgeworth is a second generation iris core.

   It provides instruction decoding, execution, unparsing, and mnemonic
   generation.

   Edgeworth extends the iris architecture to 32-bits and is a modified harvard
   architecture where data and code can be used somewhat interchangeably.

   It also provides a half instruction format for making it possible to have
   two instructions per wide instruction (most register based instructions do
   not need all of the encoding space so lets put several of them in there)
*/
package edgeworth

import "fmt"

const (
	RegisterCount    = 256
	InstructionCount = 16777216
	DataCount        = 16777216
)

type Instruction uint64
type HalfInstruction uint32
type ControlBits uint16
type HalfControlBits byte
type RegisterIndex byte
type Word uint32
type Register uint32

type Core struct {
	registers [RegisterCount]Register
	data      [DataCount]Word
	code      [InstructionCount]Instruction
}
type GetRegisterIndexError struct {
	IndexProvided int
}

func (e *GetRegisterIndexError) Error() string {
	return fmt.Sprintf("Register index %d is out of range!", e.IndexProvided)
}

func (inst Instruction) GetControlBits() ControlBits {
	return ControlBits(inst & 0x000000000000FFFF)
}

func (inst Instruction) GetRegisterIndex(index int) (RegisterIndex, error) {
	var a Instruction
	a = (inst & 0xFFFFFFFFFFFF0000) >> 16
	if index >= 0 && index < 6 {
		return RegisterIndex(a >> (8 * Instruction(index))), nil
	} else {
		return 0, &GetRegisterIndexError{index}
	}
}
func (inst HalfInstruction) GetControlBits() HalfControlBits {
	return HalfControlBits(inst & 0x000000FF)
}

func (inst HalfInstruction) GetRegisterIndex(index int) (RegisterIndex, error) {
	var a HalfInstruction
	a = (inst & 0xFFFFFF00) >> 8
	if index >= 0 && index < 3 {
		return RegisterIndex(a >> (8 * HalfInstruction(index))), nil
	} else {
		return 0, &GetRegisterIndexError{index}
	}
}
