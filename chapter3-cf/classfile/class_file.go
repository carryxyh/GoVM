package chapter3_cf

import "fmt"

type ClassFile struct {
	//魔数 class的魔术 -> 0xCAFEBABE
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (this *ClassFile) read(reader *ClassReader) {
	this.readAndCheckMagic(reader)
	this.readAndCheckVersion(reader)
	this.constantPool = readConstantPool(reader)
	this.accessFlags = reader.readUint16()
	this.thisClass = reader.readUint16()
	this.superClass = reader.readUint16()
	this.interfaces = reader.readUint16s()
	this.fields = readMembers(reader, this.constantPool)
	this.methods = readMembers(reader, this.constantPool)
	this.attributes = readAttributes(reader, this.constantPool)
}

/**
	检查魔数
 */
func (this *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic error")
	}
}

/**
	检查字节码版本
 */
func (this *ClassFile) readAndCheckVersion(reader *ClassReader) {
	this.minorVersion = reader.readUint16()
	this.majorVersion = reader.readUint16()
	switch this.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if this.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

/**
	以下相当于Getter方法，用于暴露ClassFile的属性
 */
func (this *ClassFile) MinorVersion() uint16 {
	return this.minorVersion
}

func (this *ClassFile) MajorVersion() uint16 {
	return this.majorVersion
}

func (this *ClassFile) ConstantPool() ConstantPool {
	return this.constantPool
}

func (this *ClassFile) AccessFlags() uint16 {
	return this.accessFlags
}

func (this *ClassFile) Fields() []*MemberInfo {
	return this.fields
}

func (this *ClassFile) Methods() []*MemberInfo {
	return this.methods
}

func (this *ClassFile) ClassName() string {
	return this.constantPool.getClassName(this.thisClass)
}

func (this *ClassFile) SuperClassName() string {
	if this.superClass > 0 {
		return this.constantPool.getClassName(this.superClass)
	}
	//只有java.lang.Object没有超类
	return ""
}

func (this *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(this.interfaces))
	for i, cpIndex := range this.interfaces {
		interfaceNames[i] = this.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (self *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}