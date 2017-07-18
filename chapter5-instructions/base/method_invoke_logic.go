package base

import (
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

/**
	重点是参数传递：
	1.确定方法的参数在局部变量表中占用多少位置，注意，这个数量不一定等于java代码中看到的参数个数，原因：
		1.long和double类型的参数要占用两个位置
		2.对于实例方法，java编译器会在参数列表的前面添加一个参数，这个参数的类型就是this引用
	2.假设实际的参数占用 n 个位置，一次把这 n 个变量从调用者的操作数栈中弹出，放进调用方法的局部变量表中
 */
func InvokeMethod(invokerFrame *chapter4_rtdt.Frame, method *heap.Method) {
	//创建一个新的栈帧，并且压入线程的栈顶
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			//调用方操作数栈中弹出参数
			slot := invokerFrame.OperandStack().PopSlot()
			//放到方法栈帧的局部变量表中
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
