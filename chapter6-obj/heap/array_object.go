package heap

/**
	将使用int8型来标识boolean类型
 */
func (self *Object) Bytes() []int8 {
	return self.data.([]int8)
}

func (self *Object) Shorts() []int16 {
	return self.data.([]int16)
}

func (self *Object) Ints() []int32 {
	return self.data.([]int32)
}

func (self *Object) Longs() []int64 {
	return self.data.([]int64)
}

func (self *Object) Chars() []uint16 {
	return self.data.([]uint16)
}

func (self *Object) Floats() []float32 {
	return self.data.([]float32)
}

func (self *Object) Doubles() []float64 {
	return self.data.([]float64)
}

func (self *Object) Refs() []*Object {
	return self.data.([]*Object)
}

func ArrayCopy(src, dest *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos : srcPos + length]
		_dst := dest.data.([]int8)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos : srcPos + length]
		_dst := dest.data.([]int16)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos + length]
		_dst := dest.data.([]int32)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []int64:
		_src := src.data.([]int64)[srcPos : srcPos + length]
		_dst := dest.data.([]int64)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos : srcPos + length]
		_dst := dest.data.([]uint16)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos : srcPos + length]
		_dst := dest.data.([]float32)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos : srcPos + length]
		_dst := dest.data.([]float64)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos + length]
		_dst := dest.data.([]*Object)[dstPos : dstPos + length]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}

/**
	数组长度
 */
func (self *Object) ArrayLength() int32 {
	switch self.data.(type) {
	case []int8:
		return int32(len(self.data.([]int8)))
	case []int16:
		return int32(len(self.data.([]int16)))
	case []int32:
		return int32(len(self.data.([]int32)))
	case []int64:
		return int32(len(self.data.([]int64)))
	case []uint16:
		return int32(len(self.data.([]uint16)))
	case []float32:
		return int32(len(self.data.([]float32)))
	case []float64:
		return int32(len(self.data.([]float64)))
	case []*Object:
		return int32(len(self.data.([]*Object)))
	default:
		panic("Not array!")
	}
}