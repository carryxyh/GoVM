package comparisons

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	比较long
 */
type LCMP struct {
	base.NoOperandsInstruction
}

/**
	注意这里需要把比较结果推到栈顶
 */
func (self *LCMP) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
