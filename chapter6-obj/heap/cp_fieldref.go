package heap

import "GoVM/chapter3-cf/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, fieldrefInfo *chapter3_cf.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&fieldrefInfo.ConstantMemberrefInfo)
	return ref
}

/**
	解析字段
 */
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

/**
	如果类D想通过字段符号引用访问类C的某个字段，首先要解析符号引用得到类C，
	然后根据字段名和描述符查找字段。如果字段查找失败，抛出异常
	如果查找成功，但是D没用足够的权限访问字段，则抛出异常
 */
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

/**
	从类、接口、父类中依次寻找
 */
func lookupField(class *Class, name, descriptor string) *Field {
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range class.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}
	return nil
}