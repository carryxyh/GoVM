package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

const jlClass = "java/lang/Class"

func init() {
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	native.Register(jlClass, "isInterface", "()Z", isInterface)
}

func getPrimitiveClass(frame *chapter4_rtdt.Frame) {
	nameObj := frame.LocalVars().GetThis()
	name := heap.GoString(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

func getName0(frame *chapter4_rtdt.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

/**
	断言相关
 */
func desiredAssertionStatus0(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushBoolean(false)
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *chapter4_rtdt.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
//func isPrimitive(frame *chapter4_rtdt.Frame) {
//	vars := frame.LocalVars()
//	this := vars.GetThis()
//	class := this.Extra().(*heap.Class)
//
//	stack := frame.OperandStack()
//	stack.PushBoolean(class.IsPrimitive())
//}