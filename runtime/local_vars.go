package runtime

import "math"

// this is a slice not a array.
// array size is set when init
type LocalVars []interface{}

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]interface{}, maxLocals)
	}
	return nil
}

func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index] = val
}

func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].(int32)
}

func (lv LocalVars) SetFloat(index uint, val float32) {
	lv[index] = val
}

func (lv LocalVars) GetFloat(index uint) float32 {
	return lv[index].(float32)
}

func (lv LocalVars) SetLong(index uint, val int64)  {
	lv[index] = int32(val)
	lv[index + 1] = int32(val >> 32)
}

func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].(int32))
	high := uint32(lv[index + 1].(int32))
	return int64(high) << 32 | int64(low)
}

func (lv LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.SetLong(index, int64(bits))
}

func (lv LocalVars) GetDouble(index uint) float64 {
	bits := lv.GetLong(index)
	return math.Float64frombits(uint64(bits))
}

func (lv LocalVars) SetRef(index uint, val *Object) {
	lv[index] = val
}

func (lv LocalVars) GetRef(index uint) *Object {
	return lv[index].(*Object)
}