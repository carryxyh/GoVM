package heap

type Object struct {
	class *Class
	data  interface{}
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
