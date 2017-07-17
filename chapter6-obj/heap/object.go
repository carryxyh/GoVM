package heap

type Object struct {
	class  *Class
	fields Slots
}

func NewObject(class *Class) *Object {
	return &Object{
		class:        class,
		fields: NewSlots(class.InstanceSlotCount),
	}
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(self.class)
}

func (self *Object) Fields() Slots {
	return self.fields
}
