package main

import (
	"flag"
	"fmt"
	"os"
	"GoVM/chapter2-class/classpath"
	"GoVM/chapter3-cf/classfile"
	"strings"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	//classpath option
	cpOption    string
	XjreOption  string
	class       string
	args        []string
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)

	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)

	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s \n", cmd.class)
	//	return
	//}
	//
	//fmt.Printf("class data:%v  \n", classData)
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