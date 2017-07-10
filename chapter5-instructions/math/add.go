package math

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

// Add double
type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// Add float
type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// Add int
type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

// Add long
type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}

