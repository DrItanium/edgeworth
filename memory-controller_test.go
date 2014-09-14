package edgeworth

import "testing"

func memoryControllerTestBase(t *testing.T, value MemoryAddress) {
	a := MemoryAddress(value)
	control := MemoryAddress(value >> 3)
	control2 := MemoryAddress(value & 0x7)
	addr, offset := a.MemoryControllerAddress()
	if addr != control || offset != control2 {
		t.Errorf("Expected: %d, %d, Got: %d, %d", control, control2, addr, offset)
	} else {
		t.Logf("Got: %d, %d from %d", addr, offset, value)
	}
}
func TestGetMemoryControllerAddress_1(t *testing.T) {
	memoryControllerTestBase(t, 0x2)
}

func TestGetMemoryControllerAddress_2(t *testing.T) {
	memoryControllerTestBase(t, 0xABC8)
}

func TestGetMemoryControllerAddress_3(t *testing.T) {
	memoryControllerTestBase(t, 0xFDED)
}

func TestMemoryController_1(t *testing.T) {
	controlValue := MemoryWord(0xFDED)
	controlAddress := MemoryAddress(0x0)
	var memory Memory
	memory.StoreWord(controlAddress, controlValue)
	// now that we have set it up lets now access that value
	result := memory.LoadWord(controlAddress)
	if result != controlValue {
		t.Errorf("Expected value: %d, Got: %d from address: %d", controlAddress, result, controlAddress)
	}
}

func TestMemoryController_2(t *testing.T) {
	controlValue := MemoryWord(0xABCDEF0123456789)
	controlAddress := MemoryAddress(0x1)
	var memory Memory
	memory.StoreWord(controlAddress, controlValue)
	// now that we have set it up lets now access that value
	result := memory.LoadWord(controlAddress)
	if result != controlValue {
		t.Errorf("Expected value: %d, Got: %d from address: %d", controlAddress, result, controlAddress)
	}
	// check the actual memory contents
	addr, _ := controlAddress.MemoryControllerAddress()
	result0 := memory.LoadWord(addr << 3)
	result1 := memory.LoadWord((addr + 1) << 3)
	t.Logf("lower address contents: %x, upper address contents: %x", result0, result1)
}
