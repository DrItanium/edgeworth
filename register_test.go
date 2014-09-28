package edgeworth

import (
	"math"
	"testing"
)

const (
	TestingValue64bit           = 0x0123456789ABCDEF
	TestingValue64bitFloat      = 1000.1234
	TestingValue32bitLower      = 0x01234567
	TestingValue32bitUpper      = 0x89ABCDEF
	TestingValue32bitFloatLower = 112.32
	TestingValue32bitFloatUpper = 6561.92
)

var Float64bitsRaw uint64 = math.Float64bits(TestingValue64bitFloat)
var Float32bitsRawLower uint32 = math.Float32bits(TestingValue32bitFloatLower)
var Float32bitsRawUpper uint32 = math.Float32bits(TestingValue32bitFloatUpper)

func expectationCheckUint64(control, value uint64, addendum string, t *testing.T) {
	if value != control {
		t.Errorf("Expected: %d, got: %d %s", control, value, addendum)
	}
}

func expectationCheckUint32(control, value uint32, addendum string, t *testing.T) {
	if value != control {
		t.Errorf("Expected: %d, got: %d %s", control, value, addendum)
	}
}

func expectationCheckFloat64(control, value float64, addendum string, t *testing.T) {
	if value != control {
		t.Errorf("Expected: %f, got: %f %s", control, value, addendum)
	}
}

func expectationCheckFloat32(control, value float32, addendum string, t *testing.T) {
	if value != control {
		t.Errorf("Expected: %f, got: %f %s", control, value, addendum)
	}
}

func expectationCheckByte(control, value byte, addendum string, t *testing.T) {
	if value != control {
		t.Errorf("Expected: %d, got: %d %s", control, value, addendum)
	}
}

func TestRegisterContents_1(t *testing.T) {
	// uint64 raw
	var r Register
	value := uint64(TestingValue64bit)
	r.SetValue(RegisterTypeInteger, value)
	expectationCheckUint64(TestingValue64bit, r.Bits, "(uint64 raw)", t)
	expectationCheckByte(RegisterTypeInteger, r.Tag, "(tag comparison)", t)
}

func TestRegisterContents_2(t *testing.T) {
	//float64
	var r Register
	value := float64(TestingValue64bitFloat)
	r.SetFloatValue(value)
	//check the bits
	expectationCheckUint64(Float64bitsRaw, r.Bits, "(raw bits)", t)
	expectationCheckByte(RegisterTypeFloat, r.Tag, "(tag comparison)", t)
}

func TestRegisterContents_3(t *testing.T) {
	//float32 pair
	var r Register
	r.SetFromFloat32Values(TestingValue32bitFloatLower, TestingValue32bitFloatUpper)
	//check the bits
	expectationCheckUint32(Float32bitsRawLower, uint32(r.Bits), "(lower half raw bits)", t)
	expectationCheckUint32(Float32bitsRawUpper, uint32(r.Bits>>32), "(upper half raw bits)", t)
	expectationCheckByte(RegisterTypePackedFloat32, r.Tag, "(tag comparison)", t)
}

func TestRegisterContents_4(t *testing.T) {
	//float32 pair
	var r Register
	var z PackedFloat32
	z.Lower = TestingValue32bitFloatLower
	z.Upper = TestingValue32bitFloatUpper
	r.SetPackedFloat32Value(&z)
	//check the bits
	expectationCheckUint32(Float32bitsRawLower, uint32(r.Bits), "(lower half raw bits)", t)
	expectationCheckUint32(Float32bitsRawUpper, uint32(r.Bits>>32), "(upper half raw bits)", t)
	expectationCheckByte(RegisterTypePackedFloat32, r.Tag, "(tag comparison)", t)
}

func TestRegisterContents_5(t *testing.T) {
	//float32 pair
	var r Register
	r.SetFromInt32Values(TestingValue32bitLower, TestingValue32bitUpper)
	//check the bits
	expectationCheckUint32(TestingValue32bitLower, uint32(r.Bits), "(lower half raw bits)", t)
	expectationCheckUint32(TestingValue32bitUpper, uint32(r.Bits>>32), "(upper half raw bits)", t)
	expectationCheckByte(RegisterTypePackedInt32, r.Tag, "(tag comparison)", t)
}

func TestRegisterContents_6(t *testing.T) {
	//float32 pair
	var r Register
	var z PackedInt32
	z.Lower = TestingValue32bitLower
	z.Upper = TestingValue32bitUpper
	r.SetPackedInt32Value(&z)
	//check the bits
	expectationCheckUint32(TestingValue32bitLower, uint32(r.Bits), "(lower half raw bits)", t)
	expectationCheckUint32(TestingValue32bitUpper, uint32(r.Bits>>32), "(upper half raw bits)", t)
	expectationCheckByte(RegisterTypePackedInt32, r.Tag, "(tag comparison)", t)
}
