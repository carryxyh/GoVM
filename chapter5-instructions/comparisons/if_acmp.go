package comparisons

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	obj == obj
 */
type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopRef()
	val1 := stack.PopRef()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

/**
	obj != obj
 */
type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopRef()
	val1 := stack.PopRef()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}