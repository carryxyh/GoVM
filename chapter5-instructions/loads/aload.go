package loads

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

// Load reference from local variable
type ALOAD struct {
	base.Index8Instruction
}

func (self *ALOAD) Execute(frame *chapter4_rtdt.Frame) {
	_aload(frame, uint(self.Index))
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_0) Execute(frame *chapter4_rtdt.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_1) Execute(frame *chapter4_rtdt.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_2) Execute(frame *chapter4_rtdt.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_3) Execute(frame *chapter4_rtdt.Frame) {
	_aload(frame, 3)
}

func _aload(frame *chapter4_rtdt.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}
