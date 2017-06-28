package chapter3_cf

/**
	属性表，储存了方法的字节码等信息
	attribute_info {
		u2 attribute_name_index;
		u4 attribute_length;
		u1 info[attribute_length];
	}
 */
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	return nil
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	return nil
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	return nil
}
