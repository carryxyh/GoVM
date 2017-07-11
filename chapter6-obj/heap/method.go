package heap

import "GoVM/chapter3-cf/classfile"

type Method struct {
	//继承自ClassMember，字段和方法都属于类的成员
	ClassMember
	code      []byte

	//以下两个值是由java编译器计算好了的
	maxStack  uint
	maxLocals uint
}

func newMethods(class *Class, cfMethods []*chapter3_cf.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

/**
	从classFile中读取方法相关信息
 */
func (self *Method) copyAttributes(cfMethod *chapter3_cf.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.code = codeAttr.Code()
		self.maxLocals = codeAttr.MaxLocals()
	}
}