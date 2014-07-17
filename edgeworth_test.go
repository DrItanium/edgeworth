package edgeworth

import "testing"

var DefaultInstructionEncoding Instruction = 0x01234123

func Test_GetGroup_1(t *testing.T) {
	control := GroupBits(0x3)
	bits := DefaultInstructionEncoding.GetControlBits()
	if bits.GetGroup() != control {
		t.Errorf("Expected %d, got %d.", control, bits.GetGroup())
	}
}

func Test_GetOperation_1(t *testing.T) {
	control := OpBits(0x4)
	bits := DefaultInstructionEncoding.GetControlBits()
	if bits.GetOperation() != control {
		t.Errorf("Expected %d, got %d.", control, bits.GetOperation())
	}
}

func Test_GetControlBits_1(t *testing.T) {
	control := ControlBits(0x23)
	bits := DefaultInstructionEncoding.GetControlBits()
	if bits != control {
		t.Errorf("Expected %d, got %d.", control, bits)
	}
}
