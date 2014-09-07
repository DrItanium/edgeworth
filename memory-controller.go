// memory controller that the cpu interfaces with.
package edgeworth

type MemoryAddress uint64
type MemoryWord uint64
type Memory [AcceptableMemoryRange]byte

const (
	MemoryWordByteCount = 8
)

var DecomposeMasks = []MemoryWord{
	0x00000000000000FF,
	0x000000000000FF00,
	0x0000000000FF0000,
	0x00000000FF000000,
	0x000000FF00000000,
	0x0000FF0000000000,
	0x00FF000000000000,
	0xFF00000000000000,
}

// the controller loads 64-bit values only, so we have to add some extra
func BlockAddress(address MemoryAddress) MemoryAddress {
	return address & (AcceptableMemoryRange - 7)
}

func (m *Memory) LoadBlock(address MemoryAddress) MemoryWord {
	var container [8]byte
	transformedAddr := BlockAddress(address)
	for i := 0; i < 8; i++ {
		offset := transformedAddr + MemoryAddress(i)
		container[i] = m[offset]
	}
	// little endian
	return (MemoryWord(container[7]) << 56) | (MemoryWord(container[6]) << 48) |
		(MemoryWord(container[5]) << 40) | (MemoryWord(container[4]) << 32) |
		(MemoryWord(container[3]) << 24) | (MemoryWord(container[2]) << 16) |
		(MemoryWord(container[1]) << 8) | (MemoryWord(container[0]))
}

func (m *Memory) StoreBlock(address MemoryAddress, value MemoryWord) {
	transformedAddr := BlockAddress(address)
	for i := 0; i < MemoryWordByteCount; i++ {
		offset := transformedAddr + MemoryAddress(i)
		m[offset] = byte(value & DecomposeMasks[i])
	}
}
