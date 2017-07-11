package chapter3_cf

import "math"

/**
	常量池中integer
	四字节存储整数常量
	CONSTANT_INTEGER_INFO {
		u1 tag;
		u4 bytes;
	}
 */
type ConstantIntegerInfo struct {
	//实际上，比int小的boolean、byte、short、char也可以放在里面
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

/**
	常量池中float
	四字节
	CONSTANT_FLOAT_INFO {
		u1 tag;
		u4 bytes;
	}
 */
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

/**
	常量池中long
	特殊一些 八字节，分成高8字节和低8字节
	CONSTANT_LONG_INFO {
		u1 tag;
		u4 high_bytes;
		u4 low_bytes;
	}
 */
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

/**
	常量池中double
	同样特殊 八字节
	CONSTANT_DOUBLE_INFO {
		u1 tag;
		u4 high_bytes;
		u4 low_bytes;
	}
 */
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}