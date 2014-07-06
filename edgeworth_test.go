package edgeworth

import "testing"

func Test_GetControlByte_1(t *testing.T) {
	var control byte
	var input FrontendInstruction
	control = 0x12
	input = 0x00000012

	if input.GetControlByte() != control {
		t.Errorf("Failed decoding the control byte of value: %d, Expected: %d, Got: %d!", input, control, input.GetControlByte())
	}
}

func Test_GetControlByte_2(t *testing.T) {
	var control byte
	var input FrontendInstruction
	control = 0x12
	input = 0x34567812

	if input.GetControlByte() != control {
		t.Errorf("Failed decoding the control byte of value: %d, Expected: %d, Got: %d!", input, control, input.GetControlByte())
	}
}

func Test_GetRegisterIndex_1(t *testing.T) {
	var control = []byte{0x6, 0x5, 0x4, 0x3, 0x2, 0x1}
	var input FrontendInstruction
	input = 0x12345600
	for i := 0; i < 6; i++ {
		tmp, err := input.GetRegisterIndex(i)
		if err == nil {
			if tmp != control[i] {
				t.Errorf("Failed decoding the register index at %d. Expected: %d, Got: %d!", i, control[i], tmp)
			} else {
				t.Logf("Index: %d, Expected: %d, Got: %d, Result: Success", i, control[i], tmp)
			}
		} else {
			t.Error(err)
		}
	}
}

func Test_GetRegisterIndex_2(t *testing.T) {
	var input FrontendInstruction
	input = 0
	_, err := input.GetRegisterIndex(MaxRegistersPerInstruction)
	if err == nil {
		t.Errorf("Failed to generate error for accessing out of range register index index: %d", MaxRegistersPerInstruction)
	}
}
