// memory controller that the cpu interfaces with.
package edgeworth

type MemoryAddress uint64
type MemoryWord uint64
type Memory [MemoryControllerMax]MemoryWord

const (
	MemoryWordByteCount    = 8
	MemoryTwoWordByteCount = MemoryWordByteCount * 2
	MemoryControllerMax    = AcceptableMemoryRange >> 3 // eliminate the lower three bits
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

func byteArrayToWord(arr []byte, offset int) MemoryWord {
	return (MemoryWord(arr[offset+7]) << 56) | (MemoryWord(arr[offset+6]) << 48) |
		(MemoryWord(arr[offset+5]) << 40) | (MemoryWord(arr[offset+4]) << 32) |
		(MemoryWord(arr[offset+3]) << 24) | (MemoryWord(arr[offset+2]) << 16) |
		(MemoryWord(arr[offset+1]) << 8) | (MemoryWord(arr[offset]))
}

func byteArrayToDualWords(arr []byte) (MemoryWord, MemoryWord) {
	return byteArrayToWord(arr, 0), byteArrayToWord(arr, MemoryWordByteCount)
}
func getByteArrayFromWord(addr MemoryWord) []byte {
	return []byte{
		byte(addr),
		byte(addr >> 8),
		byte(addr >> 16),
		byte(addr >> 24),
		byte(addr >> 32),
		byte(addr >> 40),
		byte(addr >> 48),
		byte(addr >> 56),
	}

}
func wordsToByteArray(addr0, addr1 MemoryWord) []byte {
	return []byte{
		byte(addr0),
		byte(addr0 >> 8),
		byte(addr0 >> 16),
		byte(addr0 >> 24),
		byte(addr0 >> 32),
		byte(addr0 >> 40),
		byte(addr0 >> 48),
		byte(addr0 >> 56),
		byte(addr1),
		byte(addr1 >> 8),
		byte(addr1 >> 16),
		byte(addr1 >> 24),
		byte(addr1 >> 32),
		byte(addr1 >> 40),
		byte(addr1 >> 48),
		byte(addr1 >> 56),
	}
}
func (address MemoryAddress) MemoryControllerAddress() (MemoryAddress, MemoryAddress) {
	//we will decompose the function into the upper 61 bits and the lower three bits
	//first thing is to get rid of the upper half of the address and then construct an actuall
	//address which we can use
	return address & (AcceptableMemoryRange - 7) >> 3, address & 0x7
}

func (memory *Memory) LoadWord(address MemoryAddress) MemoryWord {
	addr, offset := address.MemoryControllerAddress()
	value0 := memory[addr]
	if offset == 0 {
		return value0
	} else {
		value1 := memory[addr+1]
		container := wordsToByteArray(value0, value1)
		return (MemoryWord(container[offset+7]) << 56) | (MemoryWord(container[offset+6]) << 48) |
			(MemoryWord(container[offset+5]) << 40) | (MemoryWord(container[offset+4]) << 32) |
			(MemoryWord(container[offset+3]) << 24) | (MemoryWord(container[offset+2]) << 16) |
			(MemoryWord(container[offset+1]) << 8) | (MemoryWord(container[offset]))
	}
}

func (memory *Memory) StoreWord(address MemoryAddress, value MemoryWord) {
	addr, offset := address.MemoryControllerAddress()
	if offset == 0 {
		memory[addr] = value
	} else {
		value0 := memory[addr]
		value1 := memory[addr+1]
		container := wordsToByteArray(value0, value1)
		bytes := getByteArrayFromWord(value)
		// map into the space we pulled out
		for i := 0; i < MemoryWordByteCount; i++ {
			container[offset+MemoryAddress(i)] = bytes[i]
		}
		// reconstruct the words
		first, second := byteArrayToDualWords(container)
		// write to memory
		memory[addr] = first
		memory[addr+1] = second
	}
}
