package data

func (obj *Object) Bytes() []int8 {
	return obj.data.([]int8)
}

func (obj *Object) Shorts() []int16 {
	return obj.data.([]int16)
}

func (obj *Object) Chars() []int16 {
	return obj.data.([]int16)
}

func (obj *Object) Ints() []int32 {
	return obj.data.([]int32)
}

func (obj *Object) Longs() []int64 {
	return obj.data.([]int64)
}

func (obj *Object) Floats() []float32 {
	return obj.data.([]float32)
}

func (obj *Object) Doubles() []float64 {
	return obj.data.([]float64)
}

func (obj *Object) Refs() []*Object {
	return obj.data.([]*Object)
}

func (obj *Object) ArrayLength() int32 {
	switch obj.data.(type) {
	case []int8:
		return int32(len(obj.data.([]int8)))
	case []int16:
		return int32(len(obj.data.([]int16)))
	case []int32:
		return int32(len(obj.data.([]int32)))
	case []int64:
		return int32(len(obj.data.([]int64)))
	case []float32:
		return int32(len(obj.data.([]float32)))
	case []float64:
		return int32(len(obj.data.([]float64)))
	case []*Object:
		return int32(len(obj.data.([]*Object)))
	default:
		panic("Not a array!")
	}
}

func (class *Class) NewArray(count uint) *Object {
	if !class.IsArray() {
		panic("Not a array: " + class.className)
	}

	switch class.className {
	case "[B":
		return &Object{class, make([]int8, count)}
	case "[S":
		return &Object{class, make([]int16, count)}
	case "[C":
		return &Object{class, make([]int16, count)}
	case "[I":
		return &Object{class, make([]int32, count)}
	case "[J":
		return &Object{class, make([]int64, count)}
	case "[F":
		return &Object{class, make([]float32, count)}
	case "[D":
		return &Object{class, make([]float64, count)}
	case "[Z":
		return &Object{class, make([]int8, count)}
	default:
		return &Object{class, make([]*Object, count)}
	}
}

func (class *Class) IsArray() bool {
	return class.className[0] == '['
}
