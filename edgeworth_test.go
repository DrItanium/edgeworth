package edgeworth

import "testing"

func Test_GetGroup_1(t *testing.T) {
	var control GroupBits
	var bits ControlBits
	var input Instruction
	control = 0x3
	input = 0x01234123
	bits = input.GetControlBits()

	if bits.GetGroup() != control {
		t.Errorf("Expected %d, got %d.", control, bits.GetGroup())
	}
}

func Test_GetOperation_1(t *testing.T) {
	var control OpBits
	var bits ControlBits
	var input Instruction
	control = 0x15 // 0xA8 >> 3
	input = 0x012341A8
	bits = input.GetControlBits()
	if bits.GetOperation() != control {
		t.Errorf("Expected %d, got %d.", control, bits.GetOperation())
	}
}
