/*
   Package edgeworth is a second generation iris core.

   It provides instruction decoding, execution, unparsing, and mnemonic
   generation.

   Edgeworth extends the iris architecture to 32-bits and is a modified harvard
   architecture where data and code can be used somewhat interchangeably
*/
package edgeworth

import "fmt"

const (
	RegisterCount    = 256
	InstructionCount = 16777216
	DataCount        = 16777216
)

type Instruction uint64
type ControlBits uint16
type RegisterIndex byte
type Datum uint32
type Register uint32

type Core struct {
	registers [RegisterCount]Register
	data      [DataCount]Datum
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
		return 0, &GetRegisterIndexError{
			index,
		}
	}
}

func (d Datum) GetTagBits() byte {
	return byte((d & 0xFF000000) >> 24)
}

func (d Datum) GetValue() Datum {
	return Datum(d & 0x00FFFFFF)
}
