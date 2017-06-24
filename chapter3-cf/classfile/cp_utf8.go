package chapter3_cf

/**
	CONSTANT_UTF8_INFO {
		u1 tag;
		u2 length;
		u1 bytes[length];
	}
 */
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

/**
	简化版
 */
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}