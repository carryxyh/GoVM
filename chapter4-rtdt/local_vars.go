package chapter4_rtdt

import "math"

/**
	局部变量表，按照索引访问
	存放基本类型或者引用类型
	注意：我们并没有针对boolean、byte、short和char类型定义存取方法，这些类型都可以转换成int值类处理
 */
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) SetInt(index uint, val uint32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

//先转成int，然后按照int变量来处理
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

//LONG类型需要处理成两个int
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index + 1].num = int32(val >> 32)
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index + 1].num)
	return int64(high) << 32 | int64(low)
}

//DOUBLE类型可以先转成LONG，然后按照LONG变量来处理
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, uint64(bits))
}

func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self LocalVars) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

func (self LocalVars) GetRef(index uint) *Object {
	return self[index].ref
}