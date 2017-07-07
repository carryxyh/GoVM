package loads

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

// Load double from local variable
type DLOAD struct {
	base.Index8Instruction
}

func (self *DLOAD) Execute(frame *chapter4_rtdt.Frame) {
	_dload(frame, uint(self.Index))
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_0) Execute(frame *chapter4_rtdt.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_1) Execute(frame *chapter4_rtdt.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_2) Execute(frame *chapter4_rtdt.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_3) Execute(frame *chapter4_rtdt.Frame) {
	_dload(frame, 3)
}

func _dload(frame *chapter4_rtdt.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
