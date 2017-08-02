package main

import (
	"GoVM/chapter6-obj/heap"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter2-class/classpath"
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter5-instructions"
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

}