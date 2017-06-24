package chapter3_cf

/**
	string info本身不存储字符串，只存了常量池索引，这个索引指向一个CONSTANT_UTF8_INFO。
	CONSTANT_STRING_INFO {
		u1 tag;
		u2 string_index;
	}
 */
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
