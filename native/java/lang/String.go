package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *chapter4_rtdt.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}