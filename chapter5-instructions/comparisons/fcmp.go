package comparisons

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	浮点数计算有可能产生NAN not a number ， 所以比较两个浮点数时，除了大于等于小于之外还无法比较
 */
type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *chapter4_rtdt.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *chapter4_rtdt.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *chapter4_rtdt.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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