// a 32-bit architecture with peculiarities which make it a perfect target as a microcode language
// It is based off of the older 16-bit iris architecture
package edgeworth

type Bit bool
type RegisterIndex byte
type ControlBits uint16
type GroupBits byte // four bits logically
type OpBits uint16  // twelve bits logically
type Instruction uint64

/*
taken from phoenix
{
	control uint16
	destination0 byte
	destination1 byte
	source0 byte
	source1 byte
	source2 byte
	source3 byte
}

and

{
	control uint16
	destination0 byte
	destination1 byte
	immediate Word
}

*/
type Word uint32

const (
	RegisterCount       = 256
	MemorySize          = 16777216
	InstructionGroupMax = 8
)

type Core struct {
	Gpr                [RegisterCount]Word
	CodeMemory         [MemorySize]Instruction
	DataMemory         [MemorySize]Word
	StackMemory        [MemorySize]Word
	InstructionPointer Word
	AdvanceIP          Bit
	TerminateExecution Bit
}

func (i *Instruction) GetControlBits() ControlBits {
	return ControlBits(*i & 0x000000000000FFFF)
}

func (i *Instruction) GetDestination0() RegisterIndex {
	return RegisterIndex(*i & 0x0000000000FF0000 >> 16)
}
func (i *Instruction) GetDestination1() RegisterIndex {
	return RegisterIndex(*i & 0x00000000FF000000 >> 24)
}

func (i *Instruction) GetImmediate() Word {
	return Word(*i & 0xFFFFFFFF00000000 >> 32)
}

func (i *Instruction) GetSource0() RegisterIndex {
	return RegisterIndex(*i & 0x000000FF00000000 >> 32)
}

func (i *Instruction) GetSource1() RegisterIndex {
	return RegisterIndex(*i & 0x0000FF0000000000 >> 40)
}

func (i *Instruction) GetSource2() RegisterIndex {
	return RegisterIndex(*i & 0x00FF000000000000 >> 48)
}

func (i *Instruction) GetSource3() RegisterIndex {
	return RegisterIndex(*i & 0xFF00000000000000 >> 56)
}

func (i *Instruction) SetControlBits(bits ControlBits) {
	*i = *i&^0x000000000000FFFF | Instruction(bits)
}

func (i *Instruction) SetDestination0(rd RegisterIndex) {
	*i = *i&^0x0000000000FF0000 | Instruction(rd)<<16
}

func (i *Instruction) SetDestination1(rd RegisterIndex) {
	*i = *i&^0x00000000FF000000 | Instruction(rd)<<24
}
func (i *Instruction) SetSource0(rs RegisterIndex) {
	*i = *i&^0x000000FF00000000 | Instruction(rs)<<32
}
func (i *Instruction) SetSource1(rs RegisterIndex) {
	*i = *i&^0x0000FF0000000000 | Instruction(rs)<<40
}
func (i *Instruction) SetSource2(rs RegisterIndex) {
	*i = *i&^0x00FF000000000000 | Instruction(rs)<<48
}

func (i *Instruction) SetSource3(rs RegisterIndex) {
	*i = *i&^0xFF00000000000000 | Instruction(rs)<<56
}

func (i *Instruction) SetImmediate(imm Word) {
	*i = *i&^0xFFFFFFFF00000000 | Instruction(imm)<<32
}

func (i *ControlBits) GetGroupBits() GroupBits {
	// we can have a maximum of 16 different groups
	return GroupBits(*i & 0x000F)
}

func (i *ControlBits) GetOpBits() OpBits {
	// we can have a maximum of 4096 operations in a group
	return OpBits(*i & 0xFFF0 >> 4)
}

func (i *ControlBits) SetOpBits(a OpBits) {
	// clear out the upper 4 bits of the opbits
	fix := a & 0x0FFF
	*i = *i&^0xFFF0 | ControlBits(fix<<4)
}

func (i *ControlBits) SetGroupBits(a GroupBits) {
	// clear out the upper 4 bits of the group bits
	fix := a & 0x0F
	*i = *i&^0x000F | ControlBits(fix)
}
