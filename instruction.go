// Instruction related operations (encoding, decoding, typedefs) actual binding is done elsewhere.
// Each edgeworth atom is 32 bits in length and each particle is 256 bits in length (yielding 8 atoms)
// atoms can be fused into up to 128 bit molecules if needed (the set operation for instance)
package edgeworth

type Instruction uint64

// Groups
const (
	GroupArithmetic = iota
	GroupControlFlow
	GroupMemory // memory manipulation instructions
	GroupComparison
	GroupSpecial
)
