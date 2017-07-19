package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *chapter4_rtdt.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	//解析出来的方法的名字如果是init 那么 调用这个方法的class必须是自己
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//从操作数栈中弹出this引用，如果为null 抛异常
	//注意，在传递参数之前，不能破坏操作数栈的状态。
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	//确保protected方法只能被该放的类或子类调用
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) && resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() && ref.Class() != currentClass && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	//如果调用的超类中的方法，但不是构造方法，且当前累的ACC_SUPER标志被设置，需要一个额外的过程查找最重要调用的方法；
	//否则前面从放方法符号中解析出来的方法就是要调用的方法
	toBeInvoked := resolvedMethod
	if currentClass.IsSuper() && resolvedClass.IsSuperClassOf(currentClass) && resolvedMethod != "<init>" {
		toBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}

	if toBeInvoked == nil || toBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, toBeInvoked)
}
