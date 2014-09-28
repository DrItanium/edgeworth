// the edgeworth system CPU - a simple 64-bit RISC cpu with a large compliment of registers, most of them purpose built
package edgeworth

import (
	"fmt"
	"math"
)

const (
	RegisterCount = 256
)

const (
	RegisterTypeInteger = iota
	RegisterTypeFloat
	RegisterTypePackedFloat32
	RegisterTypePackedInt32
)

type PackedFloat32 struct {
	Lower float32
	Upper float32
}

type PackedInt32 struct {
	Lower uint32
	Upper uint32
}

type InvalidRegisterTypeError struct {
	Type string
}

func (e InvalidRegisterTypeError) Error() string {
	return fmt.Sprintf("Error: invalid register type: %s provided for integer register", e.Type)
}

type Register struct {
	Bits uint64 // raw bits
	Tag  byte   // tag bits
}

func (r *Register) SetValue(tag byte, bits uint64) {
	r.Tag = tag
	r.Bits = bits
}
func (r *Register) SetIntegerValue(value uint64) {
	r.SetValue(RegisterTypeInteger, value)
}

func (r *Register) SetFloatValue(value float64) {
	r.SetValue(RegisterTypeFloat, math.Float64bits(value))
}

func (r *Register) GetFloatValue() float64 {
	return math.Float64frombits(r.Bits)
}

func (r *Register) GetValue() interface{} {
	switch r.Tag {
	case RegisterTypeFloat:
		return r.GetFloatValue()
	case RegisterTypePackedFloat32:
		return r.GetFloat32Values()
	case RegisterTypePackedInt32:
		return r.GetInt32Values()
	case RegisterTypeInteger:
		fallthrough
	default:
		return r.Bits
	}
}

func (r *Register) GetFloat32Values() *PackedFloat32 {
	//get the raw bits for the upper and lower half
	lower := uint32(r.Bits)
	upper := uint32((r.Bits & 0xFFFFFFFF00000000) >> 32)
	return &PackedFloat32{Lower: math.Float32frombits(lower), Upper: math.Float32frombits(upper)}
}

func (r *Register) SetPackedFloat32Value(p *PackedFloat32) {
	blower := uint64(math.Float32bits(p.Lower))
	bupper := uint64(math.Float32bits(p.Upper))
	value := ((bupper << 32) | (blower & 0x00000000FFFFFFFF))
	r.SetValue(RegisterTypePackedFloat32, value)
}
func (r *Register) SetFromFloat32Values(lower, upper float32) {
	var p PackedFloat32
	p.Lower = lower
	p.Upper = upper
	r.SetPackedFloat32Value(&p)
}
func (r *Register) GetInt32Values() *PackedInt32 {
	//get the raw bits for the upper and lower half
	lower := uint32(r.Bits)
	upper := uint32(r.Bits >> 32)
	return &PackedInt32{Lower: lower, Upper: upper}
}

func (r *Register) SetPackedInt32Value(p *PackedInt32) {
	blower := uint64(p.Lower)
	bupper := uint64(p.Upper)
	value := ((bupper << 32) | (blower & 0x00000000FFFFFFFF))
	r.SetValue(RegisterTypePackedInt32, value)
}
func (r *Register) SetFromInt32Values(lower, upper uint32) {
	var p PackedInt32
	p.Lower = lower
	p.Upper = upper
	r.SetPackedInt32Value(&p)
}
