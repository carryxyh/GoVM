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