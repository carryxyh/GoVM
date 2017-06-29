package chapter3_cf

/**
	只会出现在ClassFile结构中，用于指出源文件名
	SOURCEFILE_ATTRIBUTE {
		u2 attribute_name_index; -> 指向一个UTF8常量
		u4 attribute_length; -> 必须是2
		u2 sourcefile_index;
	}
 */
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
