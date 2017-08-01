package chapter4_rtdt

import (
	"math"
	"GoVM/chapter6-obj/heap"
)

//这也是一个栈结构，操作数栈
type OperandStack struct {
	//用于记录栈顶位置
	size  uint
	slots []heap.Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots:        make([]heap.Slot, maxStack),
		}
	}
	return nil
}

func (self *OperandStack) Clear() {
	self.size = 0
	for i := range self.slots {
		self.slots[i].Ref = nil
	}
}

/**
	返回距离操作数栈顶n个单元格的引用变量。
	n = 0 返回操作数占顶引用
	n = 1 返回制定开始的第二个引用
 */
func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return self.slots[self.size - 1 - n].Ref
}

func (self *OperandStack) PushBoolean(val bool) {
	if val {
		self.PushInt(1)
	} else {
		self.PushInt(0)
	}
}

func (self *OperandStack) PopBoolean() bool {
	return self.PopInt() == 1
}

//int操作
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].Num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].Num
}

//float
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].Num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].Num)
	return math.Float32frombits(bits)
}

//long变量入栈时，要拆成两个int。弹出时也弹两个int
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].Num = int32(val)
	self.slots[self.size + 1].Num = int32(val >> 32)
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].Num)
	high := uint32(self.slots[self.size + 1].Num)
	return int64(high) << 32 | int64(low)
}

//double先变成long，然后按照long处理
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

//引用类型
func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size].Ref = ref
	self.size++
}

func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].Ref
	self.slots[self.size].Ref = nil // 把引用设置为nil -> 帮助Go回收结构体实例
	return ref
}

//Slot
func (self *OperandStack) PushSlot(slot heap.Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() heap.Slot {
	self.size--
	return self.slots[self.size]
}