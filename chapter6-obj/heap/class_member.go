package heap

import "GoVM/chapter3-cf/classfile"

/**
	存放字段和方法公有的信息
 */
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	//主要为了通过字段或方法访问到它所属的类
	class       *Class
}

/**
	从class文件中复制数据
 */
func (self *ClassMember) copyMemberInfo(memberInfo *chapter3_cf.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags & ACC_PRIVATE
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags & ACC_PROTECTED
}

func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags & ACC_STATIC
}

func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.GetPackageName() == d.GetPackageName()
	}
	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}