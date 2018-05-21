package data

import "math"

// this is a slice not a arrays.
// arrays size is set when init
type Slot interface {
}

type Variables []Slot

func NewVariables(maxLocals uint) Variables {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (v Variables) SetInt(index uint, val int32) {
	v[index] = val
}

func (v Variables) GetInt(index uint) int32 {
	return v[index].(int32)
}

func (v Variables) SetFloat(index uint, val float32) {
	v[index] = val
}

func (v Variables) GetFloat(index uint) float32 {
	return v[index].(float32)
}

func (v Variables) SetLong(index uint, val int64) {
	v[index] = int32(val)
	v[index+1] = int32(val >> 32)
}

func (v Variables) GetLong(index uint) int64 {
	low := uint32(v[index].(int32))
	high := uint32(v[index+1].(int32))
	return int64(high)<<32 | int64(low)
}

func (v Variables) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	v.SetLong(index, int64(bits))
}

func (v Variables) GetDouble(index uint) float64 {
	bits := v.GetLong(index)
	return math.Float64frombits(uint64(bits))
}

func (v Variables) SetRef(index uint, val *Object) {
	v[index] = val
}

func (v Variables) GetRef(index uint) *Object {
	val := v[index]
	if val != nil {
		return val.(*Object)
	}
	return nil
}

func (v Variables) SetSlot(index uint, val Slot) {
	v[index] = val
}

func (v Variables) GetSlot(index uint) Slot {
	return v[index]
}
