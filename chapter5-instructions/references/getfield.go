package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

type GET_FIELD struct {
	base.Index16Instruction
}

func (self *GET_FIELD) Execute(frame *chapter4_rtdt.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncopatibleClassChangeError")
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
	// todo
	}
}