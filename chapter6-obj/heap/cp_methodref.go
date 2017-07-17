package heap

import "GoVM/chapter3-cf/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, methodrefInfo *chapter3_cf.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&methodrefInfo.ConstantMemberrefInfo)
	return ref
}