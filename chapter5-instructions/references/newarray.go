package references

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

const (
	//Array Type  atype
	AT_BOOLEAN = 4
	AT_CHAR = 5
	AT_FLOAT = 6
	AT_DOUBLE = 7
	AT_BYTE = 8
	AT_SHORT = 9
	AT_INT = 10
	AT_LONG = 11
)

// Create new array
type NEW_ARRAY struct {
	//标识创建哪种类型的数组
	atype uint8
}

func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUInt8()
}

func (self *NEW_ARRAY) Execute(frame *chapter4_rtdt.Frame) {
	stack := frame.OperandStack()
	//从操作数栈中弹出，标识数组长度
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	classLoader := frame.Method().Class().Loader()
	//根据atype值使用当前类的加载器加载数组
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}