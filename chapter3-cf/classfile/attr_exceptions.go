package chapter3_cf

/**
	记录方法抛出的异常表
	EXCEPTIONS_ATTRIBUTE {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 number_of_exceptions;
		u2 exception_index_table[number_of_exceptions];
	}
 */
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}