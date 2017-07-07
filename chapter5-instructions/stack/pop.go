/**
	栈指令直接对操作数栈进行操作
	pop和pop2指令将制定变量弹出
	dup系列指令复制栈顶变量
	swap指令交换栈顶的两个变量
 */
package stack

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	只能用于弹出int float等占用一个操作数栈未知的变量
 */
type POP struct {
	base.NoOperandsInstruction
}

/**
	用于弹出double long变量
 */
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}