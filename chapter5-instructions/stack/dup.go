package stack

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	复制栈顶的单个变量
 */
type DUP struct {
	base.NoOperandsInstruction
}

/**
	执行前
		栈底 1 栈顶
	执行后
		栈底 1 1 栈顶
 */
func (self *DUP) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// Duplicate the top operand stack value and insert two values down
type DUP_X1 struct {
	base.NoOperandsInstruction
}

/**
	执行前
		栈底 2 1 栈顶
	执行后
		栈底 1 2 1 栈顶
 */
func (self *DUP_X1) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct {
	base.NoOperandsInstruction
}

/**
	执行前
		栈底 3 2 1 栈顶
	执行后
		栈底 1 3 2 1 栈顶
 */
func (self *DUP_X2) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values
type DUP2 struct {
	base.NoOperandsInstruction
}

/**
	执行前
		栈底 2 1 栈顶
	执行后
		栈底 2 1 2 1 栈顶
 */
func (self *DUP2) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

/**
	执行前
		栈底 3 2 1 栈顶
	执行后
		栈底 2 1 3 2 1 栈顶
 */
func (self *DUP2_X1) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

/**
	执行前
		栈底 4 3 2 1 栈顶
	执行后
		栈底 2 1 4 3 2 1 栈顶
 */
func (self *DUP2_X2) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}