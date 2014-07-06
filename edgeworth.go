package edgeworth

import "fmt"

/* indirect registers */

const (
	MemoryCapacity             = 4294967296 / 4
	FrontendRegisterCount      = 16
	MaxRegistersPerInstruction = 6
)

type Word uint32
type SystemMemory [MemoryCapacity]Word

type FrontendCore struct {
	Registers [FrontendRegisterCount]Word
	Memory    *SystemMemory
}

type FrontendInstruction uint32 // 8 bits control, 8 bits reg0, reg1, 16-bits immediate or more registers

type FrontendError struct {
	Message   string
	ErrorCode int
}

func (fe *FrontendError) Error() string {
	return fmt.Sprintf("Frontend Error %d: %s", fe.ErrorCode, fe.Message)
}
func (f FrontendInstruction) GetControlByte() byte {
	return byte(f & 0x000000FF)
}
func (f FrontendInstruction) GetRegisterIndex(index int) (byte, error) {
	/* clear the control byte */
	switch index {
	case 0:
		return byte(f & 0x00000F00 >> 8), nil
	case 1:
		return byte(f & 0x0000F000 >> 12), nil
	case 2:
		return byte(f & 0x000F0000 >> 16), nil
	case 3:
		return byte(f & 0x00F00000 >> 20), nil
	case 4:
		return byte(f & 0x0F000000 >> 24), nil
	case 5:
		return byte(f & 0xF0000000 >> 28), nil
	default:
		return 0, &FrontendError{fmt.Sprintf("Attempted to access register index at index %d which is out of range", index), 1}
	}
}

func (f FrontendInstruction) GetImmediateValue() uint16 {
	return uint16(f & 0xFFFF0000 >> 16)
}
