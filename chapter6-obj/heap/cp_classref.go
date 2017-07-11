package heap

import "GoVM/chapter3-cf/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo chapter3_cf.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}