package comparisons

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	浮点数计算有可能产生NAN not a number ， 所以比较两个浮点数时，除了大于等于小于之外还无法比较
 */
type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) Execute(frame *chapter4_rtdt.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *chapter4_rtdt.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *chapter4_rtdt.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}