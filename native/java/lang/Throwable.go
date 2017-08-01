package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

const jlThrowable = "java/lang/Throwable"

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func fillInStackTrace(frame *chapter4_rtdt.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

/**
	由于栈顶两帧正在执行fillInStackTrace(int)和fillInStackTrace()方法，所以需要跳过这两帧
	这两帧下面的几帧正在执行异常类的构造函数，所以也要跳过，具体跳过多少帧要看异常类的层次
 */
func createStackTraceElements(obj *heap.Object, thread *chapter4_rtdt.Thread) []*StackTraceElement {
	skip := distanceToObject(obj.Class()) + 2
	frames := thread.GetFrames()[skip:]

	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func createStackTraceElement(frame *chapter4_rtdt.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()

	return &StackTraceElement{
		fileName:	class.SourceFile(),
		className:	class.JavaName(),
		methodName:	method.Name(),
		lineNumber:	method.GetLineNumber(frame.NextPC() - 1),
	}
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}