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
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *chapter3_cf.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	methodDescriptor := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(methodDescriptor.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(methodDescriptor.returnType)
	}
	return method
}

/**
	本地方法没有code属性，所以需要给maxStack和maxLocals赋值
	本地方法栈帧操作数栈至少要能容纳返回值，暂时给maxStack赋值为4
	本地方法栈帧的局部变量表只能用来存放参数，所以把argSlotCount赋值给maxLocals
	code字段：
		本地方法的字节码，第一条指令都是0xFE
		第二条指令则根据函数的返回值选择相应的返回指令
 */
func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4 // todo
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
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

func (self *Method) calcArgSlotCount(paramTypes []string) {
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