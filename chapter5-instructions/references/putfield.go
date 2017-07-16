package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

/**
	需要三个操作数。
	前两个操作数是常量池索引和变量值，用法和putstatic一样。
	第三个操作数是对象的应用，从操作数中弹出。
 */
type PUT_FIELD struct {
	base.Index16Instruction
}

func (self *PUT_FIELD) Execute(frame *chapter4_rtdt.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	//如果是static修饰的 抛出异常
	if field.IsStatic() {
		panic("java.lang.IncopatibleClassChangeError")
	}

	//如果是final修饰的，但是赋值没有放在init里，抛出异常
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
	// todo
	}
}