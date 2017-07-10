package extended

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type GOTO_W struct {
	offset int
}

/**
	与goto唯一的区别在于索引从2字节变成了4字节
 */
func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *chapter4_rtdt.Frame) {
	base.Branch(frame, self.offset)
}
