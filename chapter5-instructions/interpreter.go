package chapter5_instructions

import (
	"GoVM/chapter4-rtdt"
	"fmt"
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter6-obj/heap"
)

func Interpret(method *heap.Method) {
	//获取code属性
	//codeAttr := methodInfo.CodeAttribute()
	//
	////获取局部变量表和操作数栈空间以及方法的字节码
	//maxLocals := codeAttr.MaxLocals()
	//maxStack := codeAttr.MaxStack()
	//bytecode := codeAttr.Code()

	//创建线程以及栈帧
	thread := chapter4_rtdt.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *chapter4_rtdt.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v \n", frame.LocalVars())
		fmt.Printf("OperandStack: %v \n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *chapter4_rtdt.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUInt8()
		inst := NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//execute
		fmt.Printf("pc : %2d inst:%T %v \n", pc, inst, inst)
		inst.Execute(frame)
	}
}