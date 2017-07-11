package heap

import (
	"GoVM/chapter3-cf/classfile"
	"GoVM/chapter4-rtdt"
)

type Class struct {
	//类访问标识符
	accessFlags       uint16
	//thisClassName
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *chapter3_cf.ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlowCount uint
	staticSlotCount   uint
	staticVars        *chapter4_rtdt.Slots
}

func newClass(cf *chapter3_cf.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags & ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags & ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags & ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}

// getters
func (self *Class) ConstantPool() *chapter3_cf.ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}