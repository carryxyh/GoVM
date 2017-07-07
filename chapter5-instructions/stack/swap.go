package stack

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	用于交换栈顶的两个变量
 */
type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
