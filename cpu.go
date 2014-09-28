// the edgeworth system CPU - a simple 64-bit RISC cpu with a large compliment of registers, most of them purpose built
package edgeworth

import (
	"fmt"
)

const (
	RegisterCount = 256
)

type InvalidRegisterTypeError struct {
	Type string
}

func (e InvalidRegisterTypeError) Error() string {
	return fmt.Sprintf("Error: invalid register type: %s provided for integer register", e.Type)
}

type Register interface {
	GetValue() interface{}
	SetValue(interface{}) error
}

type IntegerRegister struct {
	value uint64
}

type FloatRegister struct {
	value float64
}

func (r IntegerRegister) GetValue() interface{} {
	return r.value
}

func (r IntegerRegister) SetValue(v interface{}) error {

	switch t := v.(type) {
	default:
		var q InvalidRegisterTypeError
		q.Type = fmt.Sprintf("%T", t)
		return &q
	case uint8:
		r.value = v.(uint64)
		return nil
	case int8:
		r.value = v.(uint64)
		return nil
	case uint16:
		r.value = v.(uint64)
		return nil
	case int16:
		r.value = v.(uint64)
		return nil
	case int32:
		r.value = v.(uint64)
		return nil
	case uint32:
		r.value = v.(uint64)
		return nil
	case int64:
		r.value = v.(uint64)
		return nil
	case uint64:
		r.value = v.(uint64)
		return nil
	}
}

func (r FloatRegister) GetValue() interface{} {
	return r.value
}

func (r FloatRegister) SetValue(v interface{}) error {

	switch t := v.(type) {
	default:
		var q InvalidRegisterTypeError
		q.Type = fmt.Sprintf("%T", t)
		return &q
	case float32:
		r.value = v.(float64)
		return nil
	case float64:
		r.value = v.(float64)
		return nil
	}
}
