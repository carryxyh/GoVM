package comparisons

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	把栈顶的两个int变量弹出，然后进行比较，满足条件则跳转
	x == 0
 */
type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x != 0
 */
type IF_ICMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x < 0
 */
type IF_ICMPLT struct {
	base.BranchInstruction
}

func (self *IF_ICMPLT) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x <= 0
 */
type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPLE) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x > 0
 */
type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x >= 0
 */
type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}