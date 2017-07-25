package heap

type Object struct {
	class *Class
	data  interface{}
	//用来记录Object结构体的额外信息
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class:        class,
		data: NewSlots(class.InstanceSlotCount),
	}
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(self.class)
}

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

// reflection
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}