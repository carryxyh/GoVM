package chapter4_rtdt

import "GoVM/chapter6-obj/heap"

type Object struct {
	class  *heap.Class
	fields Slots
}

func NewObject(class *heap.Class) *Object {
	return &Object{
		class:        class,
		fields: NewSlots(class.InstanceSlotCount),
	}
}