package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

/**
	和其他三条invoke指令不同，invokeinterface后面跟着 4 字节操作符 不是2个字节
	前两个字节的含义和其他invoke一样，是个运行时常量池的索引
	第三个字节的给方法传递参数需要的slot数 含义和Method结构体的argSlotCount一样
	第四个字节是给oracle的java虚拟机实现的，值必须是0
 */
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUInt16())
	//放 count 和 zero 字段
	reader.ReadUInt8()
	reader.ReadUInt8()
}

func (self *INVOKE_INTERFACE) Execute(frame *chapter4_rtdt.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	toBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if toBeInvoked == nil || toBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !toBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, toBeInvoked)
}