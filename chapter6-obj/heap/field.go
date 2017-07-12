package heap

import "GoVM/chapter3-cf/classfile"

type Field struct {
	//继承自ClassMember，字段和方法都属于类的成员
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*chapter3_cf.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}