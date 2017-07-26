package constants

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

/**
	用于从运行时常量池中加载常量值，并把它推入到操作数栈
 */
type LDC struct {
	base.Index8Instruction
}

func (self *LDC) Execute(frame *chapter4_rtdt.Frame) {
	_ldc(frame, self.Index)
}

/**
	与ldc指令一样，都是加载int、float和字符串常量
 */
type LDC_W struct {
	base.Index16Instruction
}

func (self *LDC_W) Execute(frame *chapter4_rtdt.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *chapter4_rtdt.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	default:
		panic("todo: ldc")
	}
}


/**
	用于加载long、double常量
 */
type LDC2_W struct {
	base.Index16Instruction
}

func (self *LDC2_W) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
