package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

/**
	跟instanceof非常像，区别：
		instanceof会改变操作数栈（弹出引用，推入判断结果）
		checkcast不改变操作数栈（如果判断失败，直接抛出ClassCastException）
 */
type CHECK_CAST struct {
	base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	//操作数栈中的ref
	ref := stack.PopRef()
	//不改变操作数栈
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
