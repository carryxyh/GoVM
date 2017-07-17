package heap

import "GoVM/chapter3-cf/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, interfaceMethodRef *chapter3_cf.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&interfaceMethodRef.ConstantMemberrefInfo)
	return ref
}