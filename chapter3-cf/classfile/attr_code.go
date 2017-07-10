package chapter3_cf

/**
	CODE_ATTRIBUTE {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 max_stack; -> 操作数栈的最大深度
		u2 max_locals; -> 局部变量表大小
		u4 code_length;
		u1 code[code_length];
		u2 exception_table_length;
		{
			u2 start_pc;
			u2 end_pc;
			u2 handle_pc;
			u2 catch_type;
		} exception_table[exception_table_length];
		u2 attributes_count;
		attribute_info attributes[attributes_count]
	}
 */
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

/**
	异常表
 */
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}

func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}

func (self *CodeAttribute) Code() []byte {
	return self.code
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc :        reader.readUint16(),
			endPc :        reader.readUint16(),
			handlerPc :        reader.readUint16(),
			catchType :        reader.readUint16(),
		}
	}
	return exceptionTable
}