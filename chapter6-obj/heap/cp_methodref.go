package heap

import "GoVM/chapter3-cf/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, methodrefInfo *chapter3_cf.ConstantMethodrefInfo) *FieldRef {
	ref := &MemberRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&methodrefInfo.ConstantMemberrefInfo)
	return ref
}