package comparisons

/**
	if<cond> 指令把操作数栈顶的int变量弹出，然后跟0进行比较，满足条件则跳转
 */
import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	以下皆假设弹出的值是x，则指令跳转操作的条件如下:
	x == 0
 */
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *chapter4_rtdt.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x != 0
 */
type IFNE struct {
	base.BranchInstruction
}

func (self *IFNE) Execute(frame *chapter4_rtdt.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x < 0
 */
type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *chapter4_rtdt.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x <= 0
 */
type IFLE struct {
	base.BranchInstruction
}

func (self *IFLE) Execute(frame *chapter4_rtdt.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x > 0
 */
type IFGT struct {
	base.BranchInstruction
}

func (self *IFGT) Execute(frame *chapter4_rtdt.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

/**
	x >= 0
 */
type IFGE struct {
	base.BranchInstruction
}

func (self *IFGE) Execute(frame *chapter4_rtdt.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}