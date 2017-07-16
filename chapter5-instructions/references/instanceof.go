package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}

func (self *INSTANCE_OF) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	//操作数栈中的ref
	ref := stack.PopRef()
	if ref == nil {
		//返回false
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		//true
		stack.PushInt(1)
	} else {
		//false
		stack.PushInt(0)
	}
}