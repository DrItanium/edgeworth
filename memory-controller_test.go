package edgeworth

import "testing"

func TestBlockAddress_1(t *testing.T) {
	a := MemoryAddress(0x2)
	if BlockAddress(a) != 0x0 {
		t.Error("Transform address failed, did not get a zero address!")
	}
}

func TestBlockAddress_2(t *testing.T) {
	a := MemoryAddress(0xABCDEF0100000A01)
	control := MemoryAddress(0xA00)
	if BlockAddress(a) != control {
		t.Errorf("Expected %d, got %d", control, a)
	}
}

func TestStoreBlock_1(t *testing.T) {
	var m Memory
	a := MemoryAddress(0x2) //unaligned address, will be fixed automatically
	b := MemoryWord(0x1)
	m.StoreBlock(a, b)
	if m[0] != 0x1 {
		t.Error("StoreBlock failed to write 0x1 to 0x0 (provided address 0x2)")
	}
}
