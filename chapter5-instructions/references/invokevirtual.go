package references

import (
	"GoVM/chapter4-rtdt"
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter6-obj/heap"
	"fmt"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

/**
	frame.Method()指的是当前的方法，要执行的方法应该用
	toBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	来获取。

	举个例子：
	void A() {
		B();
	}

	这时候frame.Method()返回的是A，要执行B，这里会调用invokeVirtual指令
	这时候要根据invokeVirtual的index找到运行时常量池中的methodRef
	然后再执行，即heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())会返回B
**/
func (self *INVOKE_VIRTUAL) Execute(frame *chapter4_rtdt.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		// hack!
		//if methodRef.Name() == "println" {
		//	_println(frame.OperandStack(), methodRef.Descriptor())
		//	return
		//}

		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	toBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if toBeInvoked == nil || toBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, toBeInvoked)
}

// hack!
func _println(stack *chapter4_rtdt.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		goStr := heap.GoString(jStr)
		fmt.Println(goStr)
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
