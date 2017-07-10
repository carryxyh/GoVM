package math

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

// Divide double
type DDIV struct {
	base.NoOperandsInstruction
}

func (self *DDIV) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// Divide float
type FDIV struct {
	base.NoOperandsInstruction
}

func (self *FDIV) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// Divide int
type IDIV struct {
	base.NoOperandsInstruction
}

func (self *IDIV) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

// Divide long
type LDIV struct {
	base.NoOperandsInstruction
}

func (self *LDIV) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}