package chapter4_rtdt

import "math"

//这也是一个栈结构，操作数栈
type OperandStack struct {
	//用于记录栈顶位置
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots:        make([]Slot, maxStack),
		}
	}
	return nil
}

//int操作
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

//float
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

//long变量入栈时，要拆成两个int。弹出时也弹两个int
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size + 1].num = int32(val >> 32)
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size + 1].num)
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
func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil // 把引用设置为nil -> 帮助Go回收结构体实例
	return ref
}

//Slot
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot(slot Slot) Slot {
	self.size--
	return self.slots[self.size]
}