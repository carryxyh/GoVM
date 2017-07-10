package math

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

// Multiply double
type DMUL struct {
	base.NoOperandsInstruction
}

func (self *DMUL) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

// Multiply float
type FMUL struct {
	base.NoOperandsInstruction
}

func (self *FMUL) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

// Multiply int
type IMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

// Multiply long
type LMUL struct {
	base.NoOperandsInstruction
}

func (self *LMUL) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
