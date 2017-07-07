/**
	存储指令把变量从操作数栈顶弹出，然后存入局部变量表
 */
package stores

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type ISTORE struct {
	base.Index8Instruction
}

func (self *ISTORE) Execute(frame *chapter4_rtdt.Frame) {
	_istore(frame, self.Index)
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_0) Execute(frame *chapter4_rtdt.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_1) Execute(frame *chapter4_rtdt.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_2) Execute(frame *chapter4_rtdt.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_3) Execute(frame *chapter4_rtdt.Frame) {
	_istore(frame, 3)
}

func _istore(frame *chapter4_rtdt.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}