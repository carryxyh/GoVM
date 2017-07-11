package heap

import "GoVM/chapter3-cf/classfile"

type Field struct {
	//继承自ClassMember，字段和方法都属于类的成员
	ClassMember
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
