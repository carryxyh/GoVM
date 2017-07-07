package constants

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *chapter4_rtdt.Frame) {
	//do noting
}
