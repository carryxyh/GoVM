package extended

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter5-instructions/loads"
	"GoVM/chapter5-instructions/stores"
	"GoVM/chapter5-instructions/math"
	"GoVM/chapter4-rtdt"
)

/**
	一般情况下，一个方法的局部变量表的大小都不会超过256，及1字节。
	但是如果有方法的局部变量表超过256，用wide指令来扩展一些指令
 */
type WIDE struct {
	modifiedInstruction base.Instruction
}

/**
	读取1字节的操作码，然后创建子指令实例，最后读取子指令的操作数，加载指令和存储指令都只有一个操作数，需要扩展成两个字节
 */
func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUInt8()
	switch opcode {
	case 0x15:                                //iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x16:                                //lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x17:                                //fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x18:                                //dload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x19:                                //aload
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x36:                                //istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x37:                                //lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x38:                                //fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x39:                                //dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x3a:                                //astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUInt16())
		self.modifiedInstruction = inst
	case 0x84:                                //iinc 两个操作数，都需要扩展成两个字节
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUInt16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0xa9:                                // ret
		panic("Unsupported opcode: 0xa9!")
	}
}

/**
	只改变索引宽度，不改变指令操作
 */
func (self *WIDE) Execute(frame *chapter4_rtdt.Frame) {
	self.modifiedInstruction.Execute(frame)
}