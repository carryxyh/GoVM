package chapter3_cf


/**
	ConstantFieldrefInfo、ConstantMethodrefInfo、ConstantInterfaceMethodrefInfo
	这三个结构体继承自ConstantMemberrefInfo
	Go语言没有“继承”的概念，而是通过结构体嵌套的方式实现的
 */
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

/**
	字段符号引用
	CONSTANT_FIELDREF_INFO {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}
 */
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

/**
	普通（非接口）方法符号引用
	CONSTANT_METHODREF_INFO {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}
 */
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

/**
	接口方法符号引用
	CONSTANT_INTERFACEMETHODREF_INFO {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}
 */
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}