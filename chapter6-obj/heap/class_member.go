package heap

import "GoVM/chapter3-cf/classfile"

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
