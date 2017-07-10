package math

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

// Boolean XOR int
type IXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long
type LXOR struct {
	base.NoOperandsInstruction
}

func (self *LXOR) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
