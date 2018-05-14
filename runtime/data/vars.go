package data

import "math"

// this is a slice not a array.
// array size is set when init
type Slot struct {
	val interface{}
}

type Variables []Slot

func NewVariables(maxLocals uint) Variables {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (lv Variables) SetInt(index uint, val int32) {
	lv[index] = Slot{val}
}

func (lv Variables) GetInt(index uint) int32 {
	return lv[index].val.(int32)
}

func (lv Variables) SetFloat(index uint, val float32) {
	lv[index] = Slot{val}
}

func (lv Variables) GetFloat(index uint) float32 {
	return lv[index].val.(float32)
}

func (lv Variables) SetLong(index uint, val int64)  {
	lv[index] = Slot{int32(val)}
	lv[index + 1] = Slot{int32(val >> 32)}
}

func (lv Variables) GetLong(index uint) int64 {
	low := uint32(lv[index].val.(int32))
	high := uint32(lv[index + 1].val.(int32))
	return int64(high) << 32 | int64(low)
}

func (lv Variables) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.SetLong(index, int64(bits))
}

func (lv Variables) GetDouble(index uint) float64 {
	bits := lv.GetLong(index)
	return math.Float64frombits(uint64(bits))
}

func (lv Variables) SetRef(index uint, val *Object) {
	lv[index] = Slot{val}
}

func (lv Variables) GetRef(index uint) *Object {
	return lv[index].val.(*Object)
}