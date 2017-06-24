package chapter3_cf

/**
	常量数据结构如下
	cp_info {
		u1 tag; -> 用来区分常量类型
		u2 info[];
	}
 */

const (
	CONSTANT_Class = 7
	CONSTANT_String = 8
	CONSTANT_Fieldref = 9
	CONSTANT_Methodref = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_Integer = 3
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_NameAndType = 12
	CONSTANT_Utf8 = 1
	CONSTANT_MethodHandle = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokeDynamic = 18
)

type ConstantInfo interface {
	//读取常量信息
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	return nil
}