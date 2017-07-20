package main

import (
	"flag"
	"fmt"
	"os"
	"GoVM/chapter2-class/classpath"
	"GoVM/chapter3-cf/classfile"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
	"strings"
	"GoVM/chapter5-instructions"
)

type Cmd struct {
	helpFlag         bool
	versionFlag      bool
	verboseClassFlag bool
	verboseInstFlag  bool
	//classpath option
	cpOption         string
	XjreOption       string
	class            string
	args             []string
}

/**
	从命令行获取cmd
 */
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "version info")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.cpOption, "classpath", "", "class path")
	flag.StringVar(&cmd.cpOption, "cp", "", "equals classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...] \n", os.Args[0])
}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//第六节测试代码
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		chapter5_instructions.Interpret(mainMethod, cmd.verboseInstFlag)
	} else {
		fmt.Printf("main method not found in class %v \n", className)
	}

	//第五节测试代码
	//cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//className := strings.Replace(cmd.class, ".", "/", -1)
	//cf := loadClass(className, cp)
	//mainMethod := getMainMethod(cf)
	//if mainMethod != nil {
	//	chapter5_instructions.Interpret(mainMethod)
	//} else {
	//	fmt.Println("Main method not found")
	//}

	//第四节测试代码
	//frame := chapter4_rtdt.NewFrame(100, 100)
	//testLocals(frame.LocalVars())
	//testOperandStack(frame.OperandStack())

	//第三节测试代码
	//cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//className := strings.Replace(cmd.class, ".", "/", -1)
	//
	//cf := loadClass(className, cp)
	//fmt.Println(cmd.class)
	//printClassInfo(cf)

	//第二节测试代码
	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s \n", cmd.class)
	//	return
	//}
	//
	//fmt.Printf("class data:%v  \n", classData)
}

func testLocals(vars chapter4_rtdt.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *chapter4_rtdt.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

func loadClass(className string, cp *classpath.Classpath) *chapter3_cf.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := chapter3_cf.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *chapter3_cf.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}

func getMainMethod(classFile *chapter3_cf.ClassFile) *chapter3_cf.MemberInfo {
	for _, m := range classFile.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}