package chapter3_cf

/**
	常量池中主要存放两大类常量：字面量和符号引用
	字面量比较接近于 Java 层面的常量概念，如文本字符串、被声明为 final 的常量值等
	符号引用总结起来则包括了下面三类常量:
		1.类和接口的全限定名（即带有包名的 Class 名，如：org.lxh.test.TestClass）
		2.字段的名称和描述符（private、static 等描述符）
		3.方法的名称和描述符（private、static 等描述符）
 */
type ConstantPool []ConstantInfo

/**
	从classreader中取得常量池
	注意三个点：
	常量池其实就是一个表，里面放的都是ConstantInfo
	1.表头给出的常量池大小比实际 大1。表头n，那么常量池实际大小 n-1
	2.有效的常量池索引是 1 ~ n-1
	3.CONSTANT_Long_info和CONSTANT_Double_info各占两个位置 ：如果常量池中存在这两种常量，实际的常量比n-1还要少，而且在有效的 1 ~ n-1zhong的某些数也会变成无效索引
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	//表头
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	//索引从1开始，这里用了 <cpCount 说明index是从1到cpCount-1 及上文的1 ~ n-1
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			//占两个位置
			i++
		}
	}
	return cp
}

/**
	获取常量信息
 */
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

/**
	获取名称和类型
 */
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

/**
	获取class name
 */
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

/**
	从常量池获取utf8编码的string
 */
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}