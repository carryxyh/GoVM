package math

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type IAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

type LAND struct {
	base.NoOperandsInstruction
}

func (self *LAND) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushInt(result)
}
