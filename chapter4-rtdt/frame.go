package chapter4_rtdt

import "GoVM/chapter6-obj/heap"

/**
	栈帧
 */
type Frame struct {
	//用来实现链表的存储结构
	lower        *Frame
	//保存局部变量表指针
	localVars    LocalVars
	//保存操作数栈指针
	operandStack *OperandStack
	//这两个指令为了实现跳转
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:        thread,
		localVars:        newLocalVars(maxLocals),
		operandStack:        newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

func (self *Frame) Thread() *Thread {
	return self.thread
}