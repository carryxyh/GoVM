package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
	"unsafe"
)

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
	native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
}

//public final native Class<?> getClass
func getClass(frame *chapter4_rtdt.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *chapter4_rtdt.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *chapter4_rtdt.Frame) {
	this := frame.LocalVars().GetThis()

	clonable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(clonable) {
		panic("java.lang.CloneNotSupportException")
	}

	frame.OperandStack().PushRef(this.Clone())
}