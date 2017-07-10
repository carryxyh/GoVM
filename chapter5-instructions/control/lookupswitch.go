package control

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

/**
	matchOffsets 有点像map k->case值，v->跳转偏移量。
 */
func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *chapter4_rtdt.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs * 2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i + 1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}
