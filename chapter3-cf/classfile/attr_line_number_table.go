package chapter3_cf

/**
	存放方法的行号信息，是调试信息
	LINE_NUMBER_TABLE_ATTRIBUTE {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 line_number_table_length;
		{
			u2 start_pc;
			u2 lint_number;
		} line_number_table[line_number_table_length];
	}
 */
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:        reader.readUint16(),
			lineNumber:        reader.readUint16(),
		}
	}
}

/**
	这里为什么用了 pc >= entry.startPc ?

	紫苑注：
		看字节码，行号属性中，startPc 和 startPc 相差并不是1  有的是12 有的差16
		lineNumber 也并非递增，也存在lineNumber相等的LineNumberTableEntry
		没找到什么规律。。
 */
func (self *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(self.lineNumberTable) - 1; i >= 0; i-- {
		entry := self.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}