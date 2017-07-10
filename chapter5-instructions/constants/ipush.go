package constants

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	从操作数中获取一个byte型证书，扩展成int型，然后推入栈顶
 */
type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *chapter4_rtdt.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

/**
	从操作数中获取一个short型证书，扩展成int型，然后推入栈顶
 */
type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *chapter4_rtdt.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
