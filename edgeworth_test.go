package edgeworth

import "testing"

func Test_GetControlBits_1(t *testing.T) {
	var control ControlBits
	var input Instruction
	control = 0x1234
	input = 0x1234123412341234

	if input.GetControlBits() != control {
		t.Errorf("Expected %d, got %d.", control, input.GetControlBits())
	}
}

func Test_GetDestination0_1(t *testing.T) {
	var control RegisterIndex
	var input Instruction
	control = 0x34
	input = 0x1234123412341234
	value := input.GetDestination0()

	if value != control {
		t.Errorf("Expected %d, got %d.", control, value)
	}
}

func Test_GetDestination1_1(t *testing.T) {
	var control RegisterIndex
	var input Instruction
	control = 0x12
	input = 0x1234123412341234
	value := input.GetDestination1()

	if value != control {
		t.Errorf("Expected %d, got %d.", control, value)
	}
}

func Test_GetSource0_1(t *testing.T) {
	var control RegisterIndex
	var input Instruction
	control = 0x34
	input = 0x1234123412341234
	value := input.GetSource0()

	if value != control {
		t.Errorf("Expected %d, got %d.", control, value)
	}
}

func Test_GetSource1_1(t *testing.T) {
	var control RegisterIndex
	var input Instruction
	control = 0x12
	input = 0x1234123412341234
	value := input.GetSource1()

	if value != control {
		t.Errorf("Expected %d, got %d.", control, value)
	}
}

func Test_GetSource2_1(t *testing.T) {
	var control RegisterIndex
	var input Instruction
	control = 0x34
	input = 0x1234123412341234
	value := input.GetSource2()

	if value != control {
		t.Errorf("Expected %d, got %d.", control, value)
	}
}

func Test_GetSource3_1(t *testing.T) {
	var control RegisterIndex
	var input Instruction
	control = 0x12
	input = 0x1234123412341234
	value := input.GetSource3()

	if value != control {
		t.Errorf("Expected %d, got %d.", control, value)
	}
}

func Test_GetImmediate_1(t *testing.T) {
	var control Word
	var input Instruction
	control = 0x12341234
	input = 0x1234123412341234
	value := input.GetImmediate()

	if value != control {
		t.Errorf("Expected %d, got %d.", control, value)
	}
}
