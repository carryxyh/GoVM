package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PopRef()
}
