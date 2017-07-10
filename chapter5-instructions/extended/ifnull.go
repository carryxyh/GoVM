package extended

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	根据引用是否为null来跳转
	ref == null
 */
type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *chapter4_rtdt.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

/**
	ref != null
 */
type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNONNULL) Execute(frame *chapter4_rtdt.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}