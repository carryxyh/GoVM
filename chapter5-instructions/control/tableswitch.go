package control

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type TABLE_SWITCH struct {
	//默认情况下执行跳转所需的字节码的偏移量
	defaultOffset int32
	//low high 记录case的取值范围
	low           int32
	high          int32
	//索引表，存放 high - low + 1 个int值，对应各种case情况下，执行跳转所需要的偏移量
	jumpOffsets   []int32
}

/**
	TABLE_SWITCH指令操作码后面有 0~3 个字节的padding，保证defaultOffset在字节码中的地址是4的倍数
 */
func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

/**
	弹出一个int看看是否在 low - high 范围内 不在走default 在走对应的偏移量
 */
func (self *TABLE_SWITCH) Execute(frame *chapter4_rtdt.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index - self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}