package edgeworth

import (
	"testing"
)

func Test_InitializeCore_1(t *testing.T) {
	var c Core
	c.InitializeCore()
	if c.GetStackPointer() != EndOfStackSpace {
		t.Error("InitalizeCore did not set the stack pointer correctly")
	}
	for i := 0; i < InstructionCount; i++ {
		if c.code[i] != 0 {
			t.Error("InitializeCore did not clear out code memory correctly")
		}
	}

	for i := 0; i < DataCount; i++ {
		if c.data[i] != 0 {
			t.Error("InitializeCore did not clear out data memory correctly")
		}
	}

	for i := 0; i < StackCount; i++ {
		if c.stack[i] != 0 {
			t.Error("InitializeCore did not clear out stack memory correctly")
		}
	}

	t.Log("InitializeCore working successfully")
}

func Test_GetControlBits_1(t *testing.T) {
	/* we want to see if the control bits are extracted correctly */
	var control ControlBits
	var value Instruction
	control = 0xFFFF
	value = 0x000000000000FFFF
	if value.GetControlBits() != control {
		t.Error("Instruction's GetControlBits did not decode the control bits correctly")
	}
}

func Test_GetControlBits_2(t *testing.T) {
	/* we want to see if the control bits are extracted correctly */
	var control ControlBits
	var value Instruction
	control = 0xABCD
	value = 0x000000000000ABCD
	if value.GetControlBits() != control {
		t.Error("Instruction's GetControlBits did not decode the control bits correctly")
	}
}

func Test_GetRegisterIndex_1(t *testing.T) {
	var control = []RegisterIndex{
		0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45,
	}
	var value Instruction
	value = 0x452301EFCDABFFFF
	for i := 0; i < 6; i++ {
		result, err := value.GetRegisterIndex(i)
		if err != nil {
			t.Errorf("An error was thrown attempting to retrieve register index: %d!", i)
		} else {
			if result != control[i] {
				t.Errorf("Invalid decoding of register index! Field: %d, Expected: %d, Got: %d", i, control[i], result)
			}
		}
	}

}
