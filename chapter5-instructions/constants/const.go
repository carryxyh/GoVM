package constants

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
)

/**
	aconst_null 把 null 引用推入操作数栈顶
 */
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushRef(nil)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	dconst_0 把 double 型 0 推入操作数栈顶
 */
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_0) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	dconst_1 把 double 型 1 推入操作数栈顶
 */
type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (self DCONST_1) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	fconst_0 把 float 型 0 推入操作数栈顶
 */
type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_0) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	fconst_1 把 float 型 1 推入操作数栈顶
 */
type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_1) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	fconst_2 把 float 型 2 推入操作数栈顶
 */
type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_2) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	iconstn_m1 把 int 型 -1 推入操作数栈顶
 */
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushInt(-1)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	iconstn_0 把 int 型 0 推入操作数栈顶
 */
type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_0) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushInt(0)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	iconstn_1 把 int 型 1 推入操作数栈顶
 */
type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_1) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushInt(1)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	iconstn_2 把 int 型 2 推入操作数栈顶
 */
type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_2) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushInt(2)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	iconstn_3 把 int 型 3 推入操作数栈顶
 */
type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_3) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushInt(3)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	iconstn_4 把 int 型 4 推入操作数栈顶
 */
type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_4) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushInt(4)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	iconstn_5 把 int 型 5 推入操作数栈顶
 */
type ICONST_5 struct {
	base.NoOperandsInstruction
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	lconst_0 把 long 型 0 推入操作数栈顶
 */
type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushLong(0)
}

/*---------------------------------------------------------------------------------------------------------*/

/**
	lconst_1 把 long 型 1 推入操作数栈顶
 */
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_1) Execute(frame *chapter4_rtdt.Frame) {
	frame.OperandStack().PushLong(1)
}