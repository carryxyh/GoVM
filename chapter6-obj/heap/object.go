package heap

type Object struct {
	class *Class
	data  interface{}
	/**
		用来记录Object结构体的额外信息:
			1.某个类的对象 对应的 Class结构体指针，这里的这个Class是JVM方法区中的Class结构体 -> heap.Class
	 */
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

//getter start
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
//getter end

// reflection
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

/**
	获取一个对象中的某个字段的值
 */
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}