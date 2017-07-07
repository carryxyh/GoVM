/**
	存储指令把变量从操作数栈顶弹出，然后存入局部变量表
 */
package stores

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type DSTORE struct {
	base.Index8Instruction
}

func (self *DSTORE) Execute(frame *chapter4_rtdt.Frame) {
	_dstore(frame, self.Index)
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_0) Execute(frame *chapter4_rtdt.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_1) Execute(frame *chapter4_rtdt.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_2) Execute(frame *chapter4_rtdt.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_3) Execute(frame *chapter4_rtdt.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *chapter4_rtdt.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}