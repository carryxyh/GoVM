package heap

import "GoVM/chapter3-cf/classfile"

type Method struct {
	//继承自ClassMember，字段和方法都属于类的成员
	ClassMember
	code         []byte
	argSlotCount uint

	//以下两个值是由java编译器计算好了的
	maxStack     uint
	maxLocals    uint
}

func newMethods(class *Class, cfMethods []*chapter3_cf.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
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

func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags & ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags & ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags & ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags & ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags & ACC_STRICT
}

func (method *Method) MaxStack() uint {
	return method.maxStack
}

func (method *Method) MaxLocals() uint {
	return method.maxLocals
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}