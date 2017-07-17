package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

/**
	操作的是一个uint16的索引，来自字节码。通过这个索引，可以从当前类的运行时常量池中找到一个类符号引用。
	解析这个类符号引用，拿到类数据
	然后创建对象，并把对象引用推入栈顶
 */
type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *chapter4_rtdt.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantianionError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}


