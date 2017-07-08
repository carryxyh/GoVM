package math

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"math"
)

/**
	取余指令
 */
type DREM struct {
	base.NoOperandsInstruction
}

/**
	浮点数有Infinity -> 无限大值，所以即使是除0 也不会ArithmeticException
 */
func (self *DREM) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

type LREM struct {
	base.NoOperandsInstruction
}

func (self *LREM) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}