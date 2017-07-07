/**
	存储指令把变量从操作数栈顶弹出，然后存入局部变量表
 */
package stores

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type FSTORE struct {
	base.Index8Instruction
}

func (self *FSTORE) Execute(frame *chapter4_rtdt.Frame) {
	_fstore(frame, self.Index)
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_0) Execute(frame *chapter4_rtdt.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_1) Execute(frame *chapter4_rtdt.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_2) Execute(frame *chapter4_rtdt.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_3) Execute(frame *chapter4_rtdt.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *chapter4_rtdt.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}