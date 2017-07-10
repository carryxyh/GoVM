package base

import "GoVM/chapter4-rtdt"

/**
	操作码后面可以跟零字节或多字节的操作数，操作数就相当于函数的参数
 */
type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	//执行具体指令
	Execute(frame *chapter4_rtdt.Frame)
}

/**
	没有操作数的指令，及没有参数的指令
 */
type NoOperandsInstruction struct {

}

/**
	什么也不做，这是个没有操作数的指令
 */
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

/**
	表示跳转指令
 */
type BranchInstruction struct {
	Offset int
}

/**
	读取一个16字节的int
 */
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/**
	存储和加载指令需要根据索引存取 -> 局部变量表
	索引由单字节操作数给出
	用Index表示局部变量表索引
 */
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUInt8())
}

/**
	用于指令访问运行时常量池，Index8Instruction用于访问局部变量表
	有一些指令需要访问运行时常量池，常量池索引由两个字节的操作数给出
 */
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUInt16())
}