package chapter3_cf

type ConstantPool []ConstantInfo

/**
	从classreader中取得常量池
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	return nil
}

/**
	获取常量信息
 */
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	return nil
}

/**
	获取名称和类型
 */
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	return nil
}

/**
	获取class name
 */
func (self ConstantPool) getClassName(index uint16) string {
	return ""
}

/**
	从常量池获取utf8编码的string
 */
func (self ConstantPool) getUtf8(index uint16) string {
	return nil;
}