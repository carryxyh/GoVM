package conversions

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

// Convert long to double
type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2D) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// Convert long to float
type L2F struct {
	base.NoOperandsInstruction
}

func (self *L2F) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// Convert long to int
type L2I struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}