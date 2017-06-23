package chapter3_cf

type ConstantPool []ConstantInfo

/**
	从classreader中取得常量池
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	return nil
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	return nil
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	return nil
}

func (self ConstantPool) getClassName(index uint16) string {
	return ""
}

func (self ConstantPool) getUtf8(index uint16) string {
	return nil;
}