package base

type BytecodeReader struct {
	//字节码
	code []byte
	//标记读取到了哪个字节
	pc   int
}

/**
	为了避免每次解码指令都新创建一个BytecodeReader实例
 */
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUInt8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUInt8())
}

func (self *BytecodeReader) ReadUInt16() uint16 {
	b1 := uint16(self.ReadUInt8())
	b2 := uint16(self.ReadUInt8())
	return (b1 << 8) | b2
}

func (self *BytecodeReader) ReadInt16() uint16 {
	return int16(self.ReadUInt16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	b1 := int32(self.ReadUInt8())
	b2 := int32(self.ReadUInt8())
	b3 := int32(self.ReadUInt8())
	b4 := int32(self.ReadUInt8())
	return (b1 << 24) | (b2 << 16) | (b3 << 8) | b4
}

func (self *BytecodeReader) SkipPadding() {
	for self.pc % 4 != 0 {
		self.ReadUInt8()
	}
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) ReadInt32s(len int32) []int32 {
	ints := make([]int32, len)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}