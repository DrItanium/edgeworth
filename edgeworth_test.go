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

func Test_GetImmediate_1(t *testing.T) {
	control := Word(0x0123)
	bits := DefaultInstructionEncoding.GetImmediate()
	if bits != control {
		t.Errorf("Expected %d, got %d.", control, bits)
	}
}

func Test_SetRegisterField_1(t *testing.T) {
	var inst Instruction
	inst = 0x23
	err := inst.SetRegisterField(0, 0x41)
	if err != nil {
		t.Error(err)
	}
	err = inst.SetRegisterField(1, 0x23)
	if err != nil {
		t.Error(err)
	}
	err = inst.SetRegisterField(2, 0x01)
	if err != nil {
		t.Error(err)
	}
	if inst != DefaultInstructionEncoding {
		t.Errorf("Expected %d, got %d.", DefaultInstructionEncoding, inst)
	}
}

func Test_SetRegisterField_2(t *testing.T) {
	var inst Instruction
	inst = 0x23
	err := inst.SetRegisterField(4, 0x41)
	if err == nil {
		t.Error("Attempt to set register field 4 succeded!")
	}
}
