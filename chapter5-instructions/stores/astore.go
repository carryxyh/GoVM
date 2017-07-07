/**
	存储指令把变量从操作数栈顶弹出，然后存入局部变量表
 */
package stores

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type ASTORE struct {
	base.Index8Instruction
}

func (self *ASTORE) Execute(frame *chapter4_rtdt.Frame) {
	_astore(frame, self.Index)
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_0) Execute(frame *chapter4_rtdt.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_1) Execute(frame *chapter4_rtdt.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_2) Execute(frame *chapter4_rtdt.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_3) Execute(frame *chapter4_rtdt.Frame) {
	_astore(frame, 3)
}

func _astore(frame *chapter4_rtdt.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}