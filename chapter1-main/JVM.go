package main

import (
	"GoVM/chapter6-obj/heap"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter2-class/classpath"
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter5-instructions"
	"strings"
	"fmt"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *chapter4_rtdt.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:                cmd,
		classLoader:        classLoader,
		mainThread:        chapter4_rtdt.NewThread(),
	}
}

func (self *JVM) start() {
	self.initVM()
	self.execMain()
}

func (self *JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	chapter5_instructions.Interpret(self.mainThread, self.cmd.verboseInstFlag)
}

func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethd := mainClass.GetMainMethod()
	if mainMethd == nil {
		fmt.Printf("Main method not found in class %s \n", self.cmd.class)
		return
	}

	argsArr := self.createArgsArray()
	frame := self.mainThread.NewFrame(mainMethd)
	frame.LocalVars().SetRef(0, argsArr)
	self.mainThread.PushFrame(frame)
	chapter5_instructions.Interpret(self.mainThread, self.cmd.verboseInstFlag)
}

func (self *JVM) createArgsArray() *heap.Object {
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(self.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range self.cmd.args {
		jArgs[i] = heap.JString(self.classLoader, arg)
	}
	return argsArr
}