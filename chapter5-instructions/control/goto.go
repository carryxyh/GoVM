package control

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	无条件跳转
 */
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *chapter4_rtdt.Frame) {
	base.Branch(frame, self.Offset)
}
