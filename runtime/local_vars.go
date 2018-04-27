package runtime

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

func (lv LocalVars) SetFloat32(index uint, val float32) {
	lv[index] = val
}

func (lv LocalVars) GetFloat32(index uint) float32 {
	return lv[index].(float32)
}

func (lv LocalVars) SetLong(index uint, val int64)  {
	lv[index] = val
}

func (lv LocalVars) GetLong(index uint) int64 {
	return lv[index].(int64)
}

func (lv LocalVars) SetDouble(index uint, val float64) {
	lv[index] = val
}

func (lv LocalVars) GetDouble(index uint) float64 {
	return lv[index].(float64)
}

func (lv LocalVars) SetRef(index uint, val *Object) {
	lv[index] = val
}

func (lv LocalVars) GetRef(index uint) *Object {
	return lv[index].(*Object)
}