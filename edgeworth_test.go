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

func Test_SetControlBits_1(t *testing.T) {
	var control Instruction
	var result Instruction
	var bits ControlBits
	bits = 0x1234
	result = 0
	control = 0x0000000000001234
	result.SetControlBits(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_SetDestination0_1(t *testing.T) {
	var control Instruction
	var result Instruction
	var bits RegisterIndex
	bits = 0x34
	result = 0
	control = 0x0000000000340000
	result.SetDestination0(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}
func Test_SetDestination1_1(t *testing.T) {
	var control Instruction
	var result Instruction
	var bits RegisterIndex
	bits = 0x12
	result = 0
	control = 0x0000000012000000
	result.SetDestination1(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}
func Test_SetSource0_1(t *testing.T) {
	var control Instruction
	var result Instruction
	var bits RegisterIndex
	bits = 0x34
	result = 0
	control = 0x0000003400000000
	result.SetSource0(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}
func Test_SetSource1_1(t *testing.T) {
	var control Instruction
	var result Instruction
	var bits RegisterIndex
	bits = 0x12
	result = 0
	control = 0x0000120000000000
	result.SetSource1(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_SetSource2_1(t *testing.T) {
	var control Instruction
	var result Instruction
	var bits RegisterIndex
	bits = 0x34
	result = 0
	control = 0x0034000000000000
	result.SetSource2(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}
func Test_SetSource3_1(t *testing.T) {
	var control Instruction
	var result Instruction
	var bits RegisterIndex
	bits = 0x12
	result = 0
	control = 0x1200000000000000
	result.SetSource3(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_SetImmediate_1(t *testing.T) {
	var control Instruction
	var result Instruction
	var bits Word
	bits = 0x12341234
	result = 0
	control = 0x1234123400000000
	result.SetImmediate(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_GetGroupBits_1(t *testing.T) {
	var control GroupBits
	var input ControlBits
	var result GroupBits
	control = 0x4
	input = 0x1234
	result = input.GetGroupBits()
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_GetOpBits_1(t *testing.T) {
	var control OpBits
	var input ControlBits
	var result OpBits
	control = 0x0123
	input = 0x1234
	result = input.GetOpBits()
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_SetGroupBits_1(t *testing.T) {
	var control ControlBits
	var result ControlBits
	var bits GroupBits
	bits = 0x4
	result = 0
	control = 0x0004
	result.SetGroupBits(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_SetGroupBits_2(t *testing.T) {
	// test to make sure that we aren't reading the upper four bits of the byte
	var control ControlBits
	var result ControlBits
	var bits GroupBits
	bits = 0xF4
	result = 0
	control = 0x0004
	result.SetGroupBits(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_SetOpBits_1(t *testing.T) {
	var control ControlBits
	var result ControlBits
	var bits OpBits
	bits = 0x0123
	result = 0
	control = 0x1230
	result.SetOpBits(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}

func Test_SetOpBits_2(t *testing.T) {
	// test to make sure that we aren't reading the upper four bits of the OpBits
	var control ControlBits
	var result ControlBits
	var bits OpBits
	bits = 0xF123
	result = 0
	control = 0x1230
	result.SetOpBits(bits)
	if result != control {
		t.Errorf("Expected %d, got %d", control, result)
	}
}
