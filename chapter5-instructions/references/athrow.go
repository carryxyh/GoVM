package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
	"reflect"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (self *ATHROW) Execute(frame *chapter4_rtdt.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

/**
	从当前帧开始寻找方法的异常处理表，如果找不到，弹出栈帧继续寻找
	如果找到了，跳转到异常处理之前，先把栈帧的操作数栈清空，然后把异常对象引用推入栈顶

 */
func findAndGotoExceptionHandler(thread *chapter4_rtdt.Thread, ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handlerPc := frame.Method().FindExceptionHandler(ex, pc)
		if handlerPc > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPc)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

/**
	把虚拟机栈清空，打印异常信息，虚拟机栈空了，解释器也就停止了
	异常对象的extra字段中，存放的就是java虚拟机栈新消息
 */
func handleUncaughtException(thread *chapter4_rtdt.Thread, ex *heap.Object) {
	thread.ClearStack()

	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)

	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			string() string
		})
		println("\tat " + ste.string())
	}
}