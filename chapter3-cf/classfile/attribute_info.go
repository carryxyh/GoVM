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

/**
	从常量池中解析出属性表中的属性
 */
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

/**
	从常量池中解析一条属性表中的属性
 */
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attributeNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attributeNameIndex)
	attrLen := reader.readUint16()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":

	case "ConstantValue":

	case "Deprecated":

	case "Exceptions":

	case "LineNumberTable":

	case "LocalVariableTable":

	case "SourceFile":

	case "Synthetic":

	default:
		return &UnparsedAttribute{attrName, attrLen, nil}

	}
	return nil
}
