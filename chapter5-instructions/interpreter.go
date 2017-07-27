package chapter5_instructions

import (
	"GoVM/chapter4-rtdt"
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter6-obj/heap"
	"fmt"
)

func Interpret(method *heap.Method, logInst bool, args []string) {
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

	//转成字符串数组
	jArgs := createArgsArray(method.Class().Loader(), args)
	frame.LocalVars().SetRef(0, jArgs)

	defer catchErr(thread)
	loop(thread, logInst)
}

func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
}

func catchErr(thread *chapter4_rtdt.Thread) {
	if r := recover(); r != nil {
		//fmt.Printf("LocalVars: %v \n", frame.LocalVars())
		//fmt.Printf("OperandStack: %v \n", frame.OperandStack())
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *chapter4_rtdt.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUInt8()
		inst := NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if (logInst) {
			logInstruction(frame, inst)
		}

		//execute
		//fmt.Printf("pc : %2d inst:%T %v \n", pc, inst, inst)
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *chapter4_rtdt.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *chapter4_rtdt.Thread) {
	//for !thread.IsStackEmpty() {
	//	frame := thread.PopFrame()
	//	method := frame.Method()
	//	className := method.Class().Name()
	//	fmt.Printf(">> pc:%4d %v.%v%v \n",
	//		frame.NextPC(), className, method.Name(), method.Descriptor())
	//}
}