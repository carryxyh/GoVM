package heap

import (
	"GoVM/chapter3-cf/classfile"
	"GoVM/chapter4-rtdt"
	"strings"
)

type Class struct {
	//类访问标识符
	accessFlags       uint16
	//thisClassName
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	//实例变量(及private String name)占据的空间大小
	instanceSlotCount uint
	//类变量(及static类型的变量)占据的空间大小
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
func (self *Class) StaticVars() chapter4_rtdt.Slots {
	return self.staticVars
}

func (self *Class) NewObject() *chapter4_rtdt.Object {
	return chapter4_rtdt.NewObject(self)
}

/**
	是否有权限访问
 */
func (self *Class) isAccessibleTo(other *Class) bool {
	return other.IsPublic() || self.GetPackageName() == other.GetPackageName()
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}