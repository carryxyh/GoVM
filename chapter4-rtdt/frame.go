package chapter4_rtdt

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
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:        newLocalVars(maxLocals),
		operandStack:        newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}