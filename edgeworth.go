/*
   Package edgeworth is an accumulator based architecture used to explore using
   channels in go. It is a 32-bit word/data architecture
*/
package edgeworth

import "fmt"

/* indirect registers */

const (
	MemoryCapacity        = 4294967296 / 4 //Internally the Memory unit loads 32-bit addresses
	InstructionBundleSize = 4              //We load four instructions at a time
)

type Word uint32
type Control byte
type SystemMemory [MemoryCapacity]Word
type UnitChannel chan Instruction

type Instruction struct {
	Control Control
	Value   Word
}

type AccumulatorUnit struct {
	Input UnitChannel
	Value Word
}

type MemoryUnit struct {
	AccumulatorUnit
	Memory *SystemMemory
}

type Core struct {
	Accumulator UnitChannel
	Memory      UnitChannel
}

func (core *Core) InitializeCore(accumulator, memory, UnitChannel) {
	/* initialize all of the different pieces of the core */
	core.Accumulator = accumulator
	core.Memory = memory
}
