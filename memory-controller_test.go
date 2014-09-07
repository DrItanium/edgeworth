package edgeworth

import "testing"

func TestCheckTransformAddress_1(t *testing.T) {
	a := MemoryAddress(0x2)
	if BlockAddress(a) != 0x0 {
		t.Error("Transform address failed, did not get a zero address!")
	}
}

func TestCheckTransformAddress_2(t *testing.T) {
	a := MemoryAddress(0xABCDEF0100000A01)
	control := MemoryAddress(0xA00)
	if BlockAddress(a) != control {
		t.Errorf("Expected %d, got %d", control, a)
	}
}
