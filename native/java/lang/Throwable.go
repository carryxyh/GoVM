package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
)

const jlThrowable = "java/lang/Throwable"

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *chapter4_rtdt.Frame) {

}