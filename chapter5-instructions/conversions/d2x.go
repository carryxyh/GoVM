package conversions

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}