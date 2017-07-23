package loads

import (
	"GoVM/chapter6-obj/heap"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter5-instructions/base"
)

/**
	xaload系列指令：
		按索引取数组元素，然后推入操作数栈
 */

// Load reference from array
type AALOAD struct {
	base.NoOperandsInstruction
}

func (self *AALOAD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

// Load byte or boolean from array
type BALOAD struct {
	base.NoOperandsInstruction
}

func (self *BALOAD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

// Load char from array
type CALOAD struct {
	base.NoOperandsInstruction
}

func (self *CALOAD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

// Load double from array
type DALOAD struct {
	base.NoOperandsInstruction
}

func (self *DALOAD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

// Load float from array
type FALOAD struct {
	base.NoOperandsInstruction
}

func (self *FALOAD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

// Load int from array
type IALOAD struct {
	base.NoOperandsInstruction
}

func (self *IALOAD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

// Load long from array
type LALOAD struct {
	base.NoOperandsInstruction
}

func (self *LALOAD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

// Load short from array
type SALOAD struct {
	base.NoOperandsInstruction
}

func (self *SALOAD) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
