package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
)

const jlObject = "java/lang/Object"

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

//public final native Class<?> getClass
func getClass(frame *chapter4_rtdt.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
