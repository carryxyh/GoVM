/**
	存储指令把变量从操作数栈顶弹出，然后存入局部变量表
 */
package stores

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type LSTORE struct {
	base.Index8Instruction
}

func (self *LSTORE) Execute(frame *chapter4_rtdt.Frame) {
	_lstore(frame, self.Index)
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_0) Execute(frame *chapter4_rtdt.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_1) Execute(frame *chapter4_rtdt.Frame) {
	_lstore(frame, 0)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_2) Execute(frame *chapter4_rtdt.Frame) {
	_lstore(frame, 0)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_3) Execute(frame *chapter4_rtdt.Frame) {
	_lstore(frame, 0)
}

func _lstore(frame *chapter4_rtdt.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}