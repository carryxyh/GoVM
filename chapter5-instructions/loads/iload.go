/**
	加载指令：从局部变量表获取变量，然后推入操作数栈顶
	aload 系列指令操作引用类型变量
	dload 操作double
	fload 操作float
	iload 操作int
	lload 操作long
	xaload 操作数组
 */
package loads

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

// Load int from local variable
type ILOAD struct {
	base.Index8Instruction
}

func (self *ILOAD) Execute(frame *chapter4_rtdt.Frame) {
	_iload(frame, uint(self.Index))
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_0) Execute(frame *chapter4_rtdt.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_1) Execute(frame *chapter4_rtdt.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_2) Execute(frame *chapter4_rtdt.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_3) Execute(frame *chapter4_rtdt.Frame) {
	_iload(frame, 3)
}

func _iload(frame *chapter4_rtdt.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
